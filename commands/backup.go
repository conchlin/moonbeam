package commands

import (
	"fmt"
	"moonbeam/config"

	"github.com/bwmarrin/discordgo"
)

func HandleBackupCreation(session *discordgo.Session, message *discordgo.MessageCreate) {
	perms, e := session.UserChannelPermissions(message.Author.ID, message.ChannelID)
	if e != nil {
		fmt.Println(e.Error())
	}
	if perms&discordgo.PermissionManageMessages == discordgo.PermissionManageMessages {
		_, err := config.ParseConfigForBackup()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		session.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
			Title:       "Success",
			Description: "Backup has been created",
			Color:       0x2cdaca,
		})
	}
}

func HandleApplyingBackup(session *discordgo.Session, message *discordgo.MessageCreate) {
	perms, e := session.UserChannelPermissions(message.Author.ID, message.ChannelID)
	if e != nil {
		fmt.Println(e.Error())
	}

	if perms&discordgo.PermissionManageMessages == discordgo.PermissionManageMessages {
		backupText, err := config.ParseBackupConfig()
		if err != nil {
			fmt.Println("Error parsing backup:", err)
			return
		}

		err2 := config.ClearAndWriteConfig(backupText)
		if err2 != nil {
			fmt.Println("Error:", err)
			return
		}

		session.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
			Title:       "Success",
			Description: "Current config has been replaced by the backup.",
			Color:       0x2cdaca,
		})
	}
}
