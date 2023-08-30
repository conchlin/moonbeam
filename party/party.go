package party

import (
	"fmt"
	"strings"
	"time"
)

// A party is made up of up to six party members.
// These party members are organizing for a certain PQ/boss run Type.
type Party struct {
	ID      int
	Creator string
	Type    string
	Time    time.Time
	Members []*PartyMember // 6 is max
}

var registeredParties []*Party
var idIncrement = 1

type PartyMember struct {
	DiscordName string
	PlayerName  string
	Class       string // player's job/class
	Level       int
}

// create new party entry
func NewParty(creator, partyType string, time time.Time) *Party {
	party := &Party{
		ID:      idIncrement,
		Creator: creator,
		Type:    partyType, // todo verify
		Time:    time,
	}

	registeredParties = append(registeredParties, party)
	idIncrement++
	return party
}

// create new member entry
func NewPartyMember(discordName, playerName, class string, level int) *PartyMember {
	member := &PartyMember{
		DiscordName: discordName,
		PlayerName:  playerName,
		Class:       class,
		Level:       level,
	}
	return member
}

// adds a new party member to the Party struct
func (party *Party) AddMember(member *PartyMember) error {
	if len(party.Members) >= 6 { // six is max
		return fmt.Errorf("party is already full")
	}
	party.Members = append(party.Members, member)
	return nil
}

func (party *Party) RemoveMember(originalPoster string, member *PartyMember) error {
	indexToRemove := -1

	if len(party.Members) < 2 { // at least one member needs to be remaining
		return fmt.Errorf("a party must have at least 1 member")
	}

	// let's make sure random people cant just remove members
	if originalPoster != member.DiscordName && party.Creator != member.DiscordName {
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

func (party *Party) ShowPartyInfo() string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("Party ID: %d\n", party.ID))
	builder.WriteString(fmt.Sprintf("Party Creator: %s\n", party.Creator))
	builder.WriteString(fmt.Sprintf("Party Type: %s\n", party.Type))
	builder.WriteString(fmt.Sprintf("Party Time: %s\n", party.Time))

	builder.WriteString("Party Members:\n")
	for i, member := range party.Members {
		memberSyntax := fmt.Sprintf("%s (Lvl%d %s)", member.PlayerName, member.Level, member.Class)
		builder.WriteString(fmt.Sprintf("	Member %d: %v\n", i+1, memberSyntax))
	}

	return builder.String()
}

func ShowAllParties() string {
	var builder strings.Builder

	builder.WriteString("Registered Parties:\n")
	for _, party := range registeredParties {
		builder.WriteString("```" + party.ShowPartyInfo() + "```")
	}

	return builder.String()
}

func ParseTimeInput(input string) (time.Time, error) {
	input = strings.ReplaceAll(strings.ToLower(input), " ", "")
	timeLayout := "3:04pm"

	parsedTime, err := time.Parse(timeLayout, input)
	if err != nil {
		return time.Time{}, err
	}

	return parsedTime, nil
}
