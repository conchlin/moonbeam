package config

import (
	"encoding/json"
	"log"
	"moonbeam/utils"
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

func saveConfig(config *Config) error {
	jsonFile, err := os.Create("config/config.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	jsonEncoder := json.NewEncoder(jsonFile)
	jsonEncoder.SetIndent("", "  ")

	if err := jsonEncoder.Encode(config); err != nil {
		return err
	}

	return nil
}

func AddMember(user utils.Player) error {
	config := ParseConfig()
	newMember := MemberInfo{
		Guild:  user.Guild,
		Name:   user.Name,
		Level:  user.Level,
		Exp:    user.Exp,
		Gender: user.Gender,
		Job:    user.Job,
		Quests: user.Quests,
		Cards:  user.Cards,
		Donor:  user.Donor,
		Fame:   user.Fame,
	}

	// Append new member to existing member slice
	config.Guild.Members = append(config.Guild.Members, newMember)

	// Save the updated configuration to the JSON file
	if err := saveConfig(config); err != nil {
		return err
	}

	return nil
}
