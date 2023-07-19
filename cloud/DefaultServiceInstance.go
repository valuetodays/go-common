package cloud

type DefaultServiceInstance struct {
	InstanceId string
	ServiceId  string
	Host       string
	Port       int
	Secure     bool
	Metadata   map[string]string
}

