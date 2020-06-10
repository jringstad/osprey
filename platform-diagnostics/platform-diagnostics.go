package main

import (
	"./metrics/linux"
	"./metrics/rpi"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

// resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
const BaseUrl = "http://ec2-3-250-190-252.eu-west-1.compute.amazonaws.com:61714"

func submitReport(r map[string]interface{}) {
	bytesOut, _ := json.Marshal(r)
	resp, _ := http.Post(BaseUrl + "/telemetry/report2", "application/json", bytes.NewReader(bytesOut))
	fmt.Println(resp)
}

func runCommand(command string) string {
	out, _ := exec.Command("bash", "-c", command).Output()
	return strings.TrimSpace(string(out))
}

func collect() map[string]interface{} {
	report := map[string]interface{}{}
	for key, command := range linux.Commands {
		res := runCommand(command)
		report[key] = res
	}
	for key, command := range rpi.Commands {
		res := runCommand(command)
		report[key] = res
	}
	return report
}

func main() {
	report := collect()
	submitReport(report)
}

/*
upon regist
/proc/cpuinfo
 */
