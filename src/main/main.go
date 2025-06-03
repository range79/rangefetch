package main

import (
	"fmt"
	"os"
	"rangefetch/src/main/checker"
	"time"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--settings" {
		time.Sleep(120)
		fmt.Println("xx")
	} else {
		checker.ConfigFinder()
	}

}
