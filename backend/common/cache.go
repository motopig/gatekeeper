package common

import (
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
)

var (
	Hcache cache.Cache
)

func InitCache() {

	cache, err := cache.NewCache(Hconfig.String("cache::adapter"), `{"interval":60}`)
	if err != nil {
		panic("huake cache init shibai")
	}
	cache.Put("admin", "king of the pop", 0)

	Hcache = cache

}
