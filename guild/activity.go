package guild

import (
	"fmt"
	"moonbeam/config"
	"moonbeam/utils"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Event struct {
	Achievement string
	Guild       string
}

var playersToUpdate []string
var feedChannel string
var s *discordgo.Session
var active bool

func HandleStartFeed(session *discordgo.Session, message *discordgo.MessageCreate) {
	//always have the player notifications post in same designated channel
	feedChannel = message.ChannelID
	s = session
	active = true

	session.ChannelMessageSendEmbed(feedChannel, &discordgo.MessageEmbed{
		Title:       "Crossroads Alliance Updates",
		Description: "Updates have now been turned on",
		Color:       0x2cdaca,
	})
}

func FlagForUpdate(players []string) {
	playersToUpdate = append(playersToUpdate, players...)

	for _, player := range players {
		config.UpdateTimestamp(player)
	}
}

func CreateFeedPosts(events []Event) {
	if active {
		for _, event := range events {
			// grab the char name which is the first word of event
			split := strings.SplitAfter(event.Achievement, " ")
			charUrl := fmt.Sprintf("https://maplelegends.com/api/getavatar?name=%s", split[0])
			imgBuf, _ := utils.ParseChracterImage(charUrl)

			s.ChannelMessageSendComplex(feedChannel, &discordgo.MessageSend{
				Embed: &discordgo.MessageEmbed{
					Title:       fmt.Sprintf("%s Alliance Update", event.Guild),
					Description: event.Achievement,
					Color:       0x2cdaca,
				},
				File: &discordgo.File{
					Name:   "output.png",
					Reader: imgBuf,
				},
			})
		}
	}
}
