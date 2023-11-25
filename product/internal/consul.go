package internal

import (
	"fmt"
	"strconv"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
)

type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

type ProductSrvConf struct {
	Host    string   `mapstructure:"host" json:"host"`
	Port    string   `mapstructure:"port" json:"port"`
	SrvName string   `mapstructure:"srvName" json:"srvName"`
	Tags    []string `mapstructure:"tags" json:"tags"`
}

type ProductWebConf struct {
	Host    string   `mapstructure:"host" json:"host"`
	Port    string   `mapstructure:"port" json:"port"`
	SrvName string   `mapstructure:"srvName" json:"srvName"`
	Tags    []string `mapstructure:"tags" json:"tags"`
}

func Req(host, name, id string, port int, tags []string, isGrpc bool) error {
	config := api.DefaultConfig()

	h := ViperConf.ConsulConfig.Host
	p := ViperConf.ConsulConfig.Port

	addr := fmt.Sprintf("%s:%d", h, p)
	fmt.Println(addr, "consul addr")
	config.Address = addr

	client, err := api.NewClient(config)

	if err != nil {
		return err
	}
	randomUUID := uuid.New().String()

	agentServiceRegistration := new(api.AgentServiceRegistration)
	agentServiceRegistration.Address = fmt.Sprintf("%s", host)
	agentServiceRegistration.Port = port
	agentServiceRegistration.ID = "port-" + strconv.Itoa(port) + "-" + randomUUID[0:4]
	agentServiceRegistration.Name = name
	agentServiceRegistration.Tags = tags
	var serverAddr string

	if !isGrpc {
		serverAddr = fmt.Sprintf("http://%s:%d/health", host, port)

	} else {
		serverAddr = fmt.Sprintf("%s:%d", host, port)
	}
	fmt.Println(serverAddr, "health addr")
	var check api.AgentServiceCheck

	if !isGrpc {
		check = api.AgentServiceCheck{
			HTTP:                           serverAddr,
			Timeout:                        "3s",
			Interval:                       "3s",
			DeregisterCriticalServiceAfter: "10s",
		}
	} else {
		check = api.AgentServiceCheck{
			GRPC:                           serverAddr,
			Timeout:                        "3s",
			Interval:                       "3s",
			DeregisterCriticalServiceAfter: "10s",
		}
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

func FilterService(serviceName string) (string, error) {
	config := api.DefaultConfig()

	h := ViperConf.ConsulConfig.Host
	p := ViperConf.ConsulConfig.Port

	addr := fmt.Sprintf("%s:%d", h, p)
	config.Address = addr

	client, err := api.NewClient(config)

	if err != nil {
		return "", err
	}

	services, err := client.Agent().ServicesWithFilter(fmt.Sprintf("Service==%s", serviceName))

	if err != nil {
		return "", err
	}
	var addrs string
	for _, v := range services {
		addrs = fmt.Sprintf("%s", v.Address)
	}
	return addrs, nil
}
