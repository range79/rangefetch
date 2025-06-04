package auto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"rangefetch/src/main/common"
	"rangefetch/src/main/model"
)

func Install() {
	var selection = model.Config{
		Theme: model.ThemeSettings{
			ColorOutput:     "Blue",
			FontStyle:       "Default",
			UseDifferentimg: false,
			ImageSource:     "",
		},
		Display: model.DisplaySettings{
			HideGPU0:       false,
			HideGPU1:       false,
			HideUsername:   false,
			HideHostname:   false,
			HideKernel:     false,
			HideUptime:     false,
			HideResolution: false,
			HideShell:      false,
			HideDE:         false,
			HideWM:         false,
		},
	}

	path, err := common.ConfigPath()
	if err != nil {
		fmt.Printf("Error getting configuration path: %v\n", err)
		return
	}

	dir := filepath.Dir(path)

	err = os.MkdirAll(dir, 0755)
	if err != nil {
		fmt.Printf("Error creating directories: %v\n", err)
		return
	}

	jsonData, err := json.MarshalIndent(selection, "", "    ")
	if err != nil {
		fmt.Printf("Error converting to JSON: %v\n", err)
		return
	}

	err = ioutil.WriteFile(path, jsonData, 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}

	fmt.Printf("Configuration successfully written to '%s'!\n", path)
}
