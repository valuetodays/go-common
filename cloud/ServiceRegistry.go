package cloud

type ServiceRegistry interface {
	Register(serviceInstance ServiceInstance) bool

	FindServices(serviceName string) []ServiceInstance

	Deregister()
}

