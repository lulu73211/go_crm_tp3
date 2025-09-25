package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Type string `mapstructure:"type"` // "memory" (pour l’instant)
}

func Load(path string) (*AppConfig, error) {
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")

	// valeurs par défaut
	v.SetDefault("type", "memory")

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}
	var cfg AppConfig
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}
	return &cfg, nil
}
