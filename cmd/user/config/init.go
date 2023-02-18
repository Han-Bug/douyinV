package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type UserConfig struct {
	Viper      *viper.Viper
	DBConfig   DataBaseConfig
	EtcdConfig EtcdConfig
}

type DataBaseConfig struct {
	MySql *MySqlConfig
}

type EtcdConfig struct {
	EtcdAddress    string
	ServiceAddress string
	ServiceName    string
}

type MySqlConfig struct {
	Username string
	Password string
	Address  string
	DBName   string
	Options  string
}

func (c *MySqlConfig) GetDSN() string {
	builder := strings.Builder{}
	builder.WriteString(c.Username)
	builder.WriteString(":")
	builder.WriteString(c.Password)
	builder.WriteString("@tcp(")
	builder.WriteString(c.Address)
	builder.WriteString(")/")
	builder.WriteString(c.DBName)
	if len(c.Options) != 0 {
		builder.WriteString("?")
		builder.WriteString(c.Options)
	}
	// 	dsn = "root:password@tcp(47.108.237.99:3311)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	return builder.String()
}

// TODO 替换日志打印方式

func InitConfig(c *UserConfig, filename string) error {
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
