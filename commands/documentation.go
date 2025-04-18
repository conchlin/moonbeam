package commands

import (
	"moonbeam/utils"

	"github.com/bwmarrin/discordgo"
)

// creates discord message that provides link to the project repo
// this will also have some basic documentation
func HandleDocu(session *discordgo.Session, message *discordgo.MessageCreate) {
	var title string
	var fields []*discordgo.MessageEmbedField

	title = "Moonbeam Documentation"
	fields = []*discordgo.MessageEmbedField{
		{
			Name:  "Project Repo",
			Value: "https://github.com/conchlin/moonbeam",
		},
		{
			Name:  "App Documentation",
			Value: "https://github.com/conchlin/moonbeam#readme",
		},
		{
			Name:  "Want to contribute or have a suggestion?",
			Value: "Feel free to reach out to me on discord (conchlin)! Let's grow this app together :)",
		},
	}

	utils.SendMessageWithFields(session, message.ChannelID, title, "", fields)
}
