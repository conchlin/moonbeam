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

var commandHandlers = map[string]func(session *discordgo.Session, message *discordgo.MessageCreate){
	"$crossroads":     commands.HandleHelp,
	"$showparties":    party.ShowAllParties,
	"$docu":           commands.HandleDocu,
	"$documentation":  commands.HandleDocu,
	"$startguildfeed": guild.HandleStartFeed,
	"$createbackup":   commands.HandleBackupCreation,
	"$applybackup":    commands.HandleApplyingBackup,
	"$guildlist":      commands.HandleGuildList,
	"$createparty":    commands.HandleCreateParty,
	"$joinparty":      commands.HandleJoinParty,
	"$expel":          commands.HandleExpelMember,
	"$deleteparty":    commands.HandlePartyDeletion,
	"$invite":         commands.HandleInviteMember,
	"$random":         commands.HandleListRandomizer,
	"$character":      commands.HandleCharacterRequest,
	"$addmember":      commands.HandleNewGuildMember,
	"$removemember":   commands.HandleMemberRemoval,
}

func main() {
	config := config.ParseConfig()

	discord, err := discordgo.New("Bot " + config.Discord.SecurityToken)
	if err != nil {
		log.Fatal(err)
	}

	discord.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsGuildMembers

	// handlers
	discord.AddHandler(memberJoinHandler)
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

	fields := strings.Fields(message.Content)
	if len(fields) == 0 {
		return
	}

	command := fields[0]
	handler, exists := commandHandlers[command]
	if exists {
		handler(session, message)
	}
}

func memberJoinHandler(session *discordgo.Session, newMember *discordgo.GuildMemberAdd) {
	var welcomeMessage strings.Builder
	welcomeChannelID := "1191937521584721933"

	welcomeMessage.WriteString(fmt.Sprintf("Welcome to the Crossroads Alliance Discord, <@%s>!\n\n", newMember.User.ID))
	welcomeMessage.WriteString("To help us keep things organized, please update your server nickname to one of the following formats: ``Name (IGN)`` or ``IGN``\n")
	welcomeMessage.WriteString("You can also introduce your character by using ``$character <ign>``")

	_, err := session.ChannelMessageSend(welcomeChannelID, welcomeMessage.String())
	if err != nil {
		log.Printf("Error sending welcome message: %v", err)
	}
}
