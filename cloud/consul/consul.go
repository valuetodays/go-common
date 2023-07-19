package consul

import (
	"crypto/rand"
	"fmt"
	"github.com/valuetodays/go-common/cloud"
	"io"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
	"strings"

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

// 查找指定服务
func (c consulServiceRegistry) FindServices(serviceName string) ([]cloud.ServiceInstance, error) {
	if c.serviceInstances == nil {
		return nil, nil
	}

	services, err := c.client.Agent().Services()
	if nil != err {
		fmt.Println("Services(): ", err)
		return nil, err
	}

	var targetServiceList []cloud.ServiceInstance
	for _, si := range services {
		//fmt.Printf("-> %s serviceInfo: %v", serviceId, serviceInfo)
		//fmt.Println("\n---> ", serviceId,
		//	"kind", serviceInfo.Kind,
		//	"id", serviceInfo.ID,
		//	"service", serviceInfo.Service,
		//	"address", serviceInfo.Address,
		//	"port", serviceInfo.Port,
		//)
		if serviceName == si.Service {
			var ci = cloud.DefaultServiceInstance{ServiceId: si.Service, Host: si.Address, Port: si.Port, Metadata: si.Meta}
			targetServiceList = append(targetServiceList, ci)
			//targetServiceList = append(targetServiceList, serviceInfo.Address+":"+strconv.Itoa(serviceInfo.Port))
		}
	}

	return targetServiceList, nil
}

// 调用指定服务的
func (c consulServiceRegistry) RequestApiByService(serviceName string, method string, path string, body io.Reader) (string, error) {
	services, err := c.FindServices(serviceName)
	if nil != err {
		return "services is nil", err
	}

	length := len(services)
	if length == 0 {
		return "no service named '" + serviceName + "'", err
	} else if length == 1 {
		// todo 先不考虑权重
		return RequestApiByServiceInstance(services[0], method, path, body);
	} else {
		// todo 先随机一个，后续可以使用轮询、权重等策略
		randomBigInt, _ := rand.Int(rand.Reader, big.NewInt(int64(length)))
		randomInt, _ := strconv.Atoi(randomBigInt.String()) // string转int
		return RequestApiByServiceInstance(services[randomInt], method, path, body);
	}
}

func RequestApiByServiceInstance(serviceInstance cloud.ServiceInstance, method string, path string, body io.Reader) (string, error) {
	var ipAndPort = serviceInstance.GetHost() + ":" + strconv.Itoa(serviceInstance.GetPort())
	return RequestApi(method, ipAndPort, path, body);
}

// 简单封装一个请求api的方法
func RequestApi(method string, ipAndPort string, path string, body io.Reader) (string, error) {
	// 1.如果没有http开头就给它加一个
	if !strings.HasPrefix(ipAndPort, "http://") && !strings.HasPrefix(ipAndPort, "https://") {
		ipAndPort = "http://" + ipAndPort
	}
	// 2. 新建一个request
	req, _ := http.NewRequest(method, ipAndPort + path, body)

	// 3. 新建httpclient，并且传入request
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	// 4. 获取请求结果
	buff, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(buff), nil
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

