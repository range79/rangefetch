package util

import (
	"os/exec"
	"runtime"
	"strings"
)

func GetUptime() string {
	switch runtime.GOOS {
	case "linux", "darwin":
		//mac and linux uses same command
		out, err := exec.Command("uptime", "-p").Output()
		if err != nil {
			return "N/A"
		}
		return strings.TrimSpace(string(out))

	case "windows":
		// Windows get uptime
		out, err := exec.Command("cmd", "/C", "net stats srv").Output()
		if err != nil {
			return "N/A"
		}
		lines := strings.Split(string(out), "\n")
		for _, line := range lines {
			if strings.Contains(line, "Statistics since") || strings.Contains(line, "İstatistikler") {
				return strings.TrimSpace(line)
			}
		}
		return "Uptime info not found"

	default:
		return "Unsupported OS"
	}
}
