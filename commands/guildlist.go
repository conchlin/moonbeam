package commands

import (
	"fmt"
	"moonbeam/config"
	"moonbeam/utils"
	"strings"

	"github.com/bwmarrin/discordgo"
	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

var currentMemberList []utils.Player
var memberStringList []string

func HandleGuildList(session *discordgo.Session, message *discordgo.MessageCreate) {
	msgSplit := strings.SplitAfter(message.Content, " ")
	var formattedList strings.Builder

	go func() {
		// loads char info to currentMemberList
		loadCurrentGuildMembers(msgSplit[1])

		c := collate.New(language.English, collate.IgnoreCase)
		c.SortStrings(memberStringList)

		formattedList.WriteString(fmt.Sprintf("Total Members: %v\n\n", len(currentMemberList)))
		formattedList.WriteString("```")
		for _, char := range memberStringList {
			formattedList.WriteString(fmt.Sprintf("%v\r\n", char))
		}
		formattedList.WriteString("```")

		utils.SendMessage(session, message.ChannelID, "Guild List", formattedList.String())

		clear()
	}()
}

func loadCurrentGuildMembers(allianceMember string) {
	cfg := config.ParseConfig()

	if allianceMember == "moonbeam" {
		for _, member := range cfg.Guild.Moonbeam {
			player := config.ConvertJsonToPlayer(member)
			currentMemberList = append(currentMemberList, player)
			memberStringList = append(memberStringList, player.Name)
		}
	} else if allianceMember == "lefay" {
		for _, member := range cfg.Guild.Lefay {
			player := config.ConvertJsonToPlayer(member)
			currentMemberList = append(currentMemberList, player)
			memberStringList = append(memberStringList, player.Name)
		}
	} else if allianceMember == "basement" {
		for _, member := range cfg.Guild.Basement {
			player := config.ConvertJsonToPlayer(member)
			currentMemberList = append(currentMemberList, player)
			memberStringList = append(memberStringList, player.Name)
		}
	}
}

func clear() {
	currentMemberList = nil
	memberStringList = nil
}
