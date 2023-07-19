package consul

import (
	"fmt"
	"github.com/valuetodays/go-common/cloud"

	"github.com/hashicorp/consul/api"
)

type consulServiceRegistry struct {
	serviceInstances     map[string]map[string]cloud.ServiceInstance
	client               api.Client
	localServiceInstance cloud.ServiceInstance
}

func (c consulServiceRegistry) Register(serviceInstance cloud.ServiceInstance) bool {
	// 创建注册到consul的服务到
	registration := new(api.AgentServiceRegistration)
	registration.ID = serviceInstance.GetInstanceId()
	registration.Name = serviceInstance.GetServiceId()
	registration.Port = serviceInstance.GetPort()
	var tags []string
	if serviceInstance.IsSecure() {
		tags = append(tags, "secure=true")
	} else {
		tags = append(tags, "secure=false")
	}
	if serviceInstance.GetMetadata() != nil {
		var tags []string
		for key, value := range serviceInstance.GetMetadata() {
			tags = append(tags, key+"="+value)
		}
		registration.Tags = tags
	}
	registration.Tags = tags

	registration.Address = serviceInstance.GetHost()

	// 增加consul健康检查回调函数
	check := new(api.AgentServiceCheck)

	schema := "http"
	if serviceInstance.IsSecure() {
		schema = "https"
	}
	check.HTTP = fmt.Sprintf("%s://%s:%d/actuator/health", schema, registration.Address, registration.Port)
	check.Timeout = "5s"
	check.Interval = "5s"
	check.DeregisterCriticalServiceAfter = "20s" // 故障检查失败30s后 consul自动将注册服务删除
	registration.Check = check

	// 注册服务到consul
	err := c.client.Agent().ServiceRegister(registration)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if c.serviceInstances == nil {
		c.serviceInstances = map[string]map[string]cloud.ServiceInstance{}
	}

	services := c.serviceInstances[serviceInstance.GetServiceId()]

	if services == nil {
		services = map[string]cloud.ServiceInstance{}
	}

	services[serviceInstance.GetInstanceId()] = serviceInstance

	c.serviceInstances[serviceInstance.GetServiceId()] = services

	c.localServiceInstance = serviceInstance

	return true
}

// deregister a service
func (c consulServiceRegistry) Deregister() {
	if c.serviceInstances == nil {
		return
	}

	services := c.serviceInstances[c.localServiceInstance.GetServiceId()]

	if services == nil {
		return
	}

	delete(services, c.localServiceInstance.GetInstanceId())

	if len(services) == 0 {
		delete(c.serviceInstances, c.localServiceInstance.GetServiceId())
	}

	_ = c.client.Agent().ServiceDeregister(c.localServiceInstance.GetInstanceId())

	c.localServiceInstance = nil
}

// new a consulServiceRegistry instance
// token is optional
func NewConsulServiceRegistry(consulIpAndPort string, token string) (*consulServiceRegistry, error) {
	config := api.DefaultConfig()
	config.Address = consulIpAndPort
	config.Token = token
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &consulServiceRegistry{client: *client}, nil
}

