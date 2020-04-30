package config

import (
	"github.com/ghodss/yaml"
	"io/ioutil"
)

/**
 * @Author: eggsy
 * @Description:
 * @File:  config
 * @Version: 1.0.0
 * @Date: 12/7/19 1:48 下午
 */
var Conf *Config

type Config struct {
	Port       string `yaml:"port"`
	OrgName    string `yaml:"orgName"`
	UserName   string `yaml:"userName"`
	SdkCfgPath string `yaml:"sdkCfgPath"`
	Level      string `yaml:"level"`
	LogPath    string `yaml:"logPath"`
}

func InitConfig(filename string) {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("failed to read app.yaml")
	}
	Conf = new(Config)
	err = yaml.Unmarshal(yamlFile, Conf)
	if err != nil {
		panic("failed to unmarshal app.yaml")
	}
}
