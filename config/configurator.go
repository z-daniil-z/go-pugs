package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Ip       string   `json:"ip"`
	Port     string   `json:"port"`
	DataBase DataBase `json:"dataBase"`
}

type DataBase struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Ip       string `json:"ip"`
	Port     string `json:"port"`
	DbName   string `json:"dbName"`
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
	if path, in = os.LookupEnv("CONF_PATH"); !in {
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
	conf.Ip = lookArg("IP", "127.0.0.1")
	conf.Port = lookArg("PORT", "8000")
	conf.DataBase.User = lookArg("DB_USER", "postgres")
	conf.DataBase.Password = lookArg("DB_PASSWORD", "postgres")
	conf.DataBase.Ip = lookArg("DB_IP", "127.0.0.1")
	conf.DataBase.Port = lookArg("DB_PORT", "5432")
	conf.DataBase.DbName = lookArg("DB_NAME", "Pugs")
}

func lookArg(arg, def string) string {
	if val, in := os.LookupEnv(arg); in {
		return val
	} else {
		log.Println(fmt.Sprintf("No %s set. Default %s is %s", arg, arg, def))
		return def
	}
}
