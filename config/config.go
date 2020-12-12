package config

import (
	"github.com/BurntSushi/toml"
)

const configPath = "./config.toml"

//Config conf
type Config struct {
	Port     string
	DataPath string
}

//Conf path conf
var conf Config

//Init load config
func Init() error {
	if _, err := toml.DecodeFile(configPath, &conf); err != nil {
		return err
	}

	return nil
}

//GetConfig return config
func GetConfig() Config {
	return conf
}
