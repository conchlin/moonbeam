package commands

import (
	"moonbeam/config"
	"moonbeam/utils"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func HandleNewGuildMember(session *discordgo.Session, message *discordgo.MessageCreate) {
	msgSplit := strings.Fields(message.Content)
	if len(msgSplit) != 3 {
		utils.SendMessage(session, message.ChannelID, "Syntax Issue", "Please use the following syntax $addmember <ign> <guild>")
		return
	}

	ign := strings.TrimSpace(msgSplit[1])
	guild := strings.TrimSpace(msgSplit[2])

	//perms, e := session.UserChannelPermissions(message.Author.ID, message.ChannelID)
	//if e != nil {
	//	fmt.Println(e.Error())
	//}
	//if perms&discordgo.PermissionManageMessages == discordgo.PermissionManageMessages {
	// verify if new member is a valid character
	playerInfo, err := utils.ParseCharacterJSON(ign)
	if err != nil {
		utils.SendErrorMessage(session, message.ChannelID, err)
		return
	}
	err2 := config.AddMember(playerInfo, guild)
	if err2 != nil {
		utils.SendErrorMessage(session, message.ChannelID, err2)
		return
	}

	utils.SendMessage(session, message.ChannelID, playerInfo.Name, "Successfully added to the guild list")
	//}
}
