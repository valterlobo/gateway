package config

import (
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
)

type Config struct {
	configViper *viper.Viper
}

func NewConfig() *Config {

	configViper := viper.New()
	configViper.SetConfigFile("./config/env.json")
	newConfig := Config{configViper}

	if err := configViper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	return &newConfig
}

func (config *Config) GetValue(key string) string {

	strValue := config.configViper.GetString(key)
	return strValue
}

func (config *Config) GetValues(key string) map[string]string {

	return config.configViper.GetStringMapString(key)
}

func  ReadFileType(file string) []byte {

	content, err := ioutil.ReadFile("config/" + file)
	if err != nil {
		log.Fatal(err)
	}
	return content
}