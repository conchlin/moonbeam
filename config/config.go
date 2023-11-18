package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

// discord specific info
type discordConfig struct {
	SecurityToken string
}

// overarching categories
type Config struct {
	Discord discordConfig
}

func ParseConfig() *Config {
	tomlFile := &Config{}

	if _, err := toml.DecodeFile("config/config.toml", tomlFile); err != nil {
		log.Fatal(err)
	}

	return tomlFile
}
