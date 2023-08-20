package commands

import "github.com/bwmarrin/discordgo"

/**
This command provides very simple documentation on app commands
*/

func HandleHelp(discord *discordgo.Session, message *discordgo.MessageCreate) {
	discord.ChannelMessageSend(message.ChannelID, showDocumentation())
}

func showDocumentation() string {
	docu := `
	The Wanderlust Discord App!

	<info regarding the app goes here>
	`

	return docu
}
