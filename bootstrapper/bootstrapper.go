package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

/*
- initialize
- check internet connection
- get environment variables
- install prerequisites
- download platform diagnostics
- download platform
- register with basestation
- save ID
- initialize platform diagnostics
- initialize platform
*/

const BootstrapperConfigFile = "https://osprey-groundstation.s3-eu-west-1.amazonaws.com/bootstrapper/bootstrapper-config.json"

func runCommand(command string) string {
	out, _ := exec.Command("bash", "-c", command).Output()
	return strings.TrimSpace(string(out))
}
func log(message string, soundName string) {
	exec.Command("aplay", "sounds/"+soundName+".wav").Run()
}

func isInterfaceUp(networkInterface string) bool {
	err := exec.Command("bash", "-c", "ifconfig "+networkInterface+" | grep \"RUNNING\"").Run()
	return err == nil
}

func getConfigAndCheckConnectivity() Config {
	log("bootstrapper testing uplink", "uplink-test")
	resp, err := http.Get(BootstrapperConfigFile)
	if err != nil {
		log("failed to establish uplink (couldn't get config file)", "uplink-failed")
		panic("failed to establish uplink")
	}
	jsonBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log("failed to establish uplink (failed to read config file)", "uplink-failed")
		panic("failed to establish uplink")
	}
	var response Config
	err = json.Unmarshal(jsonBytes, &response)
	if err != nil {
		log("failed to establish uplink (failed to parse)", "uplink-failed")
		panic("failed to establish uplink")
	}

	if isInterfaceUp("eth0") {
		log("uplink success (wired)", "uplink-wired")
	} else {
		// assuming cellular
		log("uplink success (success)", "uplink-cellular")
	}
	return response
}

func downloadPackage(url string, dest string) {
	exec.Command("wget", url, "-O", dest).Run()
}

func installPackage(basePath string, packagePath string) {
	exec.Command("mkdir", "-p", basePath).Run()
	os.Chdir(basePath)
	exec.Command("bunzip2", packagePath).Run()
}

type Config struct {
	GroundstationUrl       string `json:"groundstation-url"`
	DiagnosticsPlatformUrl string `json:"diagnostics-platform"`
}

func main() {
	log("bootstrapper initializing", "initializing")
	var config = getConfigAndCheckConnectivity()
	fmt.Println(config)
	// install prerequisites -- currently none
	// download diagnostic platform
	downloadPackage(config.DiagnosticsPlatformUrl, "/tmp/platform-diagnostics.tar.bz2")
	installPackage("/opt/osprey/diagnostics/", "platform-diagnostics.tar.bz2")
}
