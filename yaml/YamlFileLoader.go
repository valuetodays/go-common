package yaml

import (
	"fmt"
	"github.com/valuetodays/go-common/utils"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// usage:
//
//	var config = ApplicationConfig{}
//	_ = LoadYamlFileAs(&config, "../resources/application-dev.yaml")
//	fmt.Println("config:", config)
func LoadYamlFileAs(respConfig any, yamlPath string) error {
	byteArray, err := LoadYamlFileAsByteArray(yamlPath)
	if nil != err {
		return err
	}
	fmt.Println("yaml 文件的内容:\n" + string(byteArray))
	err = yaml.Unmarshal(byteArray, respConfig)
	if err != nil {
		fmt.Println("解析 yaml 文件失败：", err)
		return err
	}

	//fmt.Printf("config → %+v\n", respConfig) // config → {Mysql:{Url:127.0.0.1 Port:3306} Redis:{Host:127.0.0.1 Port:6379}}
	return nil
}

func LoadYamlFileAsByteArray(yamlPath string) ([]byte, error) {
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
	//fmt.Println("yaml 文件的内容:\n" + string(dataBytes))

	return dataBytes, nil
}

// usage:
//
//	var config = ApplicationConfig{}
//	_ = LoadYamlFilesAs(&config, "../resources/application-dev.yaml", "../resources/application-dev2.yaml")
//	fmt.Println("config:", config)
func LoadYamlFilesAs(respConfig any, yamlPaths ...string) error {
	//var sepBytes = []byte("---\n")
	var allFileAsByteArray []byte
	for _, yamlPath := range yamlPaths {
		array, err := LoadYamlFileAsByteArray(yamlPath)
		if nil != err {
			return err
		}
		allFileAsByteArray = append(allFileAsByteArray, array...)
		//allFileAsByteArray = append(allFileAsByteArray, sepBytes...)
	}
	err := yaml.Unmarshal(allFileAsByteArray, respConfig)
	if err != nil {
		fmt.Println("解析 yaml 文件(多个)失败：", err)
		return err
	}

	//fmt.Printf("config → %+v\n", respConfig) // config → {Mysql:{Url:127.0.0.1 Port:3306} Redis:{Host:127.0.0.1 Port:6379}}
	return nil
}
