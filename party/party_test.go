package party

import (
	"testing"
	"wanderlust/utils"

	"github.com/bwmarrin/discordgo"
)

func createTestUser(username string) *discordgo.User {
	return &discordgo.User{
		Username: username,
	}
}

func TestAddMember(t *testing.T) {
	partyTime, _ := ParseTimeInput("6:00pm", true)
	p := NewParty("opq", "PvE", partyTime)

	// Create a discordgo.User to use as input
	user := createTestUser("Discord1")
	char := utils.Player{Name: "Player1", Job: "Warrior", Level: 120}
	member := NewPartyMember(user, char)

	err := p.AddMember(member)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if len(p.Members) != 1 {
		t.Errorf("Expected party to have 1 member, but got %d", len(p.Members))
	}

	if p.Members[0].PlayerName != "Player1" {
		t.Errorf("Expected member's PlayerName to be Player1, but got %s", p.Members[0].PlayerName)
	}
}

func TestRemoveMember(t *testing.T) {
	partyTime, _ := ParseTimeInput("6:00pm", true)
	p := NewParty("opq", "PvE", partyTime)

	// Create a discordgo.User to use as input
	user := createTestUser("Discord1")
	char := utils.Player{Name: "Player1", Job: "Warrior", Level: 120}
	member := NewPartyMember(user, char)
	p.AddMember(member)

	err := p.RemoveMember("opq", member)
	if err != nil {
		t.Errorf("Error removing party member: %v", err)
	}
	if len(p.Members) != 0 {
		t.Errorf("Expected 0 party members, but got %d", len(p.Members))
	}
}

func TestDeleteParty(t *testing.T) {
	partyTime, _ := ParseTimeInput("6:00pm", true)
	p := NewParty("opq", "PvE", partyTime)

	partyID := p.ID
	if !DeleteParty(partyID) {
		t.Errorf("Failed to delete party with ID %d", partyID)
	}
	if len(registeredParties) != 0 {
		t.Errorf("Expected 0 registered parties, but got %d", len(registeredParties))
	}
}

func TestGetPartyByID(t *testing.T) {
	partyTime, _ := ParseTimeInput("6:00pm", true)
	p := NewParty("opq", "PvE", partyTime)

	partyID := p.ID
	p2 := GetPartyByID(partyID)
	if p2 != nil {
		t.Errorf("Expected GetPartyByID to return nil, but got a party with ID %d", p2.ID)
	}
}

func TestGetPartyMemberByName(t *testing.T) {
	partyTime, _ := ParseTimeInput("6:00pm", true)
	p := NewParty("opq", "PvE", partyTime)

	// Create a discordgo.User to use as input
	user := createTestUser("Discord1")
	char := utils.Player{Name: "Player1", Job: "Warrior", Level: 120}
	member := NewPartyMember(user, char)
	p.AddMember(member)

	memberName := "Player1"
	member2, err := GetPartyMemberByName(p, memberName)
	if member2 != nil || err == nil {
		t.Errorf("Expected GetPartyMemberByName to return nil and an error, but got member: %v, error: %v", member2, err)
	}
}

func TestParseTimeInput(t *testing.T) {
	timeStr := "2023-10-28 12:00pm"
	parsedTime, err := ParseTimeInput(timeStr, false)
	if err != nil {
		t.Errorf("Error parsing time input: %v", err)
	}
	if parsedTime.Format("2006-01-02 3:04pm") != timeStr {
		t.Errorf("Parsed time does not match the input: expected %s, but got %s", timeStr, parsedTime.Format("2006-01-02 3:04pm"))
	}
}

func TestAdjustToToday(t *testing.T) {
	partyTime, _ := ParseTimeInput("6:00pm", true)
	todayTimeStr := "12:00pm"
	todayTime := adjustToToday(partyTime)
	parsedTodayTime, err := ParseTimeInput(todayTimeStr, true)
	if err != nil {
		t.Errorf("Error parsing today's time input: %v", err)
	}
	if parsedTodayTime.Format("2006-01-02 3:04pm") != todayTime.Format("2006-01-02 3:04pm") {
		t.Errorf("Parsed today's time does not match the adjusted time: expected %s, but got %s", todayTime.Format("2006-01-02 3:04pm"), parsedTodayTime.Format("2006-01-02 3:04pm"))
	}
}
