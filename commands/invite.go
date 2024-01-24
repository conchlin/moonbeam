package commands

import (
	"errors"
	"fmt"
	"moonbeam/party"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// this is just a glorified ping mechanism to alert a discord user that a party has been created
// syntax $invite <discord_name> <party_id>

// command handler
func HandleInviteMember(session *discordgo.Session, message *discordgo.MessageCreate) {
	invited, partyId, err := parseInviteString(message.Content)
	if err != nil {
		session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Error in inviting member: %v \r\n The command syntax should be $invite <player_name> <party_id>", err))
		return
	}

	partyInstance := party.GetPartyByID(partyId)
	session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("@%s You have been invited to the %s party. The id is %v", invited, partyInstance.Type, partyId))
}

// parseInviteString parses the input message for the $invite command and extracts
// the player name and party ID to send an invitation.
func parseInviteString(msg string) (string, int, error) {
	msgSplit := strings.SplitAfter(msg, " ")

	if len(msgSplit) != 3 {
		return "", 0, errors.New("command parameter is missing")
	}

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
