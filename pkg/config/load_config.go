package config

import (
	"github.com/spf13/viper"
)

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigName("master")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	var config Config

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
