package party

import (
	"fmt"
	"strings"
	"time"
	"wanderlust/utils"

	"github.com/bwmarrin/discordgo"
)

// A party is made up of up to six party members.
// These party members are organizing for a certain PQ/boss run Type.
type Party struct {
	ID      int
	Creator string
	Type    string
	PtTime  time.Time
	Deleted bool
	Members []*PartyMember // 6 is max
}

var registeredParties []*Party
var idIncrement = 1

// PartyMember represents a member of a Party struct
type PartyMember struct {
	User       *discordgo.User
	PlayerName string
	Class      string // player's job/class
	Level      int
}

// create new party entry and add it to registeredParties
func NewParty(creator, partyType string, time time.Time) *Party {
	party := &Party{
		ID:      idIncrement,
		Creator: creator,
		Type:    partyType, // todo verify
		PtTime:  time,
		Deleted: false,
	}

	registeredParties = append(registeredParties, party)
	idIncrement++
	return party
}

// create new member entry
func NewPartyMember(discordName *discordgo.User, char utils.Player) *PartyMember {
	member := &PartyMember{
		User:       discordName,
		PlayerName: char.Name,
		Class:      char.Job,
		Level:      char.Level,
	}
	return member
}

// adds a new party member to an existing Party
func (party *Party) AddMember(member *PartyMember) error {
	if len(party.Members) >= 6 { // six is max
		return fmt.Errorf("party is already full")
	}
	party.Members = append(party.Members, member)
	return nil
}

// removes a party member from an existing Party. Only the party creator of entry
// creator can removed a character
func (party *Party) RemoveMember(originalPoster string, member *PartyMember) error {
	indexToRemove := -1

	if len(party.Members) < 2 { // at least one member needs to be remaining
		return fmt.Errorf("a party must have at least 1 member")
	}

	// let's make sure random people cant just remove members
	if originalPoster != member.User.ID && party.Creator != member.User.ID {
		return fmt.Errorf("you are not either the party creator or member who registered this character")
	}

	// Find the index of the member in the Members slice
	for i, existingMember := range party.Members {
		if existingMember == member {
			indexToRemove = i
			break
		}
	}

	if indexToRemove == -1 {
		return fmt.Errorf("member not found in the party")
	}

	// Remove the member from the Members slice
	party.Members = append(party.Members[:indexToRemove], party.Members[indexToRemove+1:]...)

	return nil
}

// delete a party from the registeredParties list. Only the original creator should
// be able to perform this action.
func DeleteParty(partyID int) bool {
	for i, party := range registeredParties {
		if party.ID == partyID {
			// Remove the party from the slice
			party.Deleted = true
			registeredParties = append(registeredParties[:i], registeredParties[i+1:]...)
			return true
		}
	}
	return false // Party not found
}

// Get a party by its ID. Returns nil if the party is not found.
func GetPartyByID(partyID int) *Party {
	for _, party := range registeredParties {
		if party.ID == partyID {
			return party
		}
	}
	return nil // Party not found
}

// Get a party member by their PlayerName. Returns nil and an error if the member is not found.
func GetPartyMemberByName(party *Party, playerName string) (*PartyMember, error) {
	for _, member := range party.Members {
		fmt.Println("member: " + member.PlayerName)
		if member.PlayerName == playerName {
			return member, nil
		}
	}
	return nil, fmt.Errorf("member not found in the party")
}

// Check if a given party ID is valid and exists within registered parties
func IsValidPartyID(partyID int) bool {
	for _, party := range registeredParties {
		if party.ID == partyID {
			return true // Party ID is valid and exists
		}
	}
	return false // Party ID is not valid or doesn't exist
}

// returns a formatted string with information about a single party and its members.
func (party *Party) ShowPartyInfo() string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("Party ID: %d\n", party.ID))
	builder.WriteString(fmt.Sprintf("Party Creator: %s\n", party.Creator))
	builder.WriteString(fmt.Sprintf("Party Type: %s\n", party.Type))
	builder.WriteString(fmt.Sprintf("Party Time: %s\n", party.PtTime))

	builder.WriteString("Party Members:\n")
	for i, member := range party.Members {
		memberSyntax := fmt.Sprintf("%s (Lvl%d %s)", member.PlayerName, member.Level, member.Class)
		builder.WriteString(fmt.Sprintf("	Member %d: %v\n", i+1, memberSyntax))
	}

	return builder.String()
}

// returns a formatted string with information about all registered parties.
func ShowAllParties() string {
	var builder strings.Builder

	builder.WriteString("Registered Parties:\n")
	for _, party := range registeredParties {
		builder.WriteString("```" + party.ShowPartyInfo() + "```")
	}

	return builder.String()
}

// parses a time input string and adjusts it to represent today's date if requested.
func ParseTimeInput(input string, today bool) (parsedTime time.Time, err error) {
	parsedTime, err = time.Parse("2006-01-02 3:04pm", input)
	if err != nil {
		return parsedTime, err
	}

	if today {
		// adjust the timestamp to represent today's date rather than a user inputted one
		parsedTime = adjustToToday(parsedTime)
	}

	return parsedTime, nil
}

// adjusts a timestamp to represent today's date.
func adjustToToday(parsedTime time.Time) time.Time {
	currentTime := time.Now().UTC()
	currentYear, currentMonth, currentDay := currentTime.Year(), currentTime.Month(), currentTime.Day()

	return time.Date(currentYear, currentMonth, currentDay, parsedTime.Hour(), parsedTime.Minute(), 0, 0, parsedTime.Location())
}

// send ping to all party members as a reminder of the party event
func BroadcastPartyMessage(session *discordgo.Session, message *discordgo.MessageCreate, party *Party) {
	var builder strings.Builder
	for _, members := range party.Members {
		builder.WriteString(members.User.Mention())
	}

	builder.WriteString(fmt.Sprintf("The %s party is about to begin!", party.Type))
	// only send message if party still exists
	if !party.Deleted {
		// since the party has been notified of the event we start a timer
		// to delete the party after 30 minutes. This makes sure the party
		// list is not just filled with long since completed events.
		go utils.CreateTimer(30*time.Minute, func() {
			DeleteParty(party.ID)
		})

		session.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
			Title:       "Event time!",
			Description: builder.String(),
			Color:       0x2cdaca,
		})
	}
}
