package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"./utils"
)

/*
SD card image creation
- change password
- install autossh
- add repos
- install bootstrapper
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
	exec.Command("aplay", "sounds/"+soundName+".wav").Run()
}

func MountKeyAndReadConfig() Config {
	utils.RunCommand("sudo mkdir /mnt/osprey-key")
	utils.RunCommand("sudo mount /dev/sdb1 /mnt/osprey-key")
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
	GroundstationUrl       string `json:"groundstation-url"`
	RepoUrl				   string `json:"repo-url"`
	PackagesToInstall      []string `json:"packages-to-install"`
	ServicesToStart		   []string `json:"services-to-start"`
}

// TODO: make bootstrapper depend on network connectivity to avoid startup failure?
func main() {
	log("bootstrapper initializing", "initializing")
	config := MountKeyAndReadConfig()
	utils.AddRepo(config.RepoUrl)                                       // add repo and apt-get update
	utils.UpdateOrInstallAndReboot([]string{"osprey-bootstrapper"}) // update self, reboot if changes were made
	utils.UpdateOrInstallAndReboot(config.PackagesToInstall)     // install or update osprey, reboot if changes were made
	utils.StartServices(config.ServicesToStart)
}
