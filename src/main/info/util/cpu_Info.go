package util

import (
	"bufio"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

func cleanCPUModel(raw string) string {

	cleaned := strings.ToLower(raw)
	cleaned = strings.ReplaceAll(cleaned, " with radeon graphics", "")
	cleaned = strings.ReplaceAll(cleaned, " radeon graphics", "")
	cleaned = strings.ReplaceAll(cleaned, " with", "")

	re := regexp.MustCompile(`\s+`)
	cleaned = re.ReplaceAllString(cleaned, " ")
	return strings.Title(strings.TrimSpace(cleaned))
}

func getCPUInfoLinux() string {
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return "N/A"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "model name") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				return cleanCPUModel(parts[1])
			}
		}
	}
	return "N/A"
}

func getCPUSpeedLinux() string {
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		return "N/A"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "cpu MHz") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				mhz := strings.TrimSpace(parts[1])
				val, err := strconv.ParseFloat(mhz, 64)
				if err != nil {
					return "N/A"
				}
				ghz := val / 1000.0
				return strconv.FormatFloat(ghz, 'f', 2, 64) + " GHz"
			}
		}
	}
	return "N/A"
}

func getCPUInfoWindows() string {
	out, err := exec.Command("wmic", "cpu", "get", "Name").Output()
	if err != nil {
		return "N/A"
	}
	lines := strings.Split(string(out), "\n")
	if len(lines) > 1 {
		return cleanCPUModel(lines[1])
	}
	return "N/A"
}

func getCPUSpeedWindows() string {
	out, err := exec.Command("wmic", "cpu", "get", "CurrentClockSpeed").Output()
	if err != nil {
		return "N/A"
	}
	lines := strings.Split(string(out), "\n")
	if len(lines) > 1 {
		speed := strings.TrimSpace(lines[1])
		val, err := strconv.ParseFloat(speed, 64)
		if err != nil {
			return "N/A"
		}
		ghz := val / 1000.0
		return strconv.FormatFloat(ghz, 'f', 2, 64) + " GHz"
	}
	return "N/A"
}

// MacOS CPU modeli
func getCPUInfoDarwin() string {
	out, err := exec.Command("sysctl", "-n", "machdep.cpu.brand_string").Output()
	if err != nil {
		return "N/A"
	}
	return cleanCPUModel(string(out))
}

func getCPUSpeedDarwin() string {
	out, err := exec.Command("sysctl", "-n", "hw.cpufrequency").Output()
	if err != nil {
		return "N/A"
	}
	outStr := strings.TrimSpace(string(out))
	hz, err := strconv.ParseInt(outStr, 10, 64)
	if err != nil {
		return "N/A"
	}
	ghz := float64(hz) / 1e9
	return strconv.FormatFloat(ghz, 'f', 2, 64) + " GHz"
}

func GetFormattedCPUInfo() string {
	var model, speed string
	var cores int = runtime.NumCPU()

	switch runtime.GOOS {
	case "linux":
		model = getCPUInfoLinux()
		speed = getCPUSpeedLinux()
	case "windows":
		model = getCPUInfoWindows()
		speed = getCPUSpeedWindows()
	case "darwin":
		model = getCPUInfoDarwin()
		speed = getCPUSpeedDarwin()
	default:
		return "Unsupported OS"
	}

	return "CPU: " + model + " (" + strconv.Itoa(cores) + ") @ " + speed
}
