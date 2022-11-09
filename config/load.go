package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func init() {
	err := readConfig()
	if err != nil {
		log.Println(fmt.Sprintf("Load config failed with err:%s", err.Error()))
		os.Exit(1)
	}
}

type config struct {
	Http struct {
		IsEnable bool   `yaml:"isEnable" json:"isEnable"`
		Address  string `yaml:"address" json:"address"`
		Port     int    `yaml:"port" json:"port"`
	} `yaml:"http" json:"http"`

	Bot struct {
		Address string `yaml:"address" json:"address"`
		Port    int    `yaml:"port" json:"port"`
		QQ      int    `yaml:"qq"`
	} `yaml:"bot"`
}

var Srv = &config{}

// 读取配置
func readConfig() error {
	configFile, err := os.ReadFile("./config.yaml")
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(configFile, &Srv)
	if err != nil {
		return err
	}

	return nil
}
