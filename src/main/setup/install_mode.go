package setup

import (
	"log"
	"rangefetch/src/main/setup/auto"
	"rangefetch/src/main/setup/manual"
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
			Start()
		}

	}
}
