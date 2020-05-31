package mermaid

import (
	"reflect"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/dig"
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

func TestNewMermaid(t *testing.T) {
	type args struct {
		cmd       *cobra.Command
		config    *viper.Viper
		logger    *log.Logger
		envPrefix string
	}
	tests := []struct {
		name      string
		args      args
		want      *Mermaid
		wantErr   bool
		deepEqual bool
	}{
		// TODO: Add test cases.
		{
			name: "Basic test",
			args: args{
				cmd:       &cobra.Command{},
				config:    viper.New(),
				logger:    log.New(),
				envPrefix: "",
			},
			want: func() *Mermaid {
				worker := &Mermaid{
					Container: dig.New(),
					CMD:       &cobra.Command{},
					Config:    viper.New(),
					Logger:    log.New(),
					ENVPrefix: "",
				}
				if err := worker.Bind(); err != nil {
					t.Error(err)
				}
				return worker
			}(),
			wantErr:   false,
			deepEqual: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMermaid(tt.args.cmd, tt.args.config, tt.args.logger, tt.args.envPrefix)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMermaid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.deepEqual {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("NewMermaid() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
