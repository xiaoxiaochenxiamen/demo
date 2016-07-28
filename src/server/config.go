package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const (
	ConfigFile = "./config.json"
)

type ConfigFlag struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

var Config = loadConfigFile()

func loadConfigFile() *ConfigFlag {
	data, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(err)
	}
	config := &ConfigFlag{}
	err = json.Unmarshal(data, config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Config: %#v\n", config)
	return config
}

func (c ConfigFlag) GetServerAddress() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
