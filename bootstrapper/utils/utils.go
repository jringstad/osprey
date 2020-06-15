package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

func RunCommand(command string) string {
	out, err := exec.Command("bash", "-c", command).Output()
	Check(err, "executing command '" + command + "': " + string(out))
	return strings.TrimSpace(string(out))
}

func Check(e error, operationThatFailed string) {
	if e != nil {
		fmt.Println("Operation failed: " + operationThatFailed)
		panic(e)
	}
}

func AddRepo(repoUrl string) {
	// delete repo if it already exists, in case it changed, but don't fail if it didn't
	RunCommand("sudo rm -rf /etc/apt/sources.list.d/osprey.list || true")
	RunCommand("sudo bash -c echo \"deb [trusted=yes] " + repoUrl + " stable main\" > /etc/apt/sources.list.d/osprey.list")
	RunCommand("sudo apt-get update")
}

func UpdateOrInstallAndReboot(packages []string) {
	// ...
}

func StartServices(services []string) {
	for _, service := range services {
		RunCommand("sudo systemctl start " + service)
	}
}

/*
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
}*/