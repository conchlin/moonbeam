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
		config.AddMember(playerInfo)
		session.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
			Title:       playerInfo.Name,
			Description: "Successfully added to the guild list",
			Color:       0x2cdaca,
			Image: &discordgo.MessageEmbedImage{
				URL:    fmt.Sprintf("https://maplelegends.com/api/getavatar?name=%s", playerInfo.Name),
				Width:  70,
				Height: 78,
			},
		})
	}
}
