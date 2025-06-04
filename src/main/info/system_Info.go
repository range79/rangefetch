package info

import (
	"fmt"
	"os"
	"os/exec"
	"rangefetch/src/main/info/util"
	"runtime"
	"strings"
)

func GetSystemInfo() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("OS: %s\n", GetOSName()))
	hostname, err := os.Hostname()
	if err == nil {
		sb.WriteString(fmt.Sprintf("Hostname: %s\n", hostname))
	}
	sb.WriteString(fmt.Sprintf("%v\n", util.GetFormattedCPUInfo()))
	sb.WriteString(fmt.Sprintf("Local IP: %v\n", util.GetLocalIP()))
	sb.WriteString(fmt.Sprintf("Public IP: %s\n", util.GetPublicIP()))
	sb.WriteString(fmt.Sprintf("Memory: %s\n", util.GetMemoryInfo()))
	sb.WriteString(fmt.Sprintf("Gpu: %s", util.GetFormattedGPUInfo()))

	if runtime.GOOS == "linux" {
		sb.WriteString(fmt.Sprintf("Uptime: %s\n", getUptimeLinux()))

	} else if runtime.GOOS == "windows" {

	} else if runtime.GOOS == "darwin" {

	}

	return sb.String()
}

func getUptimeLinux() string {
	out, err := exec.Command("uptime", "-p").Output()
	if err != nil {
		return "N/A"
	}
	return strings.TrimSpace(string(out))
}
