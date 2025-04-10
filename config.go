package main

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config holds the program configuration.
type Config struct {
	SearchTerm           string   `mapstructure:"search-term"`
	JobLocation          string   `mapstructure:"job-location"`
	DescriptionBlacklist []string `mapstructure:"description-blacklist"`

	// Derived set for fast lookup (not loaded directly from YAML).
	DescriptionBlacklistSet map[string]struct{}
}

var Cfg Config

// LoadConfig loads the configuration from config.yaml.
func LoadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	if err := viper.Unmarshal(&Cfg); err != nil {
		return fmt.Errorf("unable to decode configuration: %w", err)
	}

	Cfg.DescriptionBlacklistSet = make(map[string]struct{})
	for _, word := range Cfg.DescriptionBlacklist {
		Cfg.DescriptionBlacklistSet[word] = struct{}{}
	}

	return nil
}
