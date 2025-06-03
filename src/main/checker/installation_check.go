package checker

import (
	"os"
)

func CheckFileExits() bool {
	configPath, err := ConfigPath()

	stat, err := os.Stat(configPath)
	//if dir not exits
	if os.IsNotExist(err) {
		return false
	}

	if err != nil {
		return false
	}

	if stat.Size() == 0 {
		return false
	}

	return true
}
