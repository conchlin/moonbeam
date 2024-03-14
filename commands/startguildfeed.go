package commands

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var feedChannel string
var s *discordgo.Session
var active bool

func HandleStartFeed(session *discordgo.Session, message *discordgo.MessageCreate) {
	//always have the player notifications post in same designated channel
	feedChannel = message.ChannelID
	s = session
	active = true

	session.ChannelMessageSendEmbed(feedChannel, &discordgo.MessageEmbed{
		Title:       "Moonbeam Guild Updates",
		Description: "Updates have now been turned on",
		Color:       0x2cdaca,
	})
}

func CreateFeedPosts(events []string) {
	if active {
		for _, event := range events {
			// grab the char name which is the first word of event
			split := strings.SplitAfter(event, " ")
			s.ChannelMessageSendEmbed(feedChannel, &discordgo.MessageEmbed{
				Title:       "Guild Update",
				Description: event,
				Color:       0x2cdaca,
				Image: &discordgo.MessageEmbedImage{
					URL:    fmt.Sprintf("https://maplelegends.com/api/getavatar?name=%s", split[0]),
					Width:  70,
					Height: 78,
				},
			})
		}
	}
}
