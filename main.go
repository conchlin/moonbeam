package main

import (
	"fmt"
	"log"
	"moonbeam/commands"
	"moonbeam/config"
	"moonbeam/guild"
	"moonbeam/party"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func main() {
	config := config.ParseConfig()

	discord, err := discordgo.New("Bot " + config.Discord.SecurityToken)
	if err != nil {
		log.Fatal(err)
	}

	// Add event handler
	discord.AddHandler(newMessage)
	guild.StartMemberUpdateTask()

	// Open websocket
	discord.Open()

	// Run until process is terminated
	fmt.Println("Bot running...")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	//close websocket
	defer discord.Close()
}

// an event handler that processes the message and handles specific commands.
func newMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	// Ignore bot messages
	if message.Author.ID == session.State.User.ID {
		return
	}

	// ignore messages that do not start with $
	if !strings.HasPrefix(message.Content, "$") {
		return
	}

	// handle commands
	switch {
	case message.Content == "$moonbeam":
		commands.HandleHelp(session, message)
	case message.Content == "$showparties":
		party.ShowAllParties(session, message)
	case message.Content == "$docu",
		message.Content == "$documentation":
		commands.HandleDocu(session, message)
	case message.Content == "$startguildfeed":
		guild.HandleStartFeed(session, message)
	case message.Content == "$createbackup":
		commands.HandleBackupCreation(session, message)
	case message.Content == "$applybackup":
		commands.HandleApplyingBackup(session, message)
	// for commands that use additional input we need strings.Contains
	case strings.Contains(message.Content, "$guildlist"):
		commands.HandleGuildList(session, message)
	case strings.Contains(message.Content, "$createparty"):
		commands.HandleCreateParty(session, message)
	case strings.Contains(message.Content, "$joinparty"):
		commands.HandleJoinParty(session, message)
	case strings.Contains(message.Content, "$expel"):
		commands.HandleExpelMember(session, message)
	case strings.Contains(message.Content, "$deleteparty"):
		commands.HandlePartyDeletion(session, message)
	case strings.Contains(message.Content, "$invite"):
		commands.HandleInviteMember(session, message)
	case strings.Contains(message.Content, "$random"):
		commands.HandleListRandomizer(session, message)
	case strings.Contains(message.Content, "$character"):
		commands.HandleCharacterRequest(session, message)
	case strings.Contains(message.Content, "$addmember"):
		commands.HandleNewGuildMember(session, message)
	case strings.Contains(message.Content, "$removemember"):
		commands.HandleMemberRemoval(session, message)
	}
}
