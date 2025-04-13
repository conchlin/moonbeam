package commands

import (
	"fmt"
	"log"
	"moonbeam/config"
	"moonbeam/utils"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func HandleNewGuildMember(session *discordgo.Session, message *discordgo.MessageCreate) {
	msgSplit := strings.SplitAfter(message.Content, " ")
	if len(msgSplit) != 3 {
		utils.SendMessage(session, message.ChannelID, "Error", "Please use the following syntax $addmember <ign> <--guild>")
		return
	}

	perms, e := session.UserChannelPermissions(message.Author.ID, message.ChannelID)
	if e != nil {
		fmt.Println(e.Error())
	}
	if perms&discordgo.PermissionManageMessages == discordgo.PermissionManageMessages {
		// verify if new member is a valid character
		playerInfo, err := utils.ParseCharacterJSON(msgSplit[1])
		if err != nil {
			log.Println("Failed to parse character. This character does not exist")
			return
		}
		config.AddMember(playerInfo, msgSplit[2])
		imgUrl := fmt.Sprintf("https://maplelegends.com/api/getavatar?name=%s", playerInfo.Name)
		imgBuf, _ := utils.ParseCharacterImage(imgUrl)

		utils.SendMessageWithImage(session, message.ChannelID, playerInfo.Name, "Successfully added to the guild list", imgBuf.Bytes())
	}
}
