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
	_ = LoadYamlFileAs("../resources/application-dev.yaml", &config)
	//fmt.Printf("config2 → %+v\n", config2)
	fmt.Println("====")
	fmt.Println("config:", config)
	fmt.Println("config:", config.ConsulIpAndPort)
	//fmt.Println("config2:", config2.ConsulIpAndPort)
}

func TestObjToYaml(t *testing.T) {
	config := ApplicationConfig{}
	config.ConsulIpAndPort = "aaaaa:8500"
	config.UseDebug = true
	config.Metadata = map[string]string{
		"KEY1":"Value1",
		"KEY2":"Value2",
		"KEY3":"Value3",
	}
	bytes, _ := yaml.Marshal(config)
	fmt.Print(string(bytes))
}
