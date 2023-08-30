package party

import (
	"testing"
	"time"
)

func TestAddMember(t *testing.T) {
	p := NewParty("opq", "PvE", time.Now())
	member := NewPartyMember("Discord1", "Player1", "Warrior", 120)

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

func TestShowPartyInfo(t *testing.T) {
	partyTime, _ := ParseTimeInput("6:00pm")
	p := NewParty("opq", "PvE", partyTime)
	member := NewPartyMember("Discord1", "Player1", "Warrior", 120)
	p.AddMember(member)

	expectedInfo := "Party ID: 1\nParty Creator: opq\nParty Type: PvE\nParty Time: 6:00pm\nParty Members:\n	Member 1: Player1 (Lvl120 Warrior)\n"
	actualInfo := p.ShowPartyInfo()

	if actualInfo != expectedInfo {
		t.Errorf("Expected party info:\n%s\nBut got:\n%s", expectedInfo, actualInfo)
	}
}
