package main

import (
	"./utils"
	"encoding/json"
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

func MountKeyAndReadConfig() Config {
	defer utils.RecoverErrorMessage("Failed to find key", "key-failure")
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
	utils.Log("Key is present", "key-success")
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

// TODO: make bootstrapper depend on network connectivity to avoid startup failure?
// TODO: make Log() only take one argument and generate string from text?
func main() {
	utils.Log("bootstrapper initializing", "initializing")
	utils.Log("Checking for key", "key-checking")
	config := MountKeyAndReadConfig()
	utils.AddRepo(config.RepoUrl) // add repo and apt-get update
	utils.Log("self-updating", "self-update")
	wasUpdated := utils.UpdateOrInstallAndReboot([]string{"osprey-bootstrapper"}) // update self, reboot if changes were made
	utils.Log("installing platform", "installing-platform")
	wasUpdated = wasUpdated || utils.UpdateOrInstallAndReboot(config.PackagesToInstall) // install or update osprey, reboot if changes were made
	if wasUpdated {
		utils.Log("packages were updated, rebooting", "rebooting")
		utils.Reboot()
	}
	utils.Log("Nothing needed updating", "no-updates")
	utils.Log("Starting platform services", "starting-platform-services")
	utils.StartServices(config.ServicesToStart)
	utils.Log("bootstrapper exited successfully", "finished")
}
