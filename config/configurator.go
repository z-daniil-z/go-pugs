package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Ip   string `json:"ip"`
	Port string `json:"port"`
}

func GetConfig() (*Config, error) {
	ret := &Config{}
	if _, in := os.LookupEnv("DEBUG"); in {
		if err := ret.DebugConfig(); err != nil {
			return nil, err
		}
	} else {
		ret.ProdConfig()
	}
	return ret, nil
}

func (conf *Config) DebugConfig() error {
	var path string
	var in bool
	if path, in = os.LookupEnv("PATH"); !in {
		log.Println("No path set. Default path is config/config.json")
		path = "config/config.json"
	}
	if data, err := ioutil.ReadFile(path); err != nil {
		return err
	} else {
		if err := json.Unmarshal(data, conf); err != nil {
			return err
		}
	}
	return nil
}

func (conf *Config) ProdConfig() {
	if ip, in := os.LookupEnv("IP"); in {
		conf.Ip = ip
	} else {
		log.Println("No IP set. Default IP is 127.0.0.1")
		conf.Ip = "127.0.0.1"
	}
	if port, in := os.LookupEnv("PORT"); in {
		conf.Port = port
	} else {
		log.Println("No PORT set. Default PORT is 8000")
		conf.Port = "8000"
	}
}
