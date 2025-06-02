package setup

import (
	"log"
	"rangefetch/setup/auto"
	"rangefetch/setup/manual"
	"strings"
)

func InstallMode(input string) {
	switch strings.ToLower(input) {
	case "a":

		auto.Install()

	case "m":
		manual.Install()

	default:
		{
			log.Printf("Error: Invalid choice '%s'. Please try again.\n", input)
		}

	}
}
