package mermaid

import (
	"github.com/spf13/viper"
	"go.uber.org/dig"
	"net"
	"testing"
	"time"
)

func setupViper() *viper.Viper {
	cfg := viper.New()
	cfg.SetTypeByDefaultValue(true)
	return cfg
}

func TestBindContainer(t *testing.T) {
	testCases := []struct {
		name string
		cfg  *viper.Viper
		exec func(c *dig.Container) error
	}{
		{
			name: "bool",
			cfg: func() *viper.Viper {
				cfg := setupViper()
				cfg.SetDefault("v1", true)
				return cfg
			}(),
			exec: func(container *dig.Container) error {
				type INPUT struct {
					dig.In
					V1 bool `name:"v1"`
				}
				err := container.Invoke(func(input INPUT) {})
				return err
			},
		},
		{
			name: "boolSlice",
			cfg: func() *viper.Viper {
				cfg := setupViper()
				cfg.SetDefault("v1", []bool{true, false, true})
				return cfg
			}(),
			exec: func(container *dig.Container) error {
				type INPUT struct {
					dig.In
					V1 []bool `name:"v1"`
				}
				err := container.Invoke(func(input INPUT) {})
				return err
			},
		},
		{
			name: "bytesHex",
			cfg: func() *viper.Viper {
				cfg := setupViper()
				cfg.SetDefault("v1", []byte("123"))
				return cfg
			}(),
			exec: func(container *dig.Container) error {
				type INPUT struct {
					dig.In
					V1 []byte `name:"v1"`
				}
				err := container.Invoke(func(input INPUT) {})
				return err
			},
		},
		{
			name: "count",
			cfg: func() *viper.Viper {
				cfg := setupViper()
				cfg.SetDefault("v1", 1)
				return cfg
			}(),
			exec: func(container *dig.Container) error {
				type INPUT struct {
					dig.In
					V1 int `name:"v1"`
				}
				err := container.Invoke(func(input INPUT) {})
				return err
			},
		},
		{
			name: "duration",
			cfg: func() *viper.Viper {
				cfg := setupViper()
				cfg.SetDefault("v1", time.Duration(100))
				return cfg
			}(),
			exec: func(container *dig.Container) error {
				type INPUT struct {
					dig.In
					V1 time.Duration `name:"v1"`
				}
				err := container.Invoke(func(input INPUT) {})
				return err
			},
		},
		{
			name: "durationSlice",
			cfg: func() *viper.Viper {
				cfg := setupViper()
				cfg.SetDefault(
					"v1",
					[]time.Duration{
						time.Duration(100),
						time.Duration(200),
						time.Duration(300),
					},
				)
				return cfg
			}(),
			exec: func(container *dig.Container) error {
				type INPUT struct {
					dig.In
					V1 []time.Duration `name:"v1"`
				}
				err := container.Invoke(func(input INPUT) {})
				return err
			},
		},
		{
			name: "float32",
			cfg: func() *viper.Viper {
				cfg := setupViper()
				cfg.SetDefault("v1", float32(1.8))
				return cfg
			}(),
			exec: func(container *dig.Container) error {
				type INPUT struct {
					dig.In
					V1 float64 `name:"v1"`
				}
				err := container.Invoke(func(input INPUT) {})
				return err
			},
		},
		{
			name: "float32Slice",
			cfg: func() *viper.Viper {
				cfg := setupViper()
				cfg.SetDefault("v1", []float32{float32(1.1), float32(2.2)})
				return cfg
			}(),
			exec: func(container *dig.Container) error {
				type INPUT struct {
					dig.In
					V1 []float32 `name:"v1"`
				}
				err := container.Invoke(func(input INPUT) {})
				return err
			},
		},
		{
			name: "float64",
			cfg: func() *viper.Viper {
				cfg := setupViper()
				cfg.SetDefault("v1", float64(1.8))
				return cfg
			}(),
			exec: func(container *dig.Container) error {
				type INPUT struct {
					dig.In
					V1 float64 `name:"v1"`
				}
				err := container.Invoke(func(input INPUT) {})
				return err
			},
		},
		{
			name: "float64Slice",
			cfg: func() *viper.Viper {
				cfg := setupViper()
				cfg.SetDefault("v1", []float64{float64(1.1), float64(2.2)})
				return cfg
			}(),
			exec: func(container *dig.Container) error {
				type INPUT struct {
					dig.In
					V1 []float64 `name:"v1"`
				}
				err := container.Invoke(func(input INPUT) {})
				return err
			},
		},
		{
			name: "int",
			cfg: func() *viper.Viper {
				cfg := setupViper()
				cfg.SetDefault("v1", int(1))
				return cfg
			}(),
			exec: func(container *dig.Container) error {
				type INPUT struct {
					dig.In
					V1 int `name:"v1"`
				}
				err := container.Invoke(func(input INPUT) {})
				return err
			},
		},
		{
			name: "int8",
			cfg: func() *viper.Viper {
				cfg := setupViper()
				cfg.SetDefault("v1", int8(1))
				return cfg
			}(),
			exec: func(container *dig.Container) error {
				type INPUT struct {
					dig.In
					V1 int `name:"v1"`
				}
				err := container.Invoke(func(input INPUT) {})
				return err
			},
		},
		{
			name: "int16",
			cfg: func() *viper.Viper {
				cfg := setupViper()
				cfg.SetDefault("v1", int16(1))
				return cfg
			}(),
			exec: func(container *dig.Container) error {
				type INPUT struct {
					dig.In
					V1 int `name:"v1"`
				}
				err := container.Invoke(func(input INPUT) {})
				return err
			},
		},
		{
			name: "int32",
			cfg: func() *viper.Viper {
				cfg := setupViper()
				cfg.SetDefault("v1", int32(1))
				return cfg
			}(),
			exec: func(container *dig.Container) error {
				type INPUT struct {
					dig.In
					V1 int32 `name:"v1"`
				}
				err := container.Invoke(func(input INPUT) {})
				return err
			},
		},
		{
			name: "int64",
			cfg: func() *viper.Viper {
				cfg := setupViper()
				cfg.SetDefault("v1", int64(1))
				return cfg
			}(),
			exec: func(container *dig.Container) error {
				type INPUT struct {
					dig.In
					V1 int64 `name:"v1"`
				}
				err := container.Invoke(func(input INPUT) {})
				return err
			},
		},
		{
			name: "intSlice",
			cfg: func() *viper.Viper {
				cfg := setupViper()
				cfg.SetDefault("v1", []int{1, 2, 3})
				return cfg
			}(),
			exec: func(container *dig.Container) error {
				type INPUT struct {
					dig.In
					V1 []int `name:"v1"`
				}
				err := container.Invoke(func(input INPUT) {})
				return err
			},
		},
		{
			name: "int32Slice",
			cfg: func() *viper.Viper {
				cfg := setupViper()
				cfg.SetDefault("v1", []int32{1, 2, 3})
				return cfg
			}(),
			exec: func(container *dig.Container) error {
				type INPUT struct {
					dig.In
					V1 []int32 `name:"v1"`
				}
				err := container.Invoke(func(input INPUT) {})
				return err
			},
		},
		{
			name: "int64Slice",
			cfg: func() *viper.Viper {
				cfg := setupViper()
				cfg.SetDefault("v1", []int64{1, 2, 3})
				return cfg
			}(),
			exec: func(container *dig.Container) error {
				type INPUT struct {
					dig.In
					V1 []int64 `name:"v1"`
				}
				err := container.Invoke(func(input INPUT) {})
				return err
			},
		},
		{
			name: "ip",
			cfg: func() *viper.Viper {
				cfg := setupViper()
				cfg.SetDefault("v1", net.IPv4(0, 0, 0, 0))
				return cfg
			}(),
			exec: func(container *dig.Container) error {
				type INPUT struct {
					dig.In
					V1 net.IP `name:"v1"`
				}
				err := container.Invoke(func(input INPUT) {})
				return err
			},
		},
		{
			name: "ipSlice",
			cfg: func() *viper.Viper {
				cfg := setupViper()
				cfg.SetDefault(
					"v1",
					[]net.IP{
						net.IPv4(0, 0, 0, 0),
						net.IPv4(1, 1, 1, 1),
					},
				)
				return cfg
			}(),
			exec: func(container *dig.Container) error {
				type INPUT struct {
					dig.In
					V1 []net.IP `name:"v1"`
				}
				err := container.Invoke(func(input INPUT) {})
				return err
			},
		},
		{
			name: "ipMask",
			cfg: func() *viper.Viper {
				cfg := setupViper()
				cfg.SetDefault("v1", net.ParseIP("0.0.0.0").DefaultMask())
				return cfg
			}(),
			exec: func(container *dig.Container) error {
				type INPUT struct {
					dig.In
					V1 net.IPMask `name:"v1"`
				}
				err := container.Invoke(func(input INPUT) {})
				return err
			},
		},
		{
			name: "ipNet",
			cfg: func() *viper.Viper {
				cfg := setupViper()
				_, def, _ := net.ParseCIDR("0.0.0.0/0")
				cfg.SetDefault("v1", *def)
				return cfg
			}(),
			exec: func(container *dig.Container) error {
				type INPUT struct {
					dig.In
					V1 net.IPNet `name:"v1"`
				}
				err := container.Invoke(func(input INPUT) {})
				return err
			},
		},
		{
			name: "string",
			cfg: func() *viper.Viper {
				cfg := setupViper()
				cfg.SetDefault("v1", "123")
				return cfg
			}(),
			exec: func(container *dig.Container) error {
				type INPUT struct {
					dig.In
					V1 string `name:"v1"`
				}
				err := container.Invoke(func(input INPUT) {})
				return err
			},
		},
		{
			name: "stringArray",
			cfg: func() *viper.Viper {
				cfg := setupViper()
				cfg.SetDefault("v1", []string{"123", "456", "789"})
				return cfg
			}(),
			exec: func(container *dig.Container) error {
				type INPUT struct {
					dig.In
					V1 []string `name:"v1"`
				}
				err := container.Invoke(func(input INPUT) {})
				return err
			},
		},
		{
			name: "stringSlice",
			cfg: func() *viper.Viper {
				cfg := setupViper()
				cfg.SetDefault("v1", []string{"123", "456", "789"})
				return cfg
			}(),
			exec: func(container *dig.Container) error {
				type INPUT struct {
					dig.In
					V1 []string `name:"v1"`
				}
				err := container.Invoke(func(input INPUT) {})
				return err
			},
		},
		{
			name: "stringToInt",
			cfg: func() *viper.Viper {
				cfg := setupViper()
				cfg.SetDefault("v1", map[string]int{"a": 1, "b": 2, "c": 3})
				return cfg
			}(),
			exec: func(container *dig.Container) error {
				type INPUT struct {
					dig.In
					V1 map[string]int `name:"v1"`
				}
				err := container.Invoke(func(input INPUT) {})
				return err
			},
		},
		{
			name: "stringToInt64",
			cfg: func() *viper.Viper {
				cfg := setupViper()
				cfg.SetDefault("v1", map[string]int64{"a": 1, "b": 2, "c": 3})
				return cfg
			}(),
			exec: func(container *dig.Container) error {
				type INPUT struct {
					dig.In
					V1 map[string]int64 `name:"v1"`
				}
				err := container.Invoke(func(input INPUT) {})
				return err
			},
		},
		{
			name: "stringToString",
			cfg: func() *viper.Viper {
				cfg := setupViper()
				cfg.SetDefault("v1", map[string]string{"a": "1", "b": "2", "c": "3"})
				return cfg
			}(),
			exec: func(container *dig.Container) error {
				type INPUT struct {
					dig.In
					V1 map[string]string `name:"v1"`
				}
				err := container.Invoke(func(input INPUT) {})
				return err
			},
		},
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			container := dig.New()
			BindContainer(tt.cfg, container)

			err := tt.exec(container)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
