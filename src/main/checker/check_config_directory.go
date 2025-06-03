package checker

import (
	"os"
	"path/filepath"
	"runtime"
)

func ConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "User home directory not found", err
	}

	switch runtime.GOOS {
	case "linux":

		return filepath.Join(homeDir, ".config", "rangefetch", "config.json"), nil

	case "windows":
		appData := os.Getenv("APPDATA")
		if appData == "" {
			return "appdata not found", err
		}

		return filepath.Join(appData, "rangefetch", "config.json"), nil

	case "darwin":

		return filepath.Join(homeDir, "Library", "Application Support", "rangefetch", "config.json"), nil

	case "android":

		return filepath.Join(homeDir, "rangefetch", "config.json"), nil

	default:

		return filepath.Join(homeDir, ".rangefetch_config.json"), nil
	}
}
