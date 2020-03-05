package mermaid

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// BindFlags will bind cobra flagset to viper as default value.
func BindFlags(cmd *cobra.Command, cfg *viper.Viper) error {
	// Need to set viper.typeByDefaultValue to true to get the value with correct type.
	cfg.SetTypeByDefaultValue(true)
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		// cfg.SetDefault(flag.Name, flag.Value.flag.Value.Type())
		var val interface{}

		switch flag.Value.Type() {
		case "bool":
			val, _ = cmd.Flags().GetBool(flag.Name)
		case "boolSlice":
			val, _ = cmd.Flags().GetBoolSlice(flag.Name)
		case "bytesHex":
			val, _ = cmd.Flags().GetBytesHex(flag.Name)
		case "count":
			val, _ = cmd.Flags().GetCount(flag.Name)
		case "duration":
			val, _ = cmd.Flags().GetDuration(flag.Name)
		case "durationSlice":
			val, _ = cmd.Flags().GetDurationSlice(flag.Name)
		case "float32":
			val, _ = cmd.Flags().GetFloat32(flag.Name)
		case "float32Slice":
			val, _ = cmd.Flags().GetFloat32Slice(flag.Name)
		case "float64":
			val, _ = cmd.Flags().GetFloat64(flag.Name)
		case "float64Slice":
			val, _ = cmd.Flags().GetFloat64Slice(flag.Name)
		case "int":
			val, _ = cmd.Flags().GetInt(flag.Name)
		case "int8":
			val, _ = cmd.Flags().GetInt8(flag.Name)
		case "int16":
			val, _ = cmd.Flags().GetInt16(flag.Name)
		case "int32":
			val, _ = cmd.Flags().GetInt32(flag.Name)
		case "int64":
			val, _ = cmd.Flags().GetInt64(flag.Name)
		case "int32Slice":
			val, _ = cmd.Flags().GetInt32Slice(flag.Name)
		case "int64Slice":
			val, _ = cmd.Flags().GetInt64Slice(flag.Name)
		case "ip":
			val, _ = cmd.Flags().GetIP(flag.Name)
		case "ipSlice":
			val, _ = cmd.Flags().GetIPSlice(flag.Name)
		case "ipMask":
			val, _ = cmd.Flags().GetIPv4Mask(flag.Name)
		case "ipNet":
			val, _ = cmd.Flags().GetIPNet(flag.Name)
		case "string":
			val, _ = cmd.Flags().GetString(flag.Name)
		case "stringSlice":
			val, _ = cmd.Flags().GetStringSlice(flag.Name)
		case "stringArray":
			val, _ = cmd.Flags().GetStringArray(flag.Name)
		case "stringToInt":
			val, _ = cmd.Flags().GetStringToInt(flag.Name)
		case "stringToInt64":
			val, _ = cmd.Flags().GetStringToInt64(flag.Name)
		case "stringToString":
			val, _ = cmd.Flags().GetStringToString(flag.Name)
		case "uint":
			val, _ = cmd.Flags().GetUint(flag.Name)
		case "uint8":
			val, _ = cmd.Flags().GetUint8(flag.Name)
		case "uint16":
			val, _ = cmd.Flags().GetUint16(flag.Name)
		case "uint32":
			val, _ = cmd.Flags().GetUint32(flag.Name)
		case "uint64":
			val, _ = cmd.Flags().GetUint64(flag.Name)
		case "uintSlice":
			val, _ = cmd.Flags().GetUintSlice(flag.Name)
		}
		cfg.SetDefault(flag.Name, val)
	})

	return cfg.BindPFlags(cmd.Flags())
}
