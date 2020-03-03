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
	testCases := []struct {
		name string
		want map[string]interface{}
		cmd  *cobra.Command
	}{
		{
			name: "bool",
			cmd: func() *cobra.Command {
				f := setupFlagSet()
				f.Bool("v1", true, "v1")

				cmd := &cobra.Command{}
				cmd.Flags().AddFlagSet(f)
				return cmd
			}(),
			want: map[string]interface{}{"v1": true},
		},
		{
			name: "boolSlice",
			cmd: func() *cobra.Command {
				f := setupFlagSet()
				f.BoolSlice("v1", []bool{true, false, true}, "v1")

				cmd := &cobra.Command{}
				cmd.Flags().AddFlagSet(f)
				return cmd
			}(),
			want: map[string]interface{}{"v1": []bool{true, false, true}},
		},
		{
			name: "bytesHex",
			cmd: func() *cobra.Command {
				f := setupFlagSet()
				f.BytesHex("v1", []byte("testing"), "v1")

				cmd := &cobra.Command{}
				cmd.Flags().AddFlagSet(f)
				return cmd
			}(),
			want: map[string]interface{}{"v1": []byte("testing")},
		},
		{
			name: "count",
			cmd: func() *cobra.Command {
				f := setupFlagSet()
				f.CountP("v1", "v", "a counter")

				f.Parse([]string{"-vvv"})

				cmd := &cobra.Command{}
				cmd.Flags().AddFlagSet(f)
				return cmd
			}(),
			want: map[string]interface{}{"v1": 3},
		},
		{
			name: "float64",
			cmd: func() *cobra.Command {
				f := setupFlagSet()
				f.Float64("v1", 1.8, "b1")

				cmd := &cobra.Command{}
				cmd.Flags().AddFlagSet(f)
				return cmd
			}(),
			want: map[string]interface{}{"v1": 1.8},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			cfg := viper.New()
			cfg.SetTypeByDefaultValue(true)
			BindFlags(
				tt.cmd,
				cfg,
			)
			if !reflect.DeepEqual(cfg.AllSettings(), tt.want) {
				t.Errorf("BindFlags() bind flag: %#v want: %#v", cfg.AllSettings(), tt.want)
			}
		})
	}
}
