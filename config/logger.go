package config

import "go.uber.org/zap"

func InitLogger() (logger *zap.Logger) {
	logger, _ = zap.NewDevelopment()
	defer logger.Sync()
	logger.Info("instantiation", zap.String("type", "logger"), zap.String("source", "zap"), zap.String("status", "done"))
	return logger
}
