package info

import (
	"os"
	"runtime"
	"strings"
)

func GetOSName() string {
	var os = runtime.GOOS
	switch os {
	case "windows":
		return "windows"
	case "linux":
		return getLinuxDistro()
	case "darwin":
		return "darwin"
	default:
		return runtime.GOOS
	}
}

func getLinuxDistro() string {
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return "Linux"
	}

	lines := strings.SplitSeq(string(data), "\n")
	for line := range lines {
		if strings.HasPrefix(line, "ID=") {

			val := strings.TrimPrefix(line, "ID=")
			val = strings.Trim(val, `"`)
			return val
		}
	}
	return "Linux"
}
