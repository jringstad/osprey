package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	err := exec.Command("sudo", "bash", "-c", "echo \"deb [trusted=yes] https://osprey-groundstation.s3.amazonaws.com stable main\" > /etc/apt/sources.list.d/osprey.list").Run()
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
	enableAndStartService("osprey-diagnostics.service")
	enableAndStartService("osprey-diagnostics.timer")
}
