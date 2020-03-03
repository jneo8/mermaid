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
	cmd.Flags().Int8("n2", 1, "testing var n2")
	cmd.Flags().Float64("f1", 1.23456, "testing var f1")
	cmd.Flags().String("config", "dev", "config file path")
	cmd.Flags().String("config_name", "dev", "config file name")

	_ = New(
		"ttt",
		dig.New(),
		viper.New(),
		NewLogger(),
		cmd,
	)
}
