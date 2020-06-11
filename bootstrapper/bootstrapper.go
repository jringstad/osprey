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

func downloadPackage(url string, dest string) bool {
	os.Chdir(dest)
	err := exec.Command("sudo", "wget", url).Run()
	return err == nil
}

func installPackage(basePath string, packagePath string) {
	os.Chdir(basePath)
	err := exec.Command("sudo", "tar", "-xf", packagePath).Run()
	if err != nil {
		log("failed to extract platform", "diagnostics-failed")
		panic("failure")
	}
	err = exec.Command("bash", "-c", "sudo cp misc/* /etc/misc/system/").Run()
	if err != nil {
		log("failed to extract platform", "diagnostics-failed")
		panic("failure")
	}
}

func enableAndStartService(serviceName string) {
	err := exec.Command("sudo", "systemctl", "daemon-reload").Run()
	if err != nil {
		log("failed to reload daemons", "diagnostics-failed")
		panic("failure")
	}
	err = exec.Command("sudo", "systemctl", "enable", serviceName).Run()
	if err != nil {
		log("failed to enable service", "diagnostics-failed")
		panic("failure")
	}
	err = exec.Command("sudo", "systemctl", "start", serviceName).Run()
	if err != nil {
		log("failed to start service", "diagnostics-failed")
		panic("failure")
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func addRepo() {
	d1 := []byte("deb [trusted=yes] https://osprey-groundstation.s3.amazonaws.com stable main\n")
	err := ioutil.WriteFile("/etc/apt/sources.list.d/osprey.list", d1, 0644)
	check(err)
	err = exec.Command("sudo", "apt-get", "update").Run()
	check(err)
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
	// TODO: replace all of this with an apt repo and generate packages
	addRepo()
	err := exec.Command("sudo", "apt-get", "install", "osprey-diagnostics").Run()
	check(err)
	
	/*
	var diagnosticsBasePath = "/opt/osprey/diagnostics/"

	log("downloading diagnostics platform", "diagnostics-downloading-platform")
	exec.Command("sudo", "mkdir", "-p", diagnosticsBasePath).Run()
	success := downloadPackage(config.DiagnosticsPlatformUrl, diagnosticsBasePath)
	if !success {
		log("failed to download platform", "diagnostics-downloading-failure")
		panic("failure")
	}
	installPackage(diagnosticsBasePath, "platform-diagnostics.tar.bz2")
	enableAndStartService("osprey-diagnostics.timer")
	enableAndStartService("osprey-diagnostics.service")
	*/
}
