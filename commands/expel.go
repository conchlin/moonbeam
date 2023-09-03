package commands

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"wanderlust/party"

	"github.com/bwmarrin/discordgo"
)

// handles the expulsion of a member from a party.
// syntax for command is $removemember <party_id> <character_name>

// command handler
func HandleExpelMember(session *discordgo.Session, message *discordgo.MessageCreate) {
	authorName := message.Author.Username

	partyId, playerName, strErr := splitRemoveMemberString(message.Content)
	if strErr != nil {
		session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Error in removing member: %v \r\n The command syntax should be $expel <player_name> <party_id>", strErr))
		return
	}
	partyInstance := party.GetPartyByID(partyId)
	player, err := party.GetPartyMemberByName(partyInstance, playerName)
	if err != nil {
		session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("%s cannot be found in party", playerName))
		return
	}

	removeErr := partyInstance.RemoveMember(GetNickname(session, authorName), player)
	if removeErr != nil {
		session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("error in removing member: %v", removeErr))
		return
	}

	session.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
		Title:       "Player Removed",
		Description: player.PlayerName + " has been removed from the party.",
		Color:       0x2cdaca,
	})
}

// splitRemoveMemberString parses the input message for the $expel command and extracts
// the party ID and player name to be expelled.
func splitRemoveMemberString(msg string) (int, string, error) {
	msgSplit := strings.SplitAfter(msg, " ")

	if len(msgSplit) != 3 {
		return 0, "", errors.New("command parameter is missing")
	}

	playerName := strings.TrimSpace(msgSplit[1])

	idIntVal, err := strconv.Atoi(strings.TrimSpace(msgSplit[2]))
	if err != nil {
		return 0, "", fmt.Errorf("invalid party id")
	}

	if !party.IsValidPartyID(idIntVal) {
		return 0, "", fmt.Errorf("this party id does not exist")
	}

	return idIntVal, playerName, nil
}
