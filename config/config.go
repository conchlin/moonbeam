package config

import (
	"encoding/json"
	"fmt"
	"log"
	"moonbeam/utils"
	"os"
	"time"
)

// player specific info
type MemberInfo struct {
	Guild  string `json:"guild"`
	Name   string `json:"name"`
	Level  int    `json:"level"`
	Job    string `json:"job"`
	Quests int    `json:"quests"`
	Cards  int    `json:"cards"`
	Fame   int    `json:"fame"`
}

type LastUpdate struct {
	Name     string    `json:"name"`
	UpdateAt time.Time `json:"updated_at"`
}

// overarching categories
type Config struct {
	Discord  discordConfig  `json:"discord"`
	Guild    GuildConfig    `json:"guild"`
	Activity ActivityConfig `json:"last_activity"`
}

// discord specific info
type discordConfig struct {
	SecurityToken string `json:"securityToken"`
}

// guild member specific info
type GuildConfig struct {
	Moonbeam []MemberInfo `json:"moonbeam"`
	Lefay    []MemberInfo `json:"lefay"`
	Basement []MemberInfo `json:"basement"`
}

type ActivityConfig struct {
	Updated []LastUpdate `json:"updates"`
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

func AddMember(user utils.Player, allianceMember string) error {
	config := ParseConfig()
	newMember := MemberInfo{
		Guild:  user.Guild,
		Name:   user.Name,
		Level:  user.Level,
		Job:    user.Job,
		Quests: user.Quests,
		Cards:  user.Cards,
		Fame:   user.Fame,
	}
	timeNow := LastUpdate{
		Name:     user.Name,
		UpdateAt: time.Now(),
	}

	// Append new member to existing member slice
	if allianceMember == "--moonbeam" || allianceMember == "moonbeam" {
		config.Guild.Moonbeam = append(config.Guild.Moonbeam, newMember)
	} else if allianceMember == "--lefay" || allianceMember == "lefay" {
		config.Guild.Lefay = append(config.Guild.Lefay, newMember)
	} else if allianceMember == "--basement" || allianceMember == "basement" {
		config.Guild.Basement = append(config.Guild.Basement, newMember)
	}

	config.Activity.Updated = append(config.Activity.Updated, timeNow)

	// Save the updated configuration to the JSON file
	if err := saveConfig(config); err != nil {
		return err
	}

	return nil
}

func ConvertJsonToPlayer(member MemberInfo) utils.Player {
	return utils.Player{
		Guild:  member.Guild,
		Name:   member.Name,
		Level:  member.Level,
		Job:    member.Job,
		Quests: member.Quests,
		Cards:  member.Cards,
		Fame:   member.Fame,
	}
}

func RefreshMemberList(data []utils.Player) error {
	config := ParseConfig()
	//clear before adding new list
	config.Guild.Moonbeam = nil
	config.Guild.Lefay = nil
	config.Guild.Basement = nil

	for _, entry := range data {
		updatedMember := MemberInfo{
			Guild:  entry.Guild,
			Name:   entry.Name,
			Level:  entry.Level,
			Job:    entry.Job,
			Quests: entry.Quests,
			Cards:  entry.Cards,
			Fame:   entry.Fame,
		}

		// Append new member to existing member slice
		if updatedMember.Guild == "moonbeam" {
			config.Guild.Moonbeam = append(config.Guild.Moonbeam, updatedMember)
		} else if updatedMember.Guild == "LeFay" {
			config.Guild.Lefay = append(config.Guild.Lefay, updatedMember)
		} else if updatedMember.Guild == "Basement" {
			config.Guild.Basement = append(config.Guild.Basement, updatedMember)
		}
	}

	// Save the updated configuration to the JSON file
	if err := saveConfig(config); err != nil {
		return err
	}

	return nil
}

func RemoveMember(memberName string, allianceMember string) error {
	config := ParseConfig()

	index := -1
	if allianceMember == "--moonbeam" || allianceMember == "moonbeam" ||
		allianceMember == "--lefay" || allianceMember == "lefay" ||
		allianceMember == "--basement" || allianceMember == "basement" {
		for i, member := range config.Guild.Moonbeam {
			if member.Name == memberName {
				index = i
				break
			}
		}
	}

	if index == -1 {
		return fmt.Errorf("member %s not found", memberName)
	}

	// Remove the member from the slice
	if allianceMember == "--moonbeam" || allianceMember == "moonbeam" {
		config.Guild.Moonbeam = append(config.Guild.Moonbeam[:index], config.Guild.Moonbeam[index+1:]...)
	} else if allianceMember == "--lefay" || allianceMember == "moonbeam" {
		config.Guild.Lefay = append(config.Guild.Lefay[:index], config.Guild.Lefay[index+1:]...)
	} else if allianceMember == "--basement" || allianceMember == "basement" {
		config.Guild.Basement = append(config.Guild.Basement[:index], config.Guild.Basement[index+1:]...)
	}

	// Save the updated configuration to the JSON file
	if err := saveConfig(config); err != nil {
		return err
	}

	return nil
}

func UpdateTimestamp(flaggedChar string) error {
	config := ParseConfig()

	for i, player := range config.Activity.Updated {
		if player.Name == flaggedChar {
			config.Activity.Updated[i].UpdateAt = time.Now()
		}
	}

	// Save the updated configuration to the JSON file
	if err := saveConfig(config); err != nil {
		return err
	}

	return nil
}
