package auto

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"rangefetch/src/main/common"
	"rangefetch/src/main/model"
)

const (
	ownerCanReadWriteGroupOtherCanReadOnly    os.FileMode = 0644
	ownerCanDoAllGroupOtherCanReadExecuteOnly os.FileMode = 0755
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

	err = os.MkdirAll(dir, ownerCanDoAllGroupOtherCanReadExecuteOnly)
	if err != nil {
		fmt.Printf("Error creating directories: %v\n", err)
		return
	}

	jsonData, err := json.MarshalIndent(selection, "", "    ")
	if err != nil {
		fmt.Printf("Error converting to JSON: %v\n", err)
		return
	}

	err = os.WriteFile(path, jsonData, ownerCanReadWriteGroupOtherCanReadOnly)
	if err != nil {
		fmt.Printf("Error writing file: %v\n", err)
		return
	}

	fmt.Printf("Configuration successfully written to '%s'!\n", path)
}
