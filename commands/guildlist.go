package commands

import (
	"fmt"
	"moonbeam/config"
	"moonbeam/utils"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var currentMemberList []utils.Player

func HandleGuildList(session *discordgo.Session, message *discordgo.MessageCreate) {
	msgSplit := strings.SplitAfter(message.Content, " ")
	var formattedList strings.Builder

	go func() {
		loadCurrentGuildMembers(msgSplit[1])

		formattedList.WriteString(fmt.Sprintf("Total Members: %v\n\n", len(currentMemberList)))
		formattedList.WriteString("```")
		for _, char := range currentMemberList {
			formattedList.WriteString(fmt.Sprintf("%v\t Lv%v\t %v\r\n", char.Name, char.Level, char.Job))
		}
		formattedList.WriteString("```")

		session.ChannelMessageSendComplex(message.ChannelID, &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       "Guild List",
				Description: formattedList.String(),
				Color:       0x2cdaca,
			},
		})
	}()

	//c := collate.New(language.English, collate.IgnoreCase)
	//c.SortStrings(currentMemberList)
}

func loadCurrentGuildMembers(allianceMember string) {
	cfg := config.ParseConfig()

	if allianceMember == "--moonbeam" {
		for _, member := range cfg.Guild.Moonbeam {
			player := config.ConvertJsonToPlayer(member)
			currentMemberList = append(currentMemberList, player)
		}
	} else if allianceMember == "--lefay" {
		for _, member := range cfg.Guild.Lefay {
			player := config.ConvertJsonToPlayer(member)
			currentMemberList = append(currentMemberList, player)
		}
	} else if allianceMember == "--basement" {
		for _, member := range cfg.Guild.Basement {
			player := config.ConvertJsonToPlayer(member)
			currentMemberList = append(currentMemberList, player)
		}
	}

}
