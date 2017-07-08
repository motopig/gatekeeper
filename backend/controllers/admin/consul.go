package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/motopig/gatekeeper/gate/keeper/system"
)

func Services(c *gin.Context) {

	se, err := system.Client().Agent().Services()
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"msg": "service not found!"})
		c.Abort()
		return
	}
	dc := Dc()

	c.HTML(http.StatusOK, "service.tmpl", gin.H{
		"services":   se,
		"datacenter": dc,
	})
}

func Health(c *gin.Context) {
	service := c.Param("service")
	dc := c.Param("datacenter")
	var err error

	var se = []*consulapi.ServiceEntry{}
	se, _, err = system.HealthService(service, "", false, &consulapi.QueryOptions{Datacenter: dc})
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"msg": "service not found!"})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"health": se})
}

func Nodes(c *gin.Context) {
	dc := CheckDc(c)
	nodes, _, err := system.Client().Catalog().Nodes(&consulapi.QueryOptions{Datacenter: dc})
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"msg": "node check error!"})
		c.Abort()
		return
	}
	dcs := Dc()

	c.HTML(http.StatusOK, "nodes.tmpl", gin.H{
		"nodes":      nodes,
		"datacenter": dcs,
	})
}

func DataCenter(c *gin.Context) {
	dc := Dc()

	c.JSON(http.StatusOK, gin.H{"datacenter": dc})
}

func Dc() []string {

	dc, err := system.Client().Catalog().Datacenters()
	if err != nil {
		var out []string
		return out
	}
	return dc
}

func Dnode(c *gin.Context) {
	node := c.Param("node")
	dc := c.Param("datacenter")
	detail, _, _ := Node(node, dc)
	c.JSON(http.StatusOK, gin.H{"detail": detail})

}

func Node(node string, dc string) (*consulapi.CatalogNode, *consulapi.QueryMeta, error) {

	return system.Client().Catalog().Node(node, &consulapi.QueryOptions{Datacenter: dc})
}

func CheckDc(c *gin.Context) string {
	dc, err := c.Cookie("datacenter")
	if err != nil || dc == "" {
		c.JSON(http.StatusForbidden, gin.H{"msg": "datacenter not set!"})
		c.Abort()
		return ""
	}
	return dc
}
