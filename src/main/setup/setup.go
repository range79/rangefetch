package setup

import (
	"fmt"
	"log"
	"time"
)

var welcomeShown = false

func Start() {
	if !welcomeShown {
		fmt.Println("Hey user, welcome to Rangefetch — where the magic (and code) happens!")
		fmt.Println("Would you like to install manually or automatically? (M/A)")
		welcomeShown = true
	}

	var choice string
	_, err := fmt.Scan(&choice)
	if err != nil {
		log.Println("Oops! An error occurred during setup.")
		time.Sleep(1 * time.Second)
		log.Print(err.Error())
		return
	}

	InstallMode(choice)
}
