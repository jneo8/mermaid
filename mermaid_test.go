package mermaid

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/dig"
	"testing"
)

func TestNewMermaidWorker(t *testing.T) {
	cmd := &cobra.Command{
		Use:   "testing",
		Short: "testing",
		Long:  "testing",
		RunE: func(cmd *cobra.Command, args []string) error {

			type service struct {
				N1 int
			}

			type serviceParams struct {
				dig.In
				N1 int `name:"n1"`
			}

			newService := func(params serviceParams) *service {
				return &service{
					N1: params.N1,
				}
			}

			initializers := []interface{}{
				func() string {
					return "123"
				},
				newService,
			}

			runable := func(
				service *service,
				logger *log.Logger,
				config *viper.Viper,
			) error {
				logger.Infof("%#v\n", config)
				logger.Info(config.AllSettings())
				logger.Info(service.N1)
				return nil
			}

			return Run(cmd, runable, initializers...)
		},
	}

	cmd.Flags().Int("n1", 100, "testing var n1")

	if err := cmd.Execute(); err != nil {
		t.Error(err)
	}
}
