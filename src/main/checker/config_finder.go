package checker

import (
	"log"
	"rangefetch/src/main/setup"
)

func ConfigFinder() string {

	if !CheckFileExits() {
		setup.Start()
	}

	name, err := ConfigPath()
	if err != nil {
		log.Println(err)
		return ""
	}

	return name
}
