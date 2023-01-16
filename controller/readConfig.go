package controller

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

//ReadConfig 读取外部配置文件
func ReadConfig(input string) (EnvConf, error) {
	yamlFile, err := ioutil.ReadFile(input + "/config.yaml")
	if err != nil {
		log.Println("open config file error:", err)
	}
	var confDemo EnvConf
	err = yaml.Unmarshal(yamlFile, &confDemo)
	if err != nil {
		log.Println("Unmarshal Config Error", err)
	}

	fmt.Println(confDemo)

	return confDemo, nil

}
