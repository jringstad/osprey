package autossh

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func Check(e error, operationThatFailed string) {
	if e != nil {
		fmt.Println("Operation failed: " + operationThatFailed)
		panic(e)
	}
}

func RunCommand(command string) string {
	cmd := exec.Command("bash", "-c", command)
	out, err := cmd.CombinedOutput()
	Check(err, "executing command '"+command+"': "+string(out))
	return strings.TrimSpace(string(out))
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
}

func main() {
	config := ReadConfig()
	command := `AUTOSSH_GATETIME=0 \
AUTOSSH_POLL=20 \
AUTOSSH_PORT=0 \
autossh -f -M 0 -o \"ServerAliveInterval 5\" -o \"ServerAliveCountMax 3\" \
%s@%s -N -R 8090:localhost:22 -vv \
-i /mnt/osprey-key/groundstation-cert.pem`
	RunCommand(fmt.Sprintf(command, config.GroundstationUser, config.Groundstation))
}
