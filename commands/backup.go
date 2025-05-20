package commands

import (
	"fmt"
	"moonbeam/config"
	"moonbeam/utils"

	"github.com/bwmarrin/discordgo"
)

func HandleBackupCreation(session *discordgo.Session, message *discordgo.MessageCreate) {
	//perms, e := session.UserChannelPermissions(message.Author.ID, message.ChannelID)
	//if e != nil {
	//	fmt.Println(e.Error())
	//}
	//if perms&discordgo.PermissionManageMessages == discordgo.PermissionManageMessages {
	_, err := config.ParseConfigForBackup()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	utils.SendMessage(session, message.ChannelID, "Success", "Backup has been created")
	//}
}

func HandleApplyingBackup(session *discordgo.Session, message *discordgo.MessageCreate) {
	//perms, e := session.UserChannelPermissions(message.Author.ID, message.ChannelID)
	//if e != nil {
	//	fmt.Println(e.Error())
	//}

	//if perms&discordgo.PermissionManageMessages == discordgo.PermissionManageMessages {
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

	utils.SendMessage(session, message.ChannelID, "Success", "Current config has been replaced by the backup.")
	//}
}
