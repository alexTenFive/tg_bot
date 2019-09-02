package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	TelegramToken string `yaml:"telegram_token"`
}


var (
	Cfg Config
	ConfigFile = "config.yaml"
)

func init() {
	Cfg.getConf()

}
func (c *Config) getConf() {
	yamlFile, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}

