package config

import (
	"encoding/json"
	"log"
	"os"
)

// player specific info
type MemberInfo struct {
	Guild  string `json:"guild"`
	Name   string `json:"name"`
	Level  int    `json:"level"`
	Exp    string `json:"exp"`
	Gender string `json:"gender"`
	Job    string `json:"job"`
	Quests int    `json:"quests"`
	Cards  int    `json:"cards"`
	Donor  bool   `json:"donor"`
	Fame   int    `json:"fame"`
}

// overarching categories
type Config struct {
	Discord discordConfig `json:"discord"`
	Guild   guildConfig   `json:"guild"`
}

// discord specific info
type discordConfig struct {
	SecurityToken string `json:"securityToken"`
}

// guild member specific info
type guildConfig struct {
	Members []MemberInfo `json:"members"`
}

func ParseConfig() *Config {
	jsonFile, err := os.Open("config/config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	jsonDecoder := json.NewDecoder(jsonFile)
	config := &Config{}
	if err := jsonDecoder.Decode(config); err != nil {
		log.Fatal(err)
	}

	return config
}
