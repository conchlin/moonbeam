package utils

import (
	"bytes"

	"github.com/bwmarrin/discordgo"
)

//helper functions for discord related tasks

// used for all discord messages that need a character image associated
func SendMessageWithImage(session *discordgo.Session, channelID, msgTitle, msgDescription string, charImageBytes []byte) {
	session.ChannelMessageSendComplex(channelID, &discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Title:       msgTitle,
			Description: msgDescription,
			Color:       0x2cdaca,
		},
		File: &discordgo.File{
			Name:   "output.png",
			Reader: bytes.NewReader(charImageBytes),
		},
	})
}

// default discord message
func SendMessage(session *discordgo.Session, channelID, msgTitle, msgDescription string) {
	session.ChannelMessageSendComplex(channelID, &discordgo.MessageSend{
		Embed: &discordgo.MessageEmbed{
			Title:       msgTitle,
			Description: msgDescription,
			Color:       0x2cdaca,
		},
	})
}
