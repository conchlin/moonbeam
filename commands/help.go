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
	## The **Wanderlust** Discord App!
	### Command List
	**$createparty** create a new party instance and allow others to join
	**$joinparty** join an already created party
	**$showparties** list all parties currently available
	`

	return docu
}
