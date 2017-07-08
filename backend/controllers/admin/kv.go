package admin

import (
	"net/http"
	"strings"

	"fmt"

	"github.com/gin-gonic/gin"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/motopig/gatekeeper/gate/keeper/system"
	"github.com/motopig/gatekeeper/gate/keeper/util"
)

func GetKv() *consulapi.KV {
	return system.Client().KV()
}

func Kvs(c *gin.Context) {
	c.HTML(http.StatusOK, "kv.tmpl", gin.H{})
}

func Keys(c *gin.Context) {
	dc := CheckDc(c)
	perfix := c.Param("key")

	var kvs interface{}
	var err error
	if strings.HasSuffix(perfix, "/") == false && perfix != "" {
		kvs, _, err = GetKv().Get(perfix, &consulapi.QueryOptions{Datacenter: dc})
	} else {
		kvs, _, err = GetKv().Keys(perfix, "/", &consulapi.QueryOptions{Datacenter: dc})
	}

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"msg": "key not found!"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"kvs": kvs,
	})
}

func Put(c *gin.Context) {
	dc := CheckDc(c)
	key := c.Param("key")
	value := c.PostForm("val")
	key = util.Substr(key, 1, len(key)) // 死坑

	_, err := GetKv().Put(&consulapi.KVPair{Key: key, Value: []byte(value)}, &consulapi.WriteOptions{Datacenter: dc, RelayFactor: 5})
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusForbidden, gin.H{"msg": "kv saved failed!"})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "0",
	})
}

func Del(c *gin.Context) {
	dc := CheckDc(c)
	key := c.Param("key")
	dir, _ := c.GetQuery("dir")
	if dir == "true" {
		key = key + "/"
	}
	_, errr := system.Client().KV().DeleteTree(key, &consulapi.WriteOptions{Datacenter: dc})
	if errr != nil {
		c.JSON(http.StatusOK, gin.H{"code": "1"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": "0"})
	}
}
