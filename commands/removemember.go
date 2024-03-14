package commands

import (
	"fmt"
	"moonbeam/config"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func HandleMemberRemoval(session *discordgo.Session, message *discordgo.MessageCreate) {
	msgSplit := strings.SplitAfter(message.Content, " ")

	perms, e := session.UserChannelPermissions(message.Author.ID, message.ChannelID)
	if e != nil {
		fmt.Println(e.Error())
	}
	if perms&discordgo.PermissionManageMessages == discordgo.PermissionManageMessages {
		config.RemoveMember(msgSplit[1])
		session.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
			Title:       msgSplit[1],
			Description: "Removed from the member list",
			Color:       0x2cdaca,
		})
	}
}
