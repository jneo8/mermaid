package mermaid

import (
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
	}
	cmd.Flags().Int("n1", 1, "testing var n1")
	cmd.Flags().String("config", "dev", "config file path")
	cmd.Flags().String("config_name", "dev", "config file name")

	worker := New(
		"ttt",
		dig.New(),
		viper.New(),
		NewLogger(),
		cmd,
	)
	worker.Logger.Infof("%#v", worker)
	worker.Logger.Infof("%#v", worker.Config.AllSettings())
	worker.Logger.Infof("%#v", worker.Config.GetString("abc"))
}
