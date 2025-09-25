package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Type     string `mapstructure:"type"`      // "json" | "memory" (GORM plus tard)
	JSONPath string `mapstructure:"json_path"` // e.g. data/contacts.json
}

func Load(path string) (*AppConfig, error) {
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")

	// défauts
	v.SetDefault("type", "memory")
	v.SetDefault("json_path", "data/contacts.json")

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}
	var cfg AppConfig
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}
	return &cfg, nil
}
