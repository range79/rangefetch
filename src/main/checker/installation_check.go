package checker

import (
	"os"
	"rangefetch/src/main/common"
)

func CheckFileExits() bool {
	configPath, err := common.ConfigPath()

	stat, err := os.Stat(configPath)

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
