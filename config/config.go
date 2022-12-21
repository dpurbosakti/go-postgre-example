package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type HttpConf struct {
	Host string
	Port int
}

type Config struct {
	DbConf   DbConf
	HttpConf HttpConf
}

func GetConfig() (c *Config) {
	// main viper config
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	// default values
	viper.SetDefault("FullName", "mokotest")
	viper.SetDefault("Version", "0.0.1")
	viper.SetDefault("HttpConf.Host", "127.0.0.1")
	viper.SetDefault("HttpConf.Port", "8000")

	// read the file
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		panic(err)
	}

	// map to app
	if err := viper.Unmarshal(&c); err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
		panic(err)

	}

	// done
	fmt.Println("Config is loaded successfully")
	return c
}
