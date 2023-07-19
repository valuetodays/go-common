package cloud

import "io"

type ServiceRegistry interface {
	Register(serviceInstance ServiceInstance) bool

	Deregister()

	FindServices(serviceName string) []ServiceInstance

	RequestApiByService(serviceName string, method string, path string, body io.Reader) (string, error)
}

