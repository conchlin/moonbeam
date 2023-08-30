package commands

import (
	"fmt"
	"strconv"
	"strings"
	"wanderlust/party"

	"github.com/bwmarrin/discordgo"
)

/**
Delete a party from the registeredParties list
The only user that can use this command is the original party creator
syntax -> $deleteparty <party_id>
*/

func HandlePartyDeletion(session *discordgo.Session, message *discordgo.MessageCreate) {
	session.ChannelMessageSend(message.ChannelID, partyDeletion(session, message))
}

func partyDeletion(session *discordgo.Session, message *discordgo.MessageCreate) string {
	var confirmation string

	id, err := parseDeletionString(message.Content)
	if err != nil {
		return fmt.Sprintf("Error in party deletion: %v \r\n The command syntax should be $deleteparty <party_id", err)
	}

	if party.DeleteParty(id) {
		confirmation = "The party has been successfully deleted!"
	} else {
		return fmt.Sprintf("Error in party deletion of ID %v", id)
	}

	return confirmation
}

func parseDeletionString(msg string) (int, error) {
	msgSplit := strings.SplitAfter(msg, " ")
	idIntVal, err := strconv.Atoi(strings.TrimSpace(msgSplit[1]))
	if err != nil {
		return 0, fmt.Errorf("invalid party id")
	}

	if !party.IsValidPartyID(idIntVal) {
		return 0, fmt.Errorf("this party id does not exist")
	}

	return idIntVal, nil
}
