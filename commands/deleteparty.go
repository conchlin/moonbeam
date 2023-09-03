package commands

import (
	"errors"
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

// command handler
func HandlePartyDeletion(session *discordgo.Session, message *discordgo.MessageCreate) {
	id, err := parseDeletionString(message.Content)
	if err != nil {
		session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Error in party deletion: %v \r\n The command syntax should be $deleteparty <party_id>", err))
		return
	}

	if party.DeleteParty(id) {
		session.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
			Title:       "Party Deleted",
			Description: "The party has been successfully deleted!",
			Color:       0x2cdaca,
		})
	} else {
		session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Error in party deletion of ID %v", id))
		return
	}

}

// parseDeletionString parses the input message for the $deleteparty command and extracts
// the party ID to be deleted.
func parseDeletionString(msg string) (int, error) {
	msgSplit := strings.SplitAfter(msg, " ")

	if len(msgSplit) != 2 {
		return 0, errors.New("command parameter is missing")
	}

	idIntVal, err := strconv.Atoi(strings.TrimSpace(msgSplit[1]))
	if err != nil {
		return 0, fmt.Errorf("invalid party id")
	}

	if !party.IsValidPartyID(idIntVal) {
		return 0, fmt.Errorf("this party id does not exist")
	}

	return idIntVal, nil
}
