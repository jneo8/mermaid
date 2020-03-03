package mermaid

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"reflect"
	"testing"
)

func setupFlagSet() *pflag.FlagSet {
	return pflag.NewFlagSet("test", pflag.ContinueOnError)
}

func TestBindFlags(t *testing.T) {
	tests := []struct {
		name           string
		getFlagsetFunc func() *pflag.FlagSet
		want           map[string]interface{}
	}{
		{
			name: "bool",
			getFlagsetFunc: func() *pflag.FlagSet {
				f := setupFlagSet()
				f.Bool("bool1", true, "bool1")
				return f
			},
			want: map[string]interface{}{"bool1": true},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			cfg := viper.New()
			cmd := &cobra.Command{}
			cmd.Flags().AddFlagSet(tt.getFlagsetFunc())
			BindFlags(
				cmd,
				cfg,
			)
			if !reflect.DeepEqual(cfg.AllSettings(), tt.want) {
				t.Error(tt)
			}
		})
	}
}
