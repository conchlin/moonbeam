package commands

import (
	"fmt"
	"moonbeam/lottery"
	"moonbeam/utils"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func HandleLotteryCreation(session *discordgo.Session, message *discordgo.MessageCreate) {
	msgSplit := strings.Fields(message.Content)
	if len(msgSplit) != 2 {
		utils.SendMessage(session, message.ChannelID, "Syntax Issue", "Please use the following syntax $createlottery <Buy-In Amount>")
		return
	}

	amount := strings.TrimSpace(msgSplit[1])
	n, err := strconv.Atoi(amount)
	if err != nil {
		utils.SendErrorMessage(session, message.ChannelID, fmt.Errorf("the amount entered is not a valid number"))
		return
	}

	perms, e := session.UserChannelPermissions(message.Author.ID, message.ChannelID)
	if e != nil {
		fmt.Println(e.Error())
	}

	if perms&discordgo.PermissionManageMessages != discordgo.PermissionManageMessages {
		utils.SendErrorMessage(session, message.ChannelID, fmt.Errorf("you do not have permission to use this command"))
		return
	}

	lottery.CreateLottery(n)
	utils.SendMessage(session, message.ChannelID, "New Lottery!", "A new lottery has been created with a "+amount+" buy in amount.")
}
