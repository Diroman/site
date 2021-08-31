package config

import (
	"github.com/spf13/viper"
)

type ConfigStruct struct {
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		Database string
	}
	Site struct {
		Host string
		Port int
	}
	Mail struct {
		Host     string
		Port     string
		Email    string
		Password string
	}
}

var Config = &ConfigStruct{}

func ReadConfig() error {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	v.AddConfigPath("config/")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(Config); err != nil {
		panic(err)
	}

	return nil
}
