package config

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

var (
	ErrNilConfig = errors.New("config is nil")
)

// BindEnv return viper config with specified variables
func BindEnv(v *viper.Viper, defaults map[string]interface{}, autoEnv bool, configPath *string) (*viper.Viper, error) {
	if v == nil {
		return nil, ErrNilConfig
	}

	if autoEnv {
		v.AutomaticEnv()
	}

	if configPath != nil {
		v.AddConfigPath(*configPath)
	}

	for key, value := range defaults {
		v.SetDefault(key, value)
		err := v.BindEnv(key)
		if err != nil {
			return nil, fmt.Errorf("bind env error: %w", err)
		}
	}
	return v, nil
}
