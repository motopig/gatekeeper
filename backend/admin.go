package main

import (
	"fmt"

	"github.com/motopig/gatekeeper/backend/common"
	"github.com/motopig/gatekeeper/backend/model"
	"github.com/motopig/gatekeeper/backend/router"
	"github.com/motopig/gatekeeper/gate/keeper/system"
)

func main() {

	common.InitConfig()
	model.InitMysql()
	common.InitCache()
	system.NewConsul()
	fmt.Println(common.Hcache.Get("admin"))
	router.Routers()

}
