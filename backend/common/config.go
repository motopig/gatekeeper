package common

import (
	"os"

	"github.com/astaxie/beego/config"
)

var (
	Hconfig config.Configer
)

func InitConfig() {

	iniFileDir, err := os.Getwd()

	iniFilePath := iniFileDir + "/config/config.ini"

	_, err = os.Stat(iniFilePath)
	if err != nil {
		panic("config file not found")
	}

	iniconf, err := config.NewConfig("ini", iniFilePath)
	if err != nil {
		panic("config file init failed")
	}

	Hconfig = iniconf
}
