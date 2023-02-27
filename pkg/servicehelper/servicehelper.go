package servicehelper

import (
	"errors"
	"learn-echo/config"

	log "github.com/sirupsen/logrus"
)

func SetError(scope, source, status, message string) error {
	config.Cfg.LoggerConf.WithFields(log.Fields{
		"scope":   scope,
		"source":  source,
		"status":  status,
		"message": message,
	}).Error("Service error")

	return errors.New(message)
}
