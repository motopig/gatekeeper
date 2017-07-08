package system

import (
	"math/rand"

	consulapi "github.com/hashicorp/consul/api"
)

type ConsulService struct {
	Address string
	Port    int
}

var client *consulapi.Client

func NewConsul() {
	var err error
	config := consulapi.DefaultConfig()
	client, err = consulapi.NewClient(config)
	if err != nil {
		panic("consul connect failed")
	}
}

func Client() *consulapi.Client {
	return client
}

func ParseService(service string) (ConsulService, error) {

	var services map[string]*consulapi.AgentService
	var err error
	services, err = AllService()

	if err != nil {
		return ConsulService{}, err
	}

	namedService := services[service]

	// todo cluster dc
	dc := GetConfig("datacenter::no1")
	var se = []*consulapi.ServiceEntry{}
	//var hc = consulapi.HealthChecks{}

	se, _, err = HealthService(service, "", false, &consulapi.QueryOptions{Datacenter: dc})
	//Client().Health().Service(service, "", false, &consulapi.QueryOptions{Datacenter: dc})
	//hc, _, err = CheckService(service, &consulapi.QueryOptions{Datacenter: dc})

	for k, v := range se {
		if v.Checks[1].Status != "passing" {
			se = append(se[:k], se[k+1:]...)
		}
	}

	if err != nil {
		return ConsulService{}, err
	}

	var w int
	if len(se) == 1 {
		w = 1
	} else {
		w = len(se) - 1
	}

	return ConsulService{Address: se[rand.Intn(w)].Node.Address, Port: namedService.Port}, nil
}

func AllService() (map[string]*consulapi.AgentService, error) {
	return Client().Agent().Services()
}

func HealthService(service, tag string, passingOnly bool, q *consulapi.QueryOptions) ([]*consulapi.ServiceEntry, *consulapi.QueryMeta, error) {
	return Client().Health().Service(service, tag, passingOnly, q)
}

func CheckService(service string, q *consulapi.QueryOptions) (consulapi.HealthChecks, *consulapi.QueryMeta, error) {
	return Client().Health().Checks(service, q)
}
