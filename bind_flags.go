package mermaid

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func BindFlags(cmd *cobra.Command, cfg *viper.Viper) error {

	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		// cfg.SetDefault(flag.Name, flag.Value.flag.Value.Type())
		fmt.Printf("%#v %v\n", flag, flag.Value.Type())
		var val interface{}

		switch flag.Value.Type() {
		case "bool":
			val, _ = cmd.Flags().GetBool(flag.Name)
		case "float32":
			val, _ = cmd.Flags().GetFloat32(flag.Name)
		case "float64":
			val, _ = cmd.Flags().GetFloat64(flag.Name)
		}
		cfg.SetDefault(flag.Name, val)
	})

	if err := cfg.BindPFlags(cmd.Flags()); err != nil {
		return err
	}

	for _, subCmd := range cmd.Commands() {
		if err := BindFlags(subCmd, cfg); err != nil {
			return err
		}
	}
	return nil
}
