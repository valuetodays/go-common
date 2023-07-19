package cloud

type ServiceRegistry interface {
	Register(serviceInstance ServiceInstance) bool

	FindService(serviceName string) []ServiceInstance

	Deregister()
}

