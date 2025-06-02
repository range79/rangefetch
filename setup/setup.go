package setup

import (
	"fmt"
	"log"
	"time"
)

func Start() {

	fmt.Println("Hey user, welcome to Rangefetch — where the magic (and code) happens!")
	fmt.Println("Would you like to install manually or automatically?(M/A)")
	var choice string

	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Println("Oops! An error occurred during setup.")
		time.Sleep(1000)
		log.Print(err.Error())
		return
	}
	InstallMode(choice)

}
