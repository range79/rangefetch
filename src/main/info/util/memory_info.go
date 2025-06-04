package util

import (
	"bytes"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func getMemoryInfoLinux() string {
	out, err := exec.Command("free", "-b").Output()
	if err != nil {
		return "N/A"
	}

	lines := strings.Split(string(out), "\n")
	if len(lines) < 2 {
		return "N/A"
	}

	fields := strings.Fields(lines[1])
	if len(fields) < 3 {
		return "N/A"
	}

	total, _ := strconv.ParseFloat(fields[1], 64)
	used, _ := strconv.ParseFloat(fields[2], 64)

	return formatMemory(used) + " / " + formatMemory(total)
}

func getMemoryInfoDarwin() string {

	totalBytes, err := exec.Command("sysctl", "-n", "hw.memsize").Output()
	if err != nil {
		return "N/A"
	}
	total, _ := strconv.ParseFloat(strings.TrimSpace(string(totalBytes)), 64)

	vmStat, err := exec.Command("vm_stat").Output()
	if err != nil {
		return "N/A"
	}

	pageSize := 4096.0
	usedPages := 0.0

	lines := bytes.Split(vmStat, []byte("\n"))
	for _, line := range lines {
		if bytes.Contains(line, []byte("Pages active")) ||
			bytes.Contains(line, []byte("Pages wired down")) ||
			bytes.Contains(line, []byte("Pages speculative")) {
			parts := bytes.Fields(line)
			if len(parts) >= 3 {
				countStr := strings.TrimSuffix(string(parts[2]), ".")
				count, _ := strconv.ParseFloat(countStr, 64)
				usedPages += count
			}
		}
	}

	used := usedPages * pageSize
	return formatMemory(used) + " / " + formatMemory(total)
}

func getMemoryInfoWindows() string {
	out, err := exec.Command("wmic", "OS", "get", "TotalVisibleMemorySize,FreePhysicalMemory", "/Value").Output()
	if err != nil {
		return "N/A"
	}

	lines := strings.Split(string(out), "\n")
	var totalKB, freeKB float64

	for _, line := range lines {
		if strings.HasPrefix(line, "TotalVisibleMemorySize=") {
			totalKB, _ = strconv.ParseFloat(strings.TrimPrefix(line, "TotalVisibleMemorySize="), 64)
		} else if strings.HasPrefix(line, "FreePhysicalMemory=") {
			freeKB, _ = strconv.ParseFloat(strings.TrimPrefix(line, "FreePhysicalMemory="), 64)
		}
	}

	used := (totalKB - freeKB) * 1024
	total := totalKB * 1024

	return formatMemory(used) + " / " + formatMemory(total)
}

func formatMemory(bytes float64) string {
	gb := bytes / (1024 * 1024 * 1024)
	return strconv.FormatFloat(gb, 'f', 1, 64) + " GB"
}

// Genel bellek bilgisi
func GetMemoryInfo() string {
	switch runtime.GOOS {
	case "linux":
		return getMemoryInfoLinux()
	case "darwin":
		return getMemoryInfoDarwin()
	case "windows":
		return getMemoryInfoWindows()
	default:
		return "Unsupported OS"
	}
}
