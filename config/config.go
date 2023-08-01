package config

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type HttpConf struct {
	Host string
	Port int
}

type EmailConf struct {
	Email    string
	Password string
	Host     string
	Port     int
}

type Config struct {
	DbConf     DbConf
	HttpConf   HttpConf
	EmailConf  EmailConf
	LoggerConf *log.Logger
}

var Cfg *Config

func GetConfig() {
	// main viper config
	// viper.SetConfigName("config")
	// viper.SetConfigType("yml")
	// viper.AddConfigPath(".")
	// viper.AutomaticEnv()
	viper.SetConfigFile("D:/Belajar/BE/learnEcho2/config.yml")
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
	if err := viper.Unmarshal(&Cfg); err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
		panic(err)

	}

	// initialize logger
	initLogger()

	// done
	Cfg.LoggerConf.WithFields(log.Fields{
		"source":  "config",
		"status":  "done",
		"message": "config is loaded successfully",
	}).Info("loading config")

}

func initLogger() {
	logger := log.New()
	logger.Formatter = &log.TextFormatter{
		ForceColors:     true,
		TimestampFormat: "2006/01/02 - 15:04:05",
		FullTimestamp:   true,
	}
	Cfg.LoggerConf = logger
}
