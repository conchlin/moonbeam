package commands

import (
	"fmt"
	"strconv"
	"strings"
	"wanderlust/party"

	"github.com/bwmarrin/discordgo"
)

// syntax for command is $removemember <party_id> <character_name>

func HandleExpelMember(session *discordgo.Session, message *discordgo.MessageCreate) {
	session.ChannelMessageSend(message.ChannelID, expelMember(session, message))
}

func expelMember(session *discordgo.Session, message *discordgo.MessageCreate) string {
	var confirmation string
	authorName := message.Author.Username

	partyId, playerName, strErr := splitRemoveMemberString(message.Content)
	if strErr != nil {
		return fmt.Sprintf("Error in removing member: %v \r\n The command syntax should be $expel <player_name> <party_id>", strErr)
	}
	partyInstance := party.GetPartyByID(partyId)
	player, err := party.GetPartyMemberByName(partyInstance, playerName)
	if err != nil {
		return fmt.Sprintf("%s cannot be found in party", playerName)
	}

	removeErr := partyInstance.RemoveMember(GetNickname(session, authorName), player)
	if removeErr != nil {
		return fmt.Sprintf("error in removing member: %v", removeErr)
	}

	confirmation = fmt.Sprintf("%s has been removed from the party", playerName)
	return confirmation
}

func splitRemoveMemberString(msg string) (int, string, error) {
	msgSplit := strings.SplitAfter(msg, " ")

	playerName := strings.TrimSpace(msgSplit[1])

	idIntVal, err := strconv.Atoi(strings.TrimSpace(msgSplit[2]))
	if err != nil {
		return 0, "", fmt.Errorf("invalid party id")
	}

	if !party.IsValidPartyID(idIntVal) {
		return 0, "", fmt.Errorf("this party id does not exist")
	}

	return idIntVal, playerName, nil

}
