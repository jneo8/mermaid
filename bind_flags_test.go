package mermaid

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net"
	"reflect"
	"testing"
	"time"
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
			name: "duration",
			cmd: func() *cobra.Command {
				f := setupFlagSet()
				f.Duration("v1", time.Duration(100), "v1")

				cmd := &cobra.Command{}
				cmd.Flags().AddFlagSet(f)
				return cmd
			}(),
			want: map[string]interface{}{"v1": time.Duration(100)},
		},
		{
			name: "durationSlice",
			cmd: func() *cobra.Command {
				f := setupFlagSet()
				f.DurationSlice(
					"v1",
					[]time.Duration{
						time.Duration(100),
						time.Duration(200),
					},
					"v1",
				)

				cmd := &cobra.Command{}
				cmd.Flags().AddFlagSet(f)
				return cmd
			}(),
			want: map[string]interface{}{
				"v1": []time.Duration{
					time.Duration(100),
					time.Duration(200),
				},
			},
		},
		{
			name: "float32",
			cmd: func() *cobra.Command {
				f := setupFlagSet()
				f.Float32("v1", 1.1, "v1")

				cmd := &cobra.Command{}
				cmd.Flags().AddFlagSet(f)
				return cmd
			}(),
			// viper will turn float32 to float64 when Get value.
			// See https://godoc.org/github.com/spf13/viper#Viper.Get
			want: map[string]interface{}{"v1": float64(float32(1.1))},
		},
		{
			name: "float32Slice",
			cmd: func() *cobra.Command {
				f := setupFlagSet()
				f.Float32Slice("v1", []float32{1.1, 2.2, 3.3}, "v1")

				cmd := &cobra.Command{}
				cmd.Flags().AddFlagSet(f)
				return cmd
			}(),
			want: map[string]interface{}{"v1": []float32{1.1, 2.2, 3.3}},
		},
		{
			name: "float64",
			cmd: func() *cobra.Command {
				f := setupFlagSet()
				f.Float64("v1", 1.8, "v1")

				cmd := &cobra.Command{}
				cmd.Flags().AddFlagSet(f)
				return cmd
			}(),
			want: map[string]interface{}{"v1": float64(1.8)},
		},
		{
			name: "float64Slice",
			cmd: func() *cobra.Command {
				f := setupFlagSet()
				f.Float64Slice("v1", []float64{1.1, 2.2, 3.3}, "v1")

				cmd := &cobra.Command{}
				cmd.Flags().AddFlagSet(f)
				return cmd
			}(),
			want: map[string]interface{}{"v1": []float64{1.1, 2.2, 3.3}},
		},
		{
			name: "int",
			cmd: func() *cobra.Command {
				f := setupFlagSet()
				f.Int("v1", 1, "v1")

				cmd := &cobra.Command{}
				cmd.Flags().AddFlagSet(f)
				return cmd
			}(),
			want: map[string]interface{}{"v1": 1},
		},
		// For int8, int16, int32. viper now default return type is int.
		// See https://godoc.org/github.com/spf13/viper#Viper.Get
		{
			name: "int8",
			cmd: func() *cobra.Command {
				f := setupFlagSet()
				f.Int8("v1", 1, "v1")

				cmd := &cobra.Command{}
				cmd.Flags().AddFlagSet(f)
				return cmd
			}(),
			want: map[string]interface{}{"v1": int(1)},
		},
		{
			name: "int16",
			cmd: func() *cobra.Command {
				f := setupFlagSet()
				f.Int16("v1", 1, "v1")

				cmd := &cobra.Command{}
				cmd.Flags().AddFlagSet(f)
				return cmd
			}(),
			want: map[string]interface{}{"v1": int(1)},
		},
		{
			name: "int32",
			cmd: func() *cobra.Command {
				f := setupFlagSet()
				f.Int32("v1", 1, "v1")

				cmd := &cobra.Command{}
				cmd.Flags().AddFlagSet(f)
				return cmd
			}(),
			want: map[string]interface{}{"v1": int(1)},
		},
		{
			name: "int64",
			cmd: func() *cobra.Command {
				f := setupFlagSet()
				f.Int64("v1", 1, "v1")

				cmd := &cobra.Command{}
				cmd.Flags().AddFlagSet(f)
				return cmd
			}(),
			want: map[string]interface{}{"v1": int64(1)},
		},
		{
			name: "intSlice",
			cmd: func() *cobra.Command {
				f := setupFlagSet()
				f.IntSlice("v1", []int{1, 2, 3}, "v1")
				cmd := &cobra.Command{}
				cmd.Flags().AddFlagSet(f)
				return cmd
			}(),
			want: map[string]interface{}{"v1": []int{1, 2, 3}},
		},
		{
			name: "int32Slice",
			cmd: func() *cobra.Command {
				f := setupFlagSet()
				f.Int32Slice("v1", []int32{1, 2, 3}, "v1")
				cmd := &cobra.Command{}
				cmd.Flags().AddFlagSet(f)
				return cmd
			}(),
			want: map[string]interface{}{"v1": []int32{1, 2, 3}},
		},
		{
			name: "int64Slice",
			cmd: func() *cobra.Command {
				f := setupFlagSet()
				f.Int64Slice("v1", []int64{1, 2, 3}, "v1")
				cmd := &cobra.Command{}
				cmd.Flags().AddFlagSet(f)
				return cmd
			}(),
			want: map[string]interface{}{"v1": []int64{1, 2, 3}},
		},
		{
			name: "ip",
			cmd: func() *cobra.Command {
				f := setupFlagSet()
				f.IP("v1", net.ParseIP("0.0.0.0"), "v1")

				cmd := &cobra.Command{}
				cmd.Flags().AddFlagSet(f)
				return cmd
			}(),
			want: map[string]interface{}{"v1": net.ParseIP("0.0.0.0")},
		},
		{
			name: "ipSlice",
			cmd: func() *cobra.Command {
				f := setupFlagSet()
				f.IPSlice("v1", []net.IP{net.ParseIP("0.0.0.0"), net.ParseIP("1.1.1.1")}, "v1")

				cmd := &cobra.Command{}
				cmd.Flags().AddFlagSet(f)
				return cmd
			}(),
			want: map[string]interface{}{"v1": []net.IP{net.ParseIP("0.0.0.0"), net.ParseIP("1.1.1.1")}},
		},
		{
			name: "ipMask",
			cmd: func() *cobra.Command {
				f := setupFlagSet()
				f.IPMask("v1", net.ParseIP("0.0.0.0").DefaultMask(), "v1")

				cmd := &cobra.Command{}
				cmd.Flags().AddFlagSet(f)
				return cmd
			}(),
			want: map[string]interface{}{"v1": net.ParseIP("0.0.0.0").DefaultMask()},
		},
		{
			name: "ipNet",
			cmd: func() *cobra.Command {
				f := setupFlagSet()
				_, def, _ := net.ParseCIDR("0.0.0.0/0")
				f.IPNet(
					"v1",
					*def,
					"v1",
				)

				cmd := &cobra.Command{}
				cmd.Flags().AddFlagSet(f)
				return cmd
			}(),
			want: map[string]interface{}{
				"v1": func() net.IPNet {
					_, def, _ := net.ParseCIDR("0.0.0.0/0")
					return *def
				}(),
			},
		},
		{
			name: "string",
			cmd: func() *cobra.Command {
				f := setupFlagSet()
				f.String("v1", "123", "v1")

				cmd := &cobra.Command{}
				cmd.Flags().AddFlagSet(f)
				return cmd
			}(),
			want: map[string]interface{}{"v1": "123"},
		},
		{
			name: "stringArray",
			cmd: func() *cobra.Command {
				f := setupFlagSet()
				f.StringArray("v1", []string{"123", "456", "789"}, "v1")

				cmd := &cobra.Command{}
				cmd.Flags().AddFlagSet(f)
				return cmd
			}(),
			want: map[string]interface{}{"v1": []string{"123", "456", "789"}},
		},
		{
			name: "stringSlice",
			cmd: func() *cobra.Command {
				f := setupFlagSet()
				f.StringSlice("v1", []string{"123", "456", "789"}, "v1")

				cmd := &cobra.Command{}
				cmd.Flags().AddFlagSet(f)
				return cmd
			}(),
			want: map[string]interface{}{"v1": []string{"123", "456", "789"}},
		},
		{
			name: "stringToInt",
			cmd: func() *cobra.Command {
				f := setupFlagSet()
				f.StringToInt("v1", map[string]int{"a": 1, "b": 2}, "v1")

				cmd := &cobra.Command{}
				cmd.Flags().AddFlagSet(f)
				return cmd
			}(),
			want: map[string]interface{}{"v1": map[string]int{"a": 1, "b": 2}},
		},
		{
			name: "stringToInt64",
			cmd: func() *cobra.Command {
				f := setupFlagSet()
				f.StringToInt64("v1", map[string]int64{"a": 1, "b": 2}, "v1")

				cmd := &cobra.Command{}
				cmd.Flags().AddFlagSet(f)
				return cmd
			}(),
			want: map[string]interface{}{"v1": map[string]int64{"a": 1, "b": 2}},
		},
		{
			name: "stringToString",
			cmd: func() *cobra.Command {
				f := setupFlagSet()
				f.StringToString("v1", map[string]string{"a": "a", "b": "b"}, "v1")

				cmd := &cobra.Command{}
				cmd.Flags().AddFlagSet(f)
				return cmd
			}(),
			want: map[string]interface{}{"v1": map[string]string{"a": "a", "b": "b"}},
		},
		{
			name: "uint",
			cmd: func() *cobra.Command {
				f := setupFlagSet()
				f.Uint("v1", 1, "v1")

				cmd := &cobra.Command{}
				cmd.Flags().AddFlagSet(f)
				return cmd
			}(),
			want: map[string]interface{}{"v1": uint(1)},
		},
		{
			name: "uint16",
			cmd: func() *cobra.Command {
				f := setupFlagSet()
				f.Uint16("v1", 1, "v1")

				cmd := &cobra.Command{}
				cmd.Flags().AddFlagSet(f)
				return cmd
			}(),
			want: map[string]interface{}{"v1": uint16(1)},
		},
		{
			name: "uint32",
			cmd: func() *cobra.Command {
				f := setupFlagSet()
				f.Uint32("v1", 1, "v1")

				cmd := &cobra.Command{}
				cmd.Flags().AddFlagSet(f)
				return cmd
			}(),
			want: map[string]interface{}{"v1": uint32(1)},
		},
		{
			name: "uint64",
			cmd: func() *cobra.Command {
				f := setupFlagSet()
				f.Uint64("v1", 1, "v1")

				cmd := &cobra.Command{}
				cmd.Flags().AddFlagSet(f)
				return cmd
			}(),
			want: map[string]interface{}{"v1": uint64(1)},
		},
		{
			name: "uintSlice",
			cmd: func() *cobra.Command {
				f := setupFlagSet()
				f.UintSlice("v1", []uint{1, 2, 3}, "v1")
				cmd := &cobra.Command{}
				cmd.Flags().AddFlagSet(f)
				return cmd
			}(),
			want: map[string]interface{}{"v1": []uint{1, 2, 3}},
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
