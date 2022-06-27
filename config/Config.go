package config

import (
	"fmt"
	"io/ioutil"

	"github.com/sharelin-linpeng/easyRun/jsonutil"
	"gopkg.in/yaml.v2"
)

type Config_App struct {
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	StaticPath string `yaml:"staticPath"`
	MySql      string `yaml:"mysql"`
}

type Config_Git struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
}

type Config struct {
	App Config_App `yaml:"app"`
	Git Config_Git `yaml:"git"`
}

func (config Config) showConfig() {
	println("loadConfig:" + jsonutil.Obj2Json(config))
}

var CONFIG *Config

func LoadConfigApp(path string) {
	var config Config
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("failed to read yaml file : %s\n", path)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Printf("failed to unmarshal : %s\n", path)
	}
	config.showConfig()
	CONFIG = &config
}
