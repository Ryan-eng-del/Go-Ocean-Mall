package internal

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

type ConsulConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

func Req(host, name, id string, port int, tags []string) error {
	config := api.DefaultConfig()

	h := ViperConf.ConsulConfig.Host
	p := ViperConf.ConsulConfig.Port

	addr := fmt.Sprintf("%s:%d", h, p)
	config.Address = addr

	client, err := api.NewClient(config)

	if err != nil {
		return err
	}

	agentServiceRegistration := new(api.AgentServiceRegistration)
	agentServiceRegistration.Address = config.Address
	agentServiceRegistration.Port = port
	agentServiceRegistration.ID = id
	agentServiceRegistration.Name = name
	agentServiceRegistration.Tags = tags
	serverAddr := fmt.Sprintf("http://%s:%d/health", host, port)

	check := api.AgentServiceCheck{
		HTTP:                           serverAddr,
		Timeout:                        "3s",
		Interval:                       "3s",
		DeregisterCriticalServiceAfter: "5s",
	}

	agentServiceRegistration.Check = &check

	return client.Agent().ServiceRegister(agentServiceRegistration)
}

func GetServiceList() error {
	config := api.DefaultConfig()

	h := ViperConf.ConsulConfig.Host
	p := ViperConf.ConsulConfig.Port

	addr := fmt.Sprintf("%s:%d", h, p)
	config.Address = addr

	client, err := api.NewClient(config)

	if err != nil {
		return err
	}

	services, err := client.Agent().Services()

	if err != nil {
		return err
	}

	for k, v := range services {
		fmt.Println(k)
		fmt.Println(v)
		fmt.Println("---------------")
	}

	return nil
}

func FilterService() error {
	config := api.DefaultConfig()

	h := ViperConf.ConsulConfig.Host
	p := ViperConf.ConsulConfig.Port

	addr := fmt.Sprintf("%s:%d", h, p)
	config.Address = addr

	client, err := api.NewClient(config)

	if err != nil {
		return err
	}

	services, err := client.Agent().ServicesWithFilter("Service==account_web")

	if err != nil {
		return err
	}

	for k, v := range services {
		fmt.Println(k)
		fmt.Println(v)
		fmt.Println("---------------")
	}
	return nil
}
