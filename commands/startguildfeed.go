package commands

import "github.com/bwmarrin/discordgo"

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
			s.ChannelMessageSendEmbed(feedChannel, &discordgo.MessageEmbed{
				Title:       "Guild Update",
				Description: event,
				Color:       0x2cdaca,
			})
		}
	}
}
