package checker

import (
	"fmt"
	"log"
)

func check() bool {
	filedir, err := ConfigPath()
	if err != nil {
		log.Println("Config path error:", err)
		return false
	}
	fmt.Println(filedir)
	return true
}
