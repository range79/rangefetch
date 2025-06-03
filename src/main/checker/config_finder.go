package checker

import (
	"log"
	"rangefetch/src/main/common"
	"rangefetch/src/main/setup"
)

func ConfigFinder() string {

	if !CheckFileExits() {
		setup.Start()
	}

	name, err := common.ConfigPath()
	if err != nil {
		log.Println(err)
		return ""
	}

	return name
}
