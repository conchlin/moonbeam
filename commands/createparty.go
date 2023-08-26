package commands

import (
	"fmt"
	"strings"
	"time"
	"wanderlust/party"

	"github.com/bwmarrin/discordgo"
)

/**
*	create a party utilizing the tooling found in party.go
*	This requires input from the user
*	discord syntax -> $createparty <party_id> <party_type> <party_time>
 */

func HandleCreateParty(discord *discordgo.Session, message *discordgo.MessageCreate) {
	discord.ChannelMessageSend(message.ChannelID, partyCreation(message))
}

func partyCreation(message *discordgo.MessageCreate) string {
	var confirmation string

	// verify all components of the discord syntax
	_type, time, err := splitCreatePartyString(message.Content)
	if err != nil {
		return fmt.Sprintf("There is an error in the syntax of the party command. Creation is not possible: %v", err)
	}

	// create new party of _type
	party.NewParty(_type, time)
	confirmation = fmt.Sprintf("A new %s party has been created!", _type)

	return confirmation
}

func splitCreatePartyString(msg string) (string, time.Time, error) {
	msgSplit := strings.SplitAfter(msg, " ")

	// todo check for a valid party type
	_type := strings.TrimSpace(msgSplit[1])

	timeVal, err := party.ParseTimeInput(msgSplit[2])
	if err != nil {
		return "", time.Time{}, fmt.Errorf("invalid party time")
	}

	fmt.Println(_type)
	fmt.Println(timeVal)

	return _type, timeVal, nil
}
