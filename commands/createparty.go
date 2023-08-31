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
		return fmt.Sprintf("Error in party creation: %v \r\n The command syntax should be $createparty <party_type> <start_time>", err)
	}

	// create new party of _type
	party.NewParty(message.Author.Username, _type, time)
	confirmation = fmt.Sprintf("A new %s party has been created!", _type)

	return confirmation
}

func splitCreatePartyString(msg string) (string, time.Time, error) {
	msgLower := strings.ToLower(msg)
	msgSplit := strings.SplitAfter(msgLower, " ")

	_type := strings.TrimSpace(msgSplit[1])
	today := strings.TrimSpace(msgSplit[2]) == "today"
	var timeVal time.Time
	var err error

	if today {
		// if the player has specified the party should happen today we need to
		// treat it differently. Let's add a dummy date so that we can correctly
		// parse/format our date
		timeVal, err = party.ParseTimeInput("2006-01-01 "+msgSplit[3], true)
		if err != nil {
			return "", time.Time{}, err
		}
	} else {
		// if today is not true they have provided an actual timestamp for a party
		// in the future. So we use that to parse our date
		ts := msgSplit[2] + " " + msgSplit[3]
		timeVal, err = party.ParseTimeInput(ts, false)
		if err != nil {
			return "", time.Time{}, err
		}
	}

	fmt.Println(_type)
	fmt.Println(timeVal)

	return _type, timeVal, nil
}
