package middlewares

import (
	"time"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func makeLogEntry(c echo.Context) *log.Entry {
	logger := log.New()
	logger.Formatter = &log.TextFormatter{
		ForceColors:     true,
		TimestampFormat: "2006/01/02 - 15:04:05",
		FullTimestamp:   true,
	}

	if c == nil {
		return logger.WithFields(log.Fields{
			"at": time.Now().Format("2006-01-02 15:04:05"),
		})
	}

	return logger.WithFields(log.Fields{
		// "at":     time.Now().Format("2006-01-02 15:04:05"),
		"method": c.Request().Method,
		"path":   c.Path(),
		"ip":     c.Request().RemoteAddr,
		"status": c.Response().Status,
	})
}

func LogMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		makeLogEntry(c).Info("incoming request")
		return next(c)
	}
}
