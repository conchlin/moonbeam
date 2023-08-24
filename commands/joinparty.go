package commands

import (
	"fmt"
	"strconv"
	"strings"
	"wanderlust/party"

	"github.com/bwmarrin/discordgo"
)

/**
* 	join a pre-existing party
*	discord syntax -> ~joinparty <party_id> <character_name> <job> <level>
 */

func HandleJoinParty(discord *discordgo.Session, message *discordgo.MessageCreate) {
	discord.ChannelMessageSend(message.ChannelID, joinParty(message))
}

func joinParty(message *discordgo.MessageCreate) string {
	var confirmation string

	//verify message components
	id, charName, job, level, err := splitJoinPartyString(message.Content)
	if err != nil {
		return fmt.Sprintf("There is an error in the syntax of the party command. Joining is not possible: %v", err)
	}

	newMember := party.NewPartyMember(message.Author.Username, charName, job, level)
	validParty := party.GetPartyByID(id)

	if validParty != nil {
		validParty.AddMember(newMember)
		validParty.ShowPartyInfo()
		confirmation = "You have successfully join the party!"
	} else {
		return fmt.Sprintf("There is an error in the syntax of the party command. This party id does not exist: %v", err)
	}

	return confirmation
}

func splitJoinPartyString(msg string) (int, string, string, int, error) {
	msgSplit := strings.SplitAfter(msg, " ")

	idIntVal, err := strconv.Atoi(strings.TrimSpace(msgSplit[1]))
	if err != nil {
		return 0, "", "", 0, fmt.Errorf("invalid party id")
	}

	characterName := strings.TrimSpace(msgSplit[2])
	jobName := strings.TrimSpace(msgSplit[3])

	levelIntVal, err := strconv.Atoi(strings.TrimSpace(msgSplit[4]))
	if err != nil {
		return 0, "", "", 0, fmt.Errorf("invalid job id")
	}

	fmt.Println(idIntVal)
	fmt.Println(characterName)
	fmt.Println(jobName)
	fmt.Println(levelIntVal)

	return idIntVal, characterName, jobName, levelIntVal, nil
}
