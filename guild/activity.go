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

			msg, err := s.ChannelMessageSendComplex(feedChannel, &discordgo.MessageSend{
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
			if err != nil {
				fmt.Printf("Error sending message: %v\n", err)
				continue
			}

			addReactions(feedChannel, msg.ID, event.Guild)
		}
	}
}

func addReactions(channelId, messageId, guildName string) {
	var unicode string
	if guildName == "moonbeam" {
		unicode = "moonbeam:1308663568220295229"
	} else if guildName == "LeFay" {
		unicode = "lefay:1308663511894986812"
	} else if guildName == "Basement" {
		unicode = "basement:1308663430797983745"
	} else if guildName == "Torrent" {
		unicode = "torrent:1347677865004634183"
	}

	err := s.MessageReactionAdd(channelId, messageId, unicode)
	if err != nil {
		fmt.Printf("Error adding reaction: %v\n", err)
	}

	//handle streak here
}
