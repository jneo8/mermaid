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
	cfg.SetDefault("log_level", "Info")
	level := cfg.GetString("log_level")
	switch level {
	case "Debug":
		logger.SetLevel(log.DebugLevel)
	case "Info":
		logger.SetLevel(log.InfoLevel)
	case "Warn":
		logger.SetLevel(log.WarnLevel)
	case "Error":
		logger.SetLevel(log.ErrorLevel)
	case "Fatal":
		logger.SetLevel(log.FatalLevel)
	default:
		logger.SetLevel(log.InfoLevel)
	}
}
