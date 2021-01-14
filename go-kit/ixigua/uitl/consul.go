package uitl

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"log"
)

var Client *api.Client
var Id string
var Name string
var Port int

func init() {
	config := api.DefaultConfig()
	config.Address = "http://192.168.26.10:8500"
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	Client = client
}
func SetNamePort(name string, port int) {
	Name = name
	Port = port
	Id = uuid.New().String()
}
func Register() {
	reg := api.AgentServiceRegistration{}
	reg.ID = Id
	reg.Name = Name
	reg.Address = "10.118.38.11"
	reg.Port = Port
	reg.Tags = []string{"primary"}

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
	err := Client.Agent().ServiceDeregister(Id)
	if err != nil {
		fmt.Println(err)
	}
}
