package commands

import (
	"fmt"
	"moonbeam/config"
	"strings"

	"github.com/bwmarrin/discordgo"
	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

var currentMemberList []string

func HandleGuildList(session *discordgo.Session, message *discordgo.MessageCreate) {
	var formattedList strings.Builder
	loadCurrentGuildMembers()

	c := collate.New(language.English, collate.IgnoreCase)
	c.SortStrings(currentMemberList)

	formattedList.WriteString(fmt.Sprintf("Total Members: %v\n\n", len(currentMemberList)))
	for _, char := range currentMemberList {
		formattedList.WriteString(fmt.Sprintf(char + "\r\n"))
	}

	session.ChannelMessageSendComplex(message.ChannelID, &discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Title:       "Guild List",
			Description: formattedList.String(),
			Color:       0x2cdaca,
		},
	})
}

func loadCurrentGuildMembers() {
	cfg := config.ParseConfig()

	for _, member := range cfg.Guild.Members {
		currentMemberList = append(currentMemberList, member.Name)
	}

}
