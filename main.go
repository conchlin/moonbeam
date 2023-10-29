package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"wanderlust/commands"
	"wanderlust/config"
	"wanderlust/party"

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
	case message.Content == "$wanderlust":
		commands.HandleHelp(session, message)
	case message.Content == "$showparties":
		session.ChannelMessageSend(message.ChannelID, party.ShowAllParties())
	case message.Content == "$docu",
		message.Content == "$documentation":
		commands.HandleDocu(session, message)
	// for commands that use additional input we need strings.Contains
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
	}
}
