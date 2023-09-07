package commands

import (
	"errors"
	"fmt"
	"strings"
	"time"
	"wanderlust/party"
	"wanderlust/utils"

	"github.com/bwmarrin/discordgo"
)

/**
*	create a party utilizing the tooling found in party.go
*	This requires input from the user
*	discord syntax -> $createparty <party_id> <party_type> <party_time>
 */

// command handler
func HandleCreateParty(session *discordgo.Session, message *discordgo.MessageCreate) {
	_type, _time, err := splitCreatePartyString(message.Content)
	if err != nil {
		session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Error in party creation: %v \r\n The command syntax should be $createparty <party_type> <party_date> <start_time>", err))
		return
	}

	// create new party of _type
	newParty := party.NewParty(message.Author.Username, _type, _time)
	session.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
		Title:       "New Party",
		Description: "A new " + _type + " party has been created!",
		Color:       0x2cdaca,
	})

	// calculate the difference between the two times
	currentTime := time.Now().UTC()
	partyTime := newParty.PtTime
	timeDifference := partyTime.Sub(currentTime)
	go utils.CreateTimer(timeDifference, func() {
		party.BroadcastPartyMessage(session, message, newParty)
	})
}

// splitCreatePartyString parses the input message for the $createparty command and extracts
// the party type and party Time
func splitCreatePartyString(msg string) (string, time.Time, error) {
	msgLower := strings.ToLower(msg)
	msgSplit := strings.SplitAfter(msgLower, " ")

	if len(msgSplit) != 4 {
		return "", time.Time{}, errors.New("command parameter is missing")
	}

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
