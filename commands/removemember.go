package commands

import (
	"fmt"
	"moonbeam/config"
	"moonbeam/utils"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func HandleMemberRemoval(session *discordgo.Session, message *discordgo.MessageCreate) {
	msgSplit := strings.SplitAfter(message.Content, " ")
	if len(msgSplit) != 3 {
		utils.SendMessage(session, message.ChannelID, "Error", "Please use the following syntax $removemember <ign> <--guild>")
		return
	}

	perms, e := session.UserChannelPermissions(message.Author.ID, message.ChannelID)
	if e != nil {
		fmt.Println(e.Error())
	}
	if perms&discordgo.PermissionManageMessages == discordgo.PermissionManageMessages {
		config.RemoveMember(msgSplit[1], msgSplit[2])
		utils.SendMessage(session, message.ChannelID, msgSplit[1], "Removed from the member list")
	}
}
