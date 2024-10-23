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
		session.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
			Title:       "Error",
			Description: "Please use the following syntax $addmember <ign> <--guild>",
			Color:       0x2cdaca,
		})
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
		imgBuf, _ := utils.ParseChracterImage(imgUrl)

		session.ChannelMessageSendComplex(message.ChannelID, &discordgo.MessageSend{
			Embed: &discordgo.MessageEmbed{
				Title:       playerInfo.Name,
				Description: "Successfully added to the guild list",
				Color:       0x2cdaca,
			},
			File: &discordgo.File{
				Name:   "output.png",
				Reader: imgBuf,
			},
		})
	}
}
