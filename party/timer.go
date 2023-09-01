package party

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

func CreateTimer(session *discordgo.Session, message *discordgo.MessageCreate, party *Party) {
	currentTime := time.Now().UTC()
	partyTime := party.Time
	// calculate the difference between the two times
	timeDifference := partyTime.Sub(currentTime)
	timer := time.NewTimer(timeDifference)

	// goroutine to run this part async
	timerDone := make(chan bool)
	go func() {
		<-timer.C
		broadcastPartyMessage(session, message, party)
		timerDone <- true
	}()

	// block until timerDone is true
	<-timerDone
}

// send ping to all party members as a reminder of the party event
func broadcastPartyMessage(session *discordgo.Session, message *discordgo.MessageCreate, party *Party) {
	var builder strings.Builder
	for _, members := range party.Members {
		builder.WriteString(fmt.Sprintf("@%s ", members.DiscordName))
	}

	builder.WriteString(fmt.Sprintf("The %s party is about to begin!", party.Type))
	session.ChannelMessageSend(message.ChannelID, builder.String())
}
