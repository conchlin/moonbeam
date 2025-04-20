package commands

import (
	"fmt"
	"moonbeam/lottery"
	"moonbeam/utils"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func HandleAddContestant(session *discordgo.Session, message *discordgo.MessageCreate) {
	msgSplit := strings.Fields(message.Content)

	if len(msgSplit) != 2 {
		utils.SendMessage(session, message.ChannelID, "Syntax Issue", "Please use the following syntax $addLottery <player_name>")
		return
	}

	name := strings.TrimSpace(msgSplit[1])

	perms, e := session.UserChannelPermissions(message.Author.ID, message.ChannelID)
	if e != nil {
		fmt.Println(e.Error())
	}

	if perms&discordgo.PermissionManageMessages != discordgo.PermissionManageMessages {
		utils.SendErrorMessage(session, message.ChannelID, fmt.Errorf("you do not have permission to use this command"))
		return
	}

	c := lottery.CreateContestant(name)
	activeLotto := lottery.GetActiveLottery()
	activeLotto.AddContestant(c)

	utils.SendMessage(session, message.ChannelID, name, "Successfully added to the lottery!")
}
