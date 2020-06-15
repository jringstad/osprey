package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"sync"
)

func copyAndCapture(w io.Writer, r io.Reader) ([]byte, error) {
	var out []byte
	buf := make([]byte, 1024, 1024)
	for {
		n, err := r.Read(buf[:])
		if n > 0 {
			d := buf[:n]
			out = append(out, d...)
			_, err := w.Write(d)
			if err != nil {
				return out, err
			}
		}
		if err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				err = nil
			}
			return out, err
		}
	}
}

func Check(e error, operationThatFailed string) {
	if e != nil {
		fmt.Println("Operation failed: " + operationThatFailed)
		panic(e)
	}
}

func RunCommand(command string) {
	fmt.Println(command)
	cmd := exec.Command("bash", "-c", command)

	var stdout, stderr []byte
	var errStdout, errStderr error
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}

	// cmd.Wait() should be called only after we finish reading
	// from stdoutIn and stderrIn.
	// wg ensures that we finish
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		stdout, errStdout = copyAndCapture(os.Stdout, stdoutIn)
		wg.Done()
	}()

	stderr, errStderr = copyAndCapture(os.Stderr, stderrIn)

	wg.Wait()

	err = cmd.Wait()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	if errStdout != nil || errStderr != nil {
		log.Fatal("failed to capture stdout or stderr\n")
	}
	outStr, errStr := string(stdout), string(stderr)
	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
}

func ReadConfig() Config {
	jsonFile, err := os.Open("/mnt/osprey-key/osprey-config.json")
	Check(err, "opening json file")
	defer jsonFile.Close()
	var config Config
	jsonBytes, err := ioutil.ReadAll(jsonFile)
	Check(err, "reading json file")
	json.Unmarshal(jsonBytes, &config)
	return config
}

type Config struct {
	Groundstation         string   `json:"groundstation"`
	RepoUrl               string   `json:"repo-url"`
	PackagesToInstall     []string `json:"packages-to-install"`
	ServicesToStart       []string `json:"services-to-start"`
	GroundstationCertPath string   `json:"groundstation-cert-path"`
	GroundstationUser     string   `json:"groundstation-user"`
	PlatformIdentifier    string   `json:"platform-identifier"`
}

// TODO: add autossh dependency
func main() {
	config := ReadConfig()
	command := `AUTOSSH_GATETIME=0 
AUTOSSH_POLL=20 \
AUTOSSH_PORT=0 \
autossh -M 0 -o "ServerAliveInterval 5" -o "ServerAliveCountMax 3" \
%s@%s -N -R 8090:localhost:22 -vv \
-i /mnt/osprey-key/groundstation-cert.pem`
	RunCommand(fmt.Sprintf(command, config.GroundstationUser, config.Groundstation))
}
/*
AUTOSSH_GATETIME=0 \
AUTOSSH_POLL=20 \
AUTOSSH_PORT=0 \
autossh -M 0 -o "ServerAliveInterval 5" -o "ServerAliveCountMax 3" \
ubuntu@ec2-3-250-190-252.eu-west-1.compute.amazonaws.com -N -R 8090:localhost:22 -vv \
-i /mnt/osprey-key/groundstation-cert.pem
 */