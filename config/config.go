package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config interface {
	Get(string) interface{}
	GetString(key string) string
	GetBool(key string) bool
	GetInt(key string) int
	GetInt32(key string) int32
	GetInt64(key string) int64
	GetUint(key string) uint
	GetUint32(key string) uint32
	GetUint64(key string) uint64
	GetFloat64(key string) float64
	GetTime(key string) time.Time
	GetDuration(key string) time.Duration
	GetIntSlice(key string) []int
	GetStringSlice(key string) []string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	UnmarshalKey(string, interface{}) error
}

var (
	_defaultConfig Config
)

type viperWrapper struct {
	*viper.Viper
}

func (v *viperWrapper) UnmarshalKey(key string, target interface{}) error {
	return v.Viper.UnmarshalKey(key, target)
}

func DefaultConfig() Config {
	if _defaultConfig != nil {
		return _defaultConfig
	}

	viper.AddConfigPath("./")
	viper.AddConfigPath(("/etc/"))
	viper.SetConfigName(".cast")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	_defaultConfig = &viperWrapper{Viper: viper.GetViper()}

	return _defaultConfig
}
