package mermaid

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/dig"
	"testing"
)

type service struct {
	Logger *log.Logger
	Config *viper.Viper
	N1     int
}

type serviceParams struct {
	dig.In
	Logger *log.Logger  `name:"logger"`
	Config *viper.Viper `name:"config"`
	N1     int          `name:"n1"`
}

func newService(params serviceParams) *service {
	return &service{
		Logger: params.Logger,
		Config: params.Config,
		N1:     params.N1,
	}
}

func TestNewMermaidWorker(t *testing.T) {
	cmd := &cobra.Command{
		Use:   "testing",
		Short: "testing",
		Long:  "testing",
	}

	cmd.Flags().Int("n1", 100, "testing var n1")
	cmd.Flags().String("config", "dev", "config file path")
	cmd.Flags().String("config_name", "dev", "config file name")

	worker := New(
		cmd,
		viper.New(),
		NewLogger(),
		"test",
	)
	runable := func(
		service *service,
	) error {
		service.Logger.Info(service.Config.AllSettings())
		service.Logger.Info(service.N1)
		return nil
	}

	initializers := []interface{}{
		newService,
	}

	if err := worker.Run(
		runable,
		initializers...,
	); err != nil {
		t.Error(err)
	}
}
