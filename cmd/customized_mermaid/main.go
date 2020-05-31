package main

import (
	"github.com/jneo8/mermaid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	worker, err := mermaid.NewMermaid(
		&cobra.Command{},
		viper.New(),
		log.New(),
		"",
	)
	if err != nil {
		log.Fatal(err)
	}

	runable := func(logger *log.Logger) error {
		logger.Info("Customized mermaid worker")
		return nil
	}
	initializers := []interface{}{}
	if err := worker.Execute(runable, initializers...); err != nil {
		log.Fatal(err)
	}
}
