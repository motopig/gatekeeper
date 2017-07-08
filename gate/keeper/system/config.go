package system

import (
	"os"

	"github.com/astaxie/beego/config"
)

var (
	iniconf config.Configer
)

func Configer() {

	iniFileDir, err := os.Getwd()

	iniFilePath := iniFileDir + "/config/config.ini"

	_, err = os.Stat(iniFilePath)
	if err != nil {
		panic("config file not found")
	}

	iniconf, err = config.NewConfig("ini", iniFilePath)
	if err != nil {
		panic("config file init failed")
	}

}

func GetConfig(key string) string {
	return iniconf.String(key)
}
