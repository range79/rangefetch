package info

import (
	"fmt"
	"os"
	"rangefetch/src/main/info/util"
	"strings"
)

func GetSystemInfo() string {
	const bannerWidth = 35
	const spaceBetween = 3
	const maxTotalWidth = 96

	bannerLines := strings.Split(LoadBannerFromAssets(), "\n")

	infoLines := []string{
		fmt.Sprintf("OS: %s", GetOSName()),
	}

	if hostname, err := os.Hostname(); err == nil {
		infoLines = append(infoLines, fmt.Sprintf("Hostname: %s", hostname))
	}

	infoLines = append(infoLines,
		fmt.Sprintf("Cpu: %s", util.GetFormattedCPUInfo()),
		fmt.Sprintf("Local IP: %s", util.GetLocalIP()),
		fmt.Sprintf("Public IP: %s", util.GetPublicIP()),
		fmt.Sprintf("Memory: %s", util.GetMemoryInfo()),
	)

	gpuInfo := util.GetFormattedGPUInfo()
	gpuLines := strings.Split(strings.TrimSpace(gpuInfo), "\n")
	infoLines = append(infoLines, gpuLines...)

	// Uptime da son satır olarak ekle
	infoLines = append(infoLines, fmt.Sprintf("Uptime: %s", util.GetUptime()))

	if len(infoLines) > len(bannerLines) {
		diff := len(infoLines) - len(bannerLines)
		for i := 0; i < diff; i++ {
			bannerLines = append(bannerLines, "")
		}
	}

	var sb strings.Builder

	for i := 0; i < len(bannerLines); i++ {
		bannerFormatted := fmt.Sprintf("%-*s", bannerWidth, bannerLines[i])

		infoLine := ""
		if i < len(infoLines) {

			availableWidth := maxTotalWidth - bannerWidth - spaceBetween
			line := infoLines[i]
			if len(line) > availableWidth {
				line = line[:availableWidth-3] + "..."
			}
			infoLine = line
		}

		sb.WriteString(fmt.Sprintf("%s%s%s\n", bannerFormatted, strings.Repeat(" ", spaceBetween), infoLine))
	}

	return sb.String()
}
