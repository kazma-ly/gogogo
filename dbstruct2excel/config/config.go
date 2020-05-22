package config

import (
	"io/ioutil"
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

type (
	// Config project config
	Config struct {
		MyMySQLConfig MySQLConfig `yaml:"mysql"`
	}

	// MySQLConfig mysql config
	MySQLConfig struct {
		Host     string `yaml:"host"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Port     string `yaml:"port"`
		DBName   string `yaml:"dbname"`
	}
)

var (
	config *Config
	lock   sync.RWMutex
)

// GetConfig get yaml config
func GetConfig() *Config {
	lock.Lock()
	defer lock.Unlock()

	if config == nil {
		configFile, err := os.OpenFile("./config/config.yml", os.O_RDONLY, 0755)
		if err != nil {
			log.Panicln(err)
		}

		configByte, err := ioutil.ReadAll(configFile)
		if err != nil {
			log.Panicln(err)
		}

		_config := &Config{}
		if yaml.Unmarshal(configByte, _config) != nil {
			log.Panicln(err)
		}

		config = _config
	}

	return config
}
