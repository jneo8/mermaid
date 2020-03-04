package mermaid

import (
	"github.com/spf13/viper"
	"go.uber.org/dig"
	"net"
	"time"
)

// BindContainer provide viper's flagset to container.
func BindContainer(cfg *viper.Viper, container *dig.Container) error {
	for key, val := range cfg.AllSettings() {
		var getter interface{}
		switch val.(type) {
		case bool:
			getter = func() bool {
				return cfg.GetBool(key)
			}
		case []bool:
			getter = func() []bool {
				newVal := val.([]bool)
				return newVal
			}
		case []byte:
			getter = func() []byte {
				newVal := val.([]byte)
				return newVal
			}
		case time.Duration:
			getter = func() time.Duration {
				return cfg.GetDuration(key)
			}
		case []time.Duration:
			getter = func() []time.Duration {
				newVal := val.([]time.Duration)
				return newVal
			}
		// viper will turn float32 to float64 when Get value.
		// See https://godoc.org/github.com/spf13/viper#Viper.Get
		case float32, float64:
			getter = func() float64 {
				return cfg.GetFloat64(key)
			}
		case []float32:
			getter = func() []float32 {
				newVal := val.([]float32)
				return newVal
			}
		case []float64:
			getter = func() []float64 {
				newVal := val.([]float64)
				return newVal
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
				newVal := val.([]int)
				return newVal
			}
		case []int32:
			getter = func() []int32 {
				newVal := val.([]int32)
				return newVal
			}
		case []int64:
			getter = func() []int64 {
				newVal := val.([]int64)
				return newVal
			}
		case net.IP:
			getter = func() net.IP {
				newVal := val.(net.IP)
				return newVal
			}
		case []net.IP:
			getter = func() []net.IP {
				newVal := val.([]net.IP)
				return newVal
			}
		case net.IPMask:
			getter = func() net.IPMask {
				newVal := val.(net.IPMask)
				return newVal
			}
		case net.IPNet:
			getter = func() net.IPNet {
				newVal := val.(net.IPNet)
				return newVal
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
				newVal := val.(map[string]int)
				return newVal
			}
		case map[string]int64: // string_to_int64
			getter = func() map[string]int64 {
				newVal := val.(map[string]int64)
				return newVal
			}
		case map[string]string: // string_to_sting
			getter = func() map[string]string {
				newVal := val.(map[string]string)
				return newVal
			}
		}
		if err := container.Provide(getter, dig.Name(key)); err != nil {
			return err
		}
	}
	return nil
}
