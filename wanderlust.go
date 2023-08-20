package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

const securityToken = ""

func Init() {
	discord, err := discordgo.New("Bot " + securityToken)
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

func newMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {

	// Ignore bot messages
	if message.Author.ID == discord.State.User.ID {
		return
	}

	// Respond to messages
	switch {
	case message.Content == "~wanderlust":
		discord.ChannelMessageSend(message.ChannelID, "response")
	}
}
