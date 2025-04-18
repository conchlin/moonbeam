package commands

import (
	"fmt"
	"moonbeam/config"
	"moonbeam/utils"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func HandleMemberRemoval(session *discordgo.Session, message *discordgo.MessageCreate) {
	msgSplit := strings.Fields(message.Content)
	if len(msgSplit) != 3 {
		utils.SendMessage(session, message.ChannelID, "Error", "Please use the following syntax $removemember <ign> <--guild>")
		return
	}

	ign := strings.TrimSpace(msgSplit[1])
	guild := strings.TrimSpace(msgSplit[2])

	perms, e := session.UserChannelPermissions(message.Author.ID, message.ChannelID)
	if e != nil {
		fmt.Println(e.Error())
	}
	if perms&discordgo.PermissionManageMessages != discordgo.PermissionManageMessages {
		utils.SendErrorMessage(session, message.ChannelID, fmt.Errorf("you do not have permission to use this command"))
		return
	}

	err := config.RemoveMember(ign, guild)
	if err != nil {
		utils.SendErrorMessage(session, message.ChannelID, err)
		return
	}
	utils.SendMessage(session, message.ChannelID, ign, "Removed from the member list")

}
