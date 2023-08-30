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
	**$createparty <type> <time>** create a new party instance and allow others to join
	**$joinparty <party_id> <character_name> <class> <level>** join an already created party
	**$removemember <party_id> <character_name>** removes a member from an already created party
	**$deleteparty <party_id** deletes an already existing party
	**$showparties** list all parties currently available
	`

	return docu
}
