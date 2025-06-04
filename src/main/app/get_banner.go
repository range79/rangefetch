package app

import (
	"fmt"
	"os"
	"path/filepath"
	"rangefetch/src/main/info"
)

func LoadBannerFromAssets() string {
	osName := info.GetOSName() // önceki fonksiyonun

	// assets klasöründen dosya yolu
	bannerPath := filepath.Join("assets", osName+".txt")

	// Dosya oku
	data, err := os.ReadFile(bannerPath)
	if err == nil {
		return string(data)
	}

	// Fallback: default.txt dosyası
	defaultPath := filepath.Join("assets", "default.txt")
	fallback, err := os.ReadFile(defaultPath)
	if err == nil {
		return string(fallback)
	}

	// Hiçbiri yoksa:
	return fmt.Sprintf("[no banner found for %s]", osName)
}
