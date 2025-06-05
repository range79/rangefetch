package info

import (
	"fmt"
	"os"
	"path/filepath"
	"rangefetch/src/main/checker"
)

func LoadBannerFromAssets() string {
	osName := GetOSName()
	color := checker.ConfigData().Theme.ColorOutput // örnek: "green"

	bannerPath := filepath.Join("assets", osName+".txt")

	data, err := os.ReadFile(bannerPath)
	if err == nil {
		return ApplyColor(string(data), color)
	}

	defaultPath := filepath.Join("assets", "default.txt")
	fallback, err := os.ReadFile(defaultPath)
	if err == nil {
		return ApplyColor(string(fallback), color)
	}

	return fmt.Sprintf("[no banner found for %s]", osName)
}

// Renkleri uygula
func ApplyColor(text string, color string) string {
	colorMap := map[string]string{
		"red":     "\033[31m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"blue":    "\033[34m",
		"magenta": "\033[35m",
		"cyan":    "\033[36m",
		"reset":   "\033[0m",
	}

	colorCode, ok := colorMap[color]
	if !ok {
		colorCode = colorMap["reset"]
	}

	return fmt.Sprintf("%s%s%s", colorCode, text, colorMap["reset"])
}
