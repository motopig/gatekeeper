package keeper

import (
	"github.com/motopig/gatekeeper/gate/keeper/system"
)

func Run() {
	system.Configer()
	system.InitLoger()
	system.NewConsul()
	Routers()
}
