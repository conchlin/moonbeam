package commands

import (
	"errors"
	"fmt"
	"moonbeam/party"
	"moonbeam/utils"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

/**
* 	join a pre-existing party
*	discord syntax -> ~joinparty <party_id> <character_name>
*	character info is populated using legends api
 */

// command handler
func HandleJoinParty(session *discordgo.Session, message *discordgo.MessageCreate) {
	id, charName, err := splitJoinPartyString(message.Content)
	if err != nil {
		session.ChannelMessage(message.ChannelID, fmt.Sprintf("Error in joining party: %v \r\n The command syntax is $joinparty <party_id> <player_name>", err))
		return
	}

	playerInfo, _ := utils.ParseCharacterJSON(charName)

	newMember := party.NewPartyMember(message.Author, playerInfo)
	validParty := party.GetPartyByID(id)

	if validParty != nil {
		//imgUrl := fmt.Sprintf("https://maplelegends.com/api/getavatar?name=%s", playerInfo.Name)
		//imgBuf, _ := utils.ParseCharacterImage(imgUrl)
		err := validParty.AddMember(newMember)
		if err != nil {
			utils.SendErrorMessage(session, message.ChannelID, err)
			return
		}
		validParty.ShowPartyInfo()

		utils.SendMessage(session, message.ChannelID, "New Member!", "You have successfully joined the party!")
	} else {
		session.ChannelMessage(message.ChannelID, fmt.Sprintf("Error in joining party: %v The command syntax is $joinparty <player_name>", err))
		return
	}
}

// splitJoinPartyString parses the input message for the $joinparty command and extracts
// the party ID, and character name.
func splitJoinPartyString(msg string) (int, string, error) {
	msgSplit := strings.SplitAfter(msg, " ")

	if len(msgSplit) != 3 {
		return 0, "", errors.New("command parameter is missing")
	}

	idIntVal, err := strconv.Atoi(strings.TrimSpace(msgSplit[1]))
	if err != nil {
		return 0, "", fmt.Errorf("invalid party id")
	}

	characterName := strings.TrimSpace(msgSplit[2])

	return idIntVal, characterName, nil
}
