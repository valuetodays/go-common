package yaml

import (
	"fmt"
	"testing"

	"gopkg.in/yaml.v3"
)

type ApplicationConfig struct {
	// 对应到yaml中的字段名是全部变成小写
	UseDebug        bool              `json:"usedebug"`
	ConsulIpAndPort string            `json:"consulipandport"`
	Metadata        map[string]string `json:"metadata"`
}


func TestLoadApplicationConf(t *testing.T) {
	var config = ApplicationConfig{}
	config2, _ := LoadYamlFileAs("../resources/application-dev.yaml", config)
	fmt.Printf("config2 → %+v\n", config2)
	fmt.Print("config2:", config2)
}

func TestObjToYaml(t *testing.T) {
	config := ApplicationConfig{}
	config.ConsulIpAndPort = "aaaaa:8500"
	config.UseDebug = true
	bytes, _ := yaml.Marshal(config)
	fmt.Print(string(bytes))
}
