package party

import (
	"fmt"
	"strings"
)

// A party is made up of up to six party members.
// These party members are organizing for a certain PQ/boss run Type.
type Party struct {
	ID      int
	Type    string
	Time    int64
	Members []*PartyMember // 6 is max
}

type PartyMember struct {
	DiscordName string
	PlayerName  string
	Class       string // player's job/class
}

// create new party entry
func NewParty(id int, partyType string, time int64) *Party {
	party := &Party{
		ID:   id,        // todo auto increment
		Type: partyType, // todo varify
		Time: time,
	}
	return party
}

// create new member entry
func NewPartyMember(discordName, playerName, class string) *PartyMember {
	member := &PartyMember{
		DiscordName: discordName,
		PlayerName:  playerName,
		Class:       class,
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

func (party *Party) ShowPartyInfo() string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("Party ID: %d\n", party.ID))
	builder.WriteString(fmt.Sprintf("Party Type: %s\n", party.Type))
	builder.WriteString(fmt.Sprintf("Party Time: %d\n", party.Time))

	builder.WriteString("Party Members:\n")
	for i, member := range party.Members {
		builder.WriteString(fmt.Sprintf("Member %d: %v\n", i+1, member))
	}

	return builder.String()
}
