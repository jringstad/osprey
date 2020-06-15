package main

import (
	"./utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/*
SD card image creation
- change password
- install autossh
- add repos
- install bootstrapper
TODO: make autossh depend on bootstrapper startup, to make sure it's accessing the mounted cert?
*/

/*
- initialize
- mount USB drive and verify content
- add repo URL specified on USB drive
- apt-get update
- update self and reboot if update occurred
- check if osprey is already installed
- osprey is installed:
	- update osprey and reboot if update occurred
	- start osprey
- osprey is not installed:
	- install osprey
	- reboot
*/

func log(message string, soundName string) {
	//exec.Command("aplay", "sounds/"+soundName+".wav").Run()
	fmt.Println(message)
}

func MountKeyAndReadConfig() Config {
	// unmount in case it's already mounted, but ignore failure
	utils.RunCommand("sudo umount /mnt/osprey-key || true")
	utils.RunCommand("sudo mkdir -p /mnt/osprey-key")
	utils.RunCommand("sudo mount /dev/sda1 /mnt/osprey-key")
	jsonFile, err := os.Open("/mnt/osprey-key/osprey-config.json")
	utils.Check(err, "opening json file")
	defer jsonFile.Close()
	var config Config
	jsonBytes, err := ioutil.ReadAll(jsonFile)
	utils.Check(err, "reading json file")
	json.Unmarshal(jsonBytes, &config)
	return config
}

type Config struct {
	GroundstationUrl      string   `json:"groundstation-url"`
	RepoUrl               string   `json:"repo-url"`
	PackagesToInstall     []string `json:"packages-to-install"`
	ServicesToStart       []string `json:"services-to-start"`
	GroundstationCertPath string   `json:"groundstation-cert-path"`
}

// TODO: make bootstrapper depend on network connectivity to avoid startup failure?
func main() {
	log("bootstrapper initializing", "initializing")
	config := MountKeyAndReadConfig()
	utils.AddRepo(config.RepoUrl)                                                       // add repo and apt-get update
	wasUpdated := utils.UpdateOrInstallAndReboot([]string{"osprey-bootstrapper"})       // update self, reboot if changes were made
	wasUpdated = wasUpdated || utils.UpdateOrInstallAndReboot(config.PackagesToInstall) // install or update osprey, reboot if changes were made
	if wasUpdated {
		utils.Reboot()
	}
	utils.StartServices(config.ServicesToStart)
	log("bootstrapper exited successfully", "exited")
}
