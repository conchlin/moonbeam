package commands

import (
	"moonbeam/utils"

	"github.com/bwmarrin/discordgo"
)

/**
This command provides very simple documentation on app commands
*/

func HandleHelp(session *discordgo.Session, message *discordgo.MessageCreate) {
	var title string
	var description string
	var fields []*discordgo.MessageEmbedField

	title = "The Moonbeam Discord App!"
	description = "Commands List"
	fields = []*discordgo.MessageEmbedField{
		{
			Name:  "$showparties",
			Value: "Display a list of all currently active parties.",
		},
		{
			Name:  "$createparty <type> <date> <time>",
			Value: "Establish a new party instance and enable others to participate.",
		},
		{
			Name:  "$joinparty <party_id> <character_name>",
			Value: "Become a member of a pre-existing party.",
		},
		{
			Name:  "$expel <party_id> <character_name>",
			Value: "Remove a member from an existing party. To execute this command, you must either be the party creator or the Discord account that added the character.",
		},
		{
			Name:  "$invite <discord_name> <party_id>",
			Value: "Send a notification to the specified Discord account with party details.",
		},
		{
			Name:  "$deleteparty <party_id>",
			Value: "Erase a currently active party. Only the party creator can utilize this command.",
		},
	}

	utils.SendMessageWithFields(session, message.ChannelID, title, description, fields)
}
