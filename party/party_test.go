package party

import (
	"testing"
)

func TestAddMember(t *testing.T) {
	p := NewParty(1, "opq", 1234567890)
	member := NewPartyMember("Discord1", "Player1", "Warrior")

	err := p.AddMember(member)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if len(p.Members) != 1 {
		t.Errorf("Expected party to have 1 member, but got %d", len(p.Members))
	}
}

func TestShowPartyInfo(t *testing.T) {
	p := NewParty(1, "opq", 1234567890)
	member := NewPartyMember("Discord1", "Player1", "Warrior")
	p.AddMember(member)

	expectedInfo := "Party ID: 1\nParty Type: opq\nParty Time: 1234567890\nParty Members:\nMember 1: &{Discord1 Player1 Warrior}\n"
	actualInfo := p.ShowPartyInfo()

	if actualInfo != expectedInfo {
		t.Errorf("Expected party info:\n%s\nBut got:\n%s", expectedInfo, actualInfo)
	}
}
