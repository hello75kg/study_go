package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"net/http"
	"time"
)

func Register(address string, port int, name string, tags []string) error {
	cfg := api.DefaultConfig()
	cfg.Address = "192.168.0.249:8500"

	client, _ := api.NewClient(cfg)

	registration := new(api.AgentServiceRegistration)
	registration.ID = address
	registration.Name = name
	registration.Tags = tags
	registration.Port = port
	check := api.AgentServiceCheck{
		HTTP:                           "http://192.168.0.249:8021/health",
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "5s",
	}
	registration.Check = &check

	err := client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
	return nil
}

func main() {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	go r.Run(":8021")

	_ = Register("192.168.0.249", 8021, "user-web", []string{"wshop-web", "wang"})
	AllServices()

	time.Sleep(10 * time.Second)
}

func AllServices() {
	cfg := api.DefaultConfig()
	cfg.Address = "192.168.0.249:8500"
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	services, err := client.Agent().Services()
	if err != nil {
		panic(err)
	}

	for key, service := range services {
		fmt.Println(key)
		fmt.Println(service)
	}

}

func FilterServices() {
	cfg := api.DefaultConfig()
	cfg.Address = "192.168.0.249:8500"
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	filter, err := client.Agent().ServicesWithFilter(`Service=="user-web"`)
	if err != nil {
		panic(err)
	}

	for key, service := range filter {
		fmt.Println(key)
		fmt.Println(service)
	}

}
