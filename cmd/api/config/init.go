package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type APIConfig struct {
	Viper       *viper.Viper
	HertzConfig HertzConfig
	EtcdConfig  EtcdConfig
}

type HertzConfig struct {
	Address string
}

type EtcdConfig struct {
	EtcdAddress        string
	ServiceAddress     string
	ServiceName        string
	UserServerName     string
	CommentServerName  string
	FavoriteServerName string
	FeedServerName     string
	MessageServerName  string
	PublishServerName  string
	RelationServerName string
}

// TODO 替换日志打印方式

func InitConfig(c *APIConfig, filename string) error {

	v := viper.New()
	c.Viper = v
	v.SetConfigName(filename)
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	v.AddConfigPath("./config")
	v.AddConfigPath("./../config")
	v.AddConfigPath("./../../config")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			fmt.Println("Config file not found")
			return err
		} else {
			// Config file was found but another error was produced
			fmt.Println("Config file was found but another error was produced")
			return err
		}
	}

	err := v.Unmarshal(&c)
	if err != nil {
		fmt.Println("Error when Unmarshal config")
		return err
	}

	return nil
}
