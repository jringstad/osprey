package main

import (
	"os/exec"
	"strings"
)

/*
- get environment variables (S3 bucket, EC2 gateway)
- initialize
- check internet connection
- install prerequisites
- download platform diagnostics
- download platform
- register with basestation
- save ID
- initialize platform diagnostics
- initialize platform
 */

func runCommand(command string) string {
	out, _ := exec.Command("bash", "-c", command).Output()
	return strings.TrimSpace(string(out))
}
func log(message string, soundName string) {
	exec.Command("aplay", "misc/" + soundName + ".wav").Run()
}
func checkConnectivity() {
	// TODO: hit up ground station instead
	err := exec.Command("ping", "-c", "4", "google.com").Run()
	if err != nil {
		log("failed to establish uplink", "uplink-failed")
		panic("failed to establish uplink")
	} else {
		// TODO: detect actual uplink
		log("Uplink established", "uplink-cellular")
	}
}

func main() {
	log("bootstrapper initializing", "initializing")
	checkConnectivity()
}