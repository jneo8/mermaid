package mermaid

import (
	"github.com/keepeye/logrus-filename"
	log "github.com/sirupsen/logrus"
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
