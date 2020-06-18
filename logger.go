package mermaid

import (
	"github.com/keepeye/logrus-filename"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// NewLogger return default logger.
func NewLogger() *log.Logger {
	logger := log.New()
	logger.AddHook(filename.NewHook())
	logger.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000",
		FullTimestamp:   true,
	})
	return logger
}

// SetLoggerLevel setup logger's log level using viper "log_level" field .
func SetLoggerLevel(logger *log.Logger, cfg *viper.Viper) {
	level := cfg.GetString("log_level")
	switch level {
	case "debug":
		logger.SetLevel(log.DebugLevel)
	case "info":
		logger.SetLevel(log.InfoLevel)
	case "warn":
		logger.SetLevel(log.WarnLevel)
	case "error":
		logger.SetLevel(log.ErrorLevel)
	case "fatal":
		logger.SetLevel(log.FatalLevel)
	default:
		logger.SetLevel(log.InfoLevel)
	}
	logger.Infof("Log level: %s", logger.GetLevel().String())
}
