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
func NewParty(partyType string, time time.Time) *Party {
	party := &Party{
		ID:   idIncrement,
		Type: partyType, // todo verify
		Time: time,
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

// Get a party by its ID. Returns nil if the party is not found.
func GetPartyByID(partyID int) *Party {
	for _, party := range registeredParties {
		if party.ID == partyID {
			return party
		}
	}
	return nil // Party not found
}

func (party *Party) ShowPartyInfo() string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("Party ID: %d\n", party.ID))
	builder.WriteString(fmt.Sprintf("Party Type: %s\n", party.Type))
	builder.WriteString(fmt.Sprintf("Party Time: %s\n", party.Time))

	builder.WriteString("Party Members:\n")
	for i, member := range party.Members {
		builder.WriteString(fmt.Sprintf("Member %d: %v\n", i+1, member))
	}

	return builder.String()
}

func ShowAllParties() string {
	var builder strings.Builder

	builder.WriteString("Registered Parties:\n")
	for _, party := range registeredParties {
		builder.WriteString(party.ShowPartyInfo())
		builder.WriteString("\n") // Add an empty line between parties
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
