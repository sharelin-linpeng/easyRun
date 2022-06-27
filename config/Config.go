package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

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
	configStr, _ := json.Marshal(config)
	println("loadConfig:" + string(configStr))
}

func loadConfigApp(path string) *Config {
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

	return &config
}

var CONFIG = loadConfigApp(os.Args[1])
