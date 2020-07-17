package mermaid

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/dig"
	"net"
	"time"
)

// BindContainer provide viper's flagset to container.
func BindContainer(cfg *viper.Viper, container *dig.Container) error {
	for k, v := range cfg.AllSettings() {
		key := k
		val := v

		// Skip help.
		if key == "help" {
			continue
		}
		var getter interface{}
		switch value := val.(type) {
		case bool:
			getter = func() bool {
				return cfg.GetBool(key)
			}
		case []bool:
			getter = func() []bool {
				return value
			}
		case []byte:
			getter = func() []byte {
				return value
			}
		case time.Duration:
			getter = func() time.Duration {
				return cfg.GetDuration(key)
			}
		case []time.Duration:
			getter = func() []time.Duration {
				return value
			}
		// viper will turn float32 to float64 when Get value.
		// See https://godoc.org/github.com/spf13/viper#Viper.Get
		case float32, float64:
			getter = func() float64 {
				return cfg.GetFloat64(key)
			}
		case []float32:
			getter = func() []float32 {
				return value
			}
		case []float64:
			getter = func() []float64 {
				return value
			}
		case uint:
			getter = func() uint {
				return cfg.GetUint(key)
			}
		case uint8:
			getter = func() uint8 {
				return value
			}
		case uint16:
			getter = func() uint16 {
				return value
			}
		case uint32:
			getter = func() uint32 {
				return cfg.GetUint32(key)
			}
		case uint64:
			getter = func() uint64 {
				return cfg.GetUint64(key)
			}
		case int: // count, int8, int16, int32
			getter = func() int {
				return cfg.GetInt(key)
			}
		case int64:
			getter = func() int64 {
				return cfg.GetInt64(key)
			}
		case []int:
			getter = func() []int {
				return value
			}
		case []int32:
			getter = func() []int32 {
				return value
			}
		case []int64:
			getter = func() []int64 {
				return value
			}
		case net.IP:
			getter = func() net.IP {
				return value
			}
		case []net.IP:
			getter = func() []net.IP {
				return value
			}
		case net.IPMask:
			getter = func() net.IPMask {
				return value
			}
		case net.IPNet:
			getter = func() net.IPNet {
				return value
			}
		case string:
			getter = func() string {
				return cfg.GetString(key)
			}
		case []string: // stringSlice, stringArray
			getter = func() []string {
				return cfg.GetStringSlice(key)
			}
		case map[string]int: // string_to_int
			getter = func() map[string]int {
				return value
			}
		case map[string]int64: // string_to_int64
			getter = func() map[string]int64 {
				return value
			}
		case map[string]string: // string_to_sting
			getter = func() map[string]string {
				return value
			}
		}
		if err := container.Provide(getter, dig.Name(key)); err != nil {
			return fmt.Errorf("Get key %s err: %s", key, err)
		}
	}
	return nil
}
