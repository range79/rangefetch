package util

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

func GetGPUInfoLinux() string {
	out, err := exec.Command("lspci").Output()
	if err != nil {
		return "N/A"
	}
	lines := strings.Split(string(out), "\n")
	var gpus []string

	extractModel := func(descr string, marka string) string {
		idx := strings.Index(descr, marka)
		model := descr
		if idx != -1 {
			model = descr[idx+len(marka):]
		}
		model = strings.TrimSpace(model)

		if start := strings.Index(model, "["); start != -1 {
			if end := strings.Index(model[start:], "]"); end != -1 {
				return strings.TrimSpace(model[start+1 : start+end])
			}
		}
		if start := strings.Index(model, "("); start != -1 {
			if end := strings.Index(model[start:], ")"); end != -1 {
				return strings.TrimSpace(model[:start])
			}
		}
		return model
	}

	for _, line := range lines {
		lower := strings.ToLower(line)
		if strings.Contains(lower, "vga compatible controller") || strings.Contains(lower, "3d controller") {
			parts := strings.SplitN(line, ":", 3)
			if len(parts) < 3 {
				continue
			}

			pciAddr := strings.TrimSpace(parts[0])
			descr := strings.TrimSpace(parts[2])

			var marka string
			if strings.Contains(descr, "NVIDIA") {
				marka = "NVIDIA"
			} else if strings.Contains(descr, "Intel") {
				marka = "Intel"
			} else if strings.Contains(descr, "AMD") || strings.Contains(descr, "ATI") || strings.Contains(descr, "Advanced Micro Devices") {
				marka = "AMD"
			} else {
				marka = "Unknown"
			}

			model := extractModel(descr, marka)

			final := fmt.Sprintf("%s : %s : %s", pciAddr, marka, model)
			gpus = append(gpus, final)
		}
	}

	if len(gpus) == 0 {
		return "No GPU found"
	}

	var sb strings.Builder
	for i, gpu := range gpus {
		sb.WriteString(fmt.Sprintf("GPU %d: %s\n", i, gpu))
	}

	return sb.String()
}

func GetGPUInfoWindows() string {
	cmd := exec.Command("wmic", "path", "win32_VideoController", "get", "name")
	out, err := cmd.Output()
	if err != nil {
		return "N/A"
	}
	lines := strings.Split(string(out), "\n")
	var names []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.Contains(strings.ToLower(line), "name") {
			continue
		}
		names = append(names, line)
	}
	if len(names) == 0 {
		return "No GPU found"
	}
	var sb strings.Builder
	for i, name := range names {
		sb.WriteString(fmt.Sprintf("GPU %d: %s\n", i, name))
	}
	return sb.String()
}

func GetGPUInfoMac() string {
	cmd := exec.Command("system_profiler", "SPDisplaysDataType")
	out, err := cmd.Output()
	if err != nil {
		return "N/A"
	}
	lines := strings.Split(string(out), "\n")
	var gpus []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Chipset Model:") {
			gpus = append(gpus, strings.TrimPrefix(line, "Chipset Model: "))
		}
	}
	if len(gpus) == 0 {
		return "No GPU found"
	}
	var sb strings.Builder
	for i, gpu := range gpus {
		sb.WriteString(fmt.Sprintf("GPU %d: %s\n", i, gpu))
	}
	return sb.String()
}
func GetFormattedGPUInfo() string {
	var gpu string

	switch runtime.GOOS {
	case "linux":

		gpu = GetGPUInfoLinux()
	case "windows":
		gpu = GetGPUInfoWindows()
	case "darwin":
		gpu = GetGPUInfoMac()
	default:
		return "Unsupported OS"
	}

	return gpu
}
