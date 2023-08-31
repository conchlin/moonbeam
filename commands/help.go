package commands

import "github.com/bwmarrin/discordgo"

/**
This command provides very simple documentation on app commands
*/

func HandleHelp(discord *discordgo.Session, message *discordgo.MessageCreate) {
	var title string
	var description string
	var fields []*discordgo.MessageEmbedField

	title = "The Wanderlust Discord App!"
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
			Name:  "$joinparty <party_id> <character_name> <class> <level>",
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

	discord.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
		Title:       title,
		Description: description,
		Fields:      fields,
		Color:       0x2cdaca,
	})
}
