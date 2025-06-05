package app

import (
	"fmt"
	"rangefetch/src/main/checker"
	"rangefetch/src/main/info"
)

func Init() {
	checker.ConfigData()
	fmt.Println(info.GetSystemInfo())

}
