package commands

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"wanderlust/party"
	"wanderlust/utils"

	"github.com/bwmarrin/discordgo"
)

/**
* 	join a pre-existing party
*	discord syntax -> ~joinparty <party_id> <character_name> <job> <level>
 */

// command handler
func HandleJoinParty(session *discordgo.Session, message *discordgo.MessageCreate) {
	id, charName, job, level, err := splitJoinPartyString(message.Content)
	if err != nil {
		session.ChannelMessage(message.ChannelID, fmt.Sprintf("Error in joining party: %v \r\n The command syntax is $joinparty <party_id> <player_name> <job> <level>", err))
		return
	}

	nick := utils.GetNickname(session, message.Author.ID)
	newMember := party.NewPartyMember(nick, charName, job, level)
	validParty := party.GetPartyByID(id)

	if validParty != nil {
		validParty.AddMember(newMember)
		validParty.ShowPartyInfo()
		session.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
			Title:       "New Member!",
			Description: "You have successfully joined the party!",
			Color:       0x2cdaca,
		})
	} else {
		session.ChannelMessage(message.ChannelID, fmt.Sprintf("Error in joining party: %v The command syntax is $joinparty <player_name> <job> <level>", err))
		return
	}
}

// splitJoinPartyString parses the input message for the $joinparty command and extracts
// the party ID, character name, job, and level.
func splitJoinPartyString(msg string) (int, string, string, int, error) {
	msgSplit := strings.SplitAfter(msg, " ")

	if len(msgSplit) != 5 {
		return 0, "", "", 0, errors.New("command parameter is missing")
	}

	idIntVal, err := strconv.Atoi(strings.TrimSpace(msgSplit[1]))
	if err != nil {
		return 0, "", "", 0, fmt.Errorf("invalid party id")
	}

	characterName := strings.TrimSpace(msgSplit[2])
	jobName := strings.TrimSpace(msgSplit[3])

	levelIntVal, err := strconv.Atoi(strings.TrimSpace(msgSplit[4]))
	if err != nil {
		return 0, "", "", 0, fmt.Errorf("invalid job string")
	}

	fmt.Println(idIntVal)
	fmt.Println(characterName)
	fmt.Println(jobName)
	fmt.Println(levelIntVal)

	return idIntVal, characterName, jobName, levelIntVal, nil
}
