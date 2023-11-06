package commands

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func HandleListRandomizer(session *discordgo.Session, message *discordgo.MessageCreate) {
	sliceOfStrings := splitEachString(message.Content)
	randShuffle(sliceOfStrings)

	session.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
		Title: "Randomized Result!",
		// only show the first result of the shuffled slice

		Description: strings.Trim(fmt.Sprintf("%s", sliceOfStrings[:1]), "[ ]"),
		Color:       0x2cdaca,
	})
}

func splitEachString(msg string) []string {
	var a []string
	msgLower := strings.ToLower(msg)
	msgSplit := strings.SplitAfter(msgLower, " ")

	// exclude the $random command prefix
	a = append(a, msgSplit[1:]...)

	fmt.Println(a)
	return a
}

func randShuffle(a []string) {
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
}
