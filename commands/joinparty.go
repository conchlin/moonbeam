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

func HandleJoinParty(session *discordgo.Session, message *discordgo.MessageCreate) {
	session.ChannelMessageSend(message.ChannelID, joinParty(session, message))
}

func joinParty(session *discordgo.Session, message *discordgo.MessageCreate) string {
	var confirmation string

	//verify message components
	id, charName, job, level, err := splitJoinPartyString(message.Content)
	if err != nil {
		return fmt.Sprintf("Error in joining party: %v \r\n The command syntax is $joinparty <player_name> <job> <level>", err)
	}

	nick := getNickname(session, message.Author.ID)
	newMember := party.NewPartyMember(nick, charName, job, level)
	validParty := party.GetPartyByID(id)

	if validParty != nil {
		validParty.AddMember(newMember)
		validParty.ShowPartyInfo()
		confirmation = "You have successfully join the party!"
	} else {
		return fmt.Sprintf("Error in joining party: %v The command syntax is $joinparty <player_name> <job> <level>", err)
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
		return 0, "", "", 0, fmt.Errorf("invalid job string")
	}

	fmt.Println(idIntVal)
	fmt.Println(characterName)
	fmt.Println(jobName)
	fmt.Println(levelIntVal)

	return idIntVal, characterName, jobName, levelIntVal, nil
}

/*
*
get nickname of the user
guildId of the server
userId of the user who posted
*/
func getNickname(session *discordgo.Session, userId string) string {
	guildId := "0" // use actual guild id i just didnt want to commit it to github lol
	member, err := session.GuildMember(guildId, userId)
	if err != nil {
		fmt.Println("Error fetching member data:", err)
		return ""
	}

	nickname := member.Nick
	if nickname == "" {
		nickname = member.User.Username
	}

	return nickname
}
