package yaml

import (
	"fmt"
	"github.com/valuetodays/go-common/utils"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func LoadYamlFileAs(yamlPath string, respConfig any) (any, error) {
	curDir, _ := os.Getwd()
	appPath := utils.GetAppPath()
	fmt.Println("curDir=", curDir)
	fmt.Println("appPath=", appPath)

	var dirToUse = curDir
	if curDir == "/" {
		dirToUse = utils.GetAppPath()
	}
	fullPath := filepath.Join(dirToUse, yamlPath)
	fmt.Println("fullPath=", fullPath)
	dataBytes, err := os.ReadFile(fullPath)
	if err != nil {
		fmt.Printf("读取文件失败：: %v\n", err)
		return nil, err
	}
	fmt.Println("yaml 文件的内容:\n" + string(dataBytes))
	err = yaml.Unmarshal(dataBytes, &respConfig)
	if err != nil {
		fmt.Println("解析 yaml 文件失败：", err)
		return nil, err
	}

	fmt.Printf("config → %+v\n", respConfig) // config → {Mysql:{Url:127.0.0.1 Port:3306} Redis:{Host:127.0.0.1 Port:6379}}
	return respConfig, nil
}
