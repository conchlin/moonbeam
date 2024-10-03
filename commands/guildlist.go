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
	var formattedList strings.Builder

	go func() {
		loadCurrentGuildMembers()

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

func loadCurrentGuildMembers() {
	cfg := config.ParseConfig()

	for _, member := range cfg.Guild.Members {
		player := config.ConvertJsonToPlayer(member)
		currentMemberList = append(currentMemberList, player)
	}

}
