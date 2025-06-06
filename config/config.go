package config

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"moonbeam/utils"
	"os"
	"strings"
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
	Torrent  []MemberInfo `json:"torrent"`
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

// creates the backup at the same time
func ParseConfigForBackup() (string, error) {
	jsonFile, err := os.Open("config/config.json")
	if err != nil {
		return "", fmt.Errorf("failed to open config file: %w", err)
	}
	defer jsonFile.Close()

	var builder strings.Builder
	_, err = io.Copy(&builder, jsonFile)
	if err != nil {
		return "", fmt.Errorf("failed to read config file: %w", err)
	}

	var data map[string]interface{}
	err = json.Unmarshal([]byte(builder.String()), &data)
	if err != nil {
		return "", fmt.Errorf("failed to parse JSON: %w", err)
	}

	formattedJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to format JSON: %w", err)
	}

	backupPath := "config/config_backup.json"
	err = os.WriteFile(backupPath, formattedJSON, 0644)
	if err != nil {
		return "", fmt.Errorf("failed to create backup file: %w", err)
	}

	return string(formattedJSON), nil
}

func ParseBackupConfig() (map[string]interface{}, error) {
	backupPath := "config/config_backup.json"

	backupFile, err := os.Open(backupPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open backup file: %w", err)
	}
	defer backupFile.Close()

	var builder strings.Builder
	_, err = io.Copy(&builder, backupFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read backup file: %w", err)
	}

	var data map[string]interface{}
	err = json.Unmarshal([]byte(builder.String()), &data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return data, nil
}

func ClearAndWriteConfig(newContent map[string]interface{}) error {
	configPath := "config/config.json"

	file, err := os.OpenFile(configPath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	jsonData, err := json.MarshalIndent(newContent, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to encode string as JSON: %w", err)
	}

	_, err = file.Write(jsonData)
	if err != nil {
		return fmt.Errorf("failed to write to config file: %w", err)
	}

	fmt.Println("Config file cleared and updated successfully.")
	return nil
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
	} else if allianceMember == "--torrent" || allianceMember == "torrent" {
		config.Guild.Torrent = append(config.Guild.Torrent, newMember)
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
	config.Guild.Torrent = nil

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
		} else if updatedMember.Guild == "Torrent" {
			config.Guild.Torrent = append(config.Guild.Torrent, updatedMember)
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

	var members *[]MemberInfo
	switch strings.ToLower(allianceMember) {
	case "moonbeam":
		members = &config.Guild.Moonbeam
	case "lefay":
		members = &config.Guild.Lefay
	case "basement":
		members = &config.Guild.Basement
	case "torrent":
		members = &config.Guild.Torrent
	default:
		return fmt.Errorf("unknown alliance member: %s", allianceMember)
	}

	index := -1
	for i, member := range *members {
		if strings.EqualFold(member.Name, memberName) {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("member %s not found in %s", memberName, allianceMember)
	}

	// Remove the member from the slice
	*members = append((*members)[:index], (*members)[index+1:]...)

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
