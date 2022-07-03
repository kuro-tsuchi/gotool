package config

import (
	"log"

	"github.com/spf13/viper"
)

var Conf *Config

type Config struct {
	Server ServerConfig
	Mysql  MysqlConfig
}
type ServerConfig struct {
	Port string
}

type MysqlConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func ViperInit() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Conf)
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Conf)
	log.Printf("%+v", Conf)

}
