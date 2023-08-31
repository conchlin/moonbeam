package commands

import (
	"fmt"
	"strconv"
	"strings"
	"wanderlust/party"

	"github.com/bwmarrin/discordgo"
)

// this is just a glorified ping mechanism to alert a discord user that a party has been created
// syntax $invite <discord_name> <party_id>

func HandleInviteMember(session *discordgo.Session, message *discordgo.MessageCreate) {
	session.ChannelMessageSend(message.ChannelID, inviteMember(session, message))
}

func inviteMember(session *discordgo.Session, message *discordgo.MessageCreate) string {
	var confirmation string

	invited, partyId, err := parseInviteString(message.Content)
	if err != nil {
		return fmt.Sprintf("Error in inviting member: %v \r\n The command syntax should be $invite <player_name> <party_id>", err)
	}

	partyInstance := party.GetPartyByID(partyId)
	confirmation = fmt.Sprintf("@%s You have been invited to the %s party. The id is %v", invited, partyInstance.Type, partyId)

	return confirmation
}

func parseInviteString(msg string) (string, int, error) {
	msgSplit := strings.SplitAfter(msg, " ")

	playerName := strings.TrimSpace(msgSplit[1])

	idIntVal, err := strconv.Atoi(strings.TrimSpace(msgSplit[2]))
	if err != nil {
		return "", 0, fmt.Errorf("invalid party id")
	}

	if !party.IsValidPartyID(idIntVal) {
		return "", 0, fmt.Errorf("this party id does not exist")
	}

	return playerName, idIntVal, nil
}
