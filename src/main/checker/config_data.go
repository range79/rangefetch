package checker

import (
	"encoding/json"
	"log"
	"os"
	m "rangefetch/src/main/model"
)

func ConfigData() m.Config {
	configPath := ConfigFinder()

	file, err := os.Open(configPath)
	if err != nil {
		log.Fatal("Error: Unable to open config.json. Please make sure the file exists and is readable.")
	}
	defer file.Close()

	var config m.Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		log.Fatal("Error: Failed to parse config.json:", err)
	}

	return config
}
