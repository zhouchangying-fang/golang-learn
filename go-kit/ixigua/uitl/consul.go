package uitl

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
)

var Client *api.Client

func init() {
	config := api.DefaultConfig()
	config.Address = "http://192.168.26.10:8500"
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	Client = client
}
func Register() {
	reg := api.AgentServiceRegistration{}
	reg.ID = "userService"
	reg.Name = "userService"
	reg.Address = "10.118.38.11"
	reg.Port = 8080

	check := api.AgentServiceCheck{}
	check.Interval = "5s"
	check.HTTP = "http://10.118.38.11:8080/user/100"
	reg.Check = &check

	err := Client.Agent().ServiceRegister(&reg)
	if err != nil {
		log.Fatal(err)
	}
}
func DeRegister() {
	err := Client.Agent().ServiceDeregister("userService")
	if err != nil {
		fmt.Println(err)
	}
}
