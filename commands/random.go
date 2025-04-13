package commands

import (
	"fmt"
	"math/rand"
	"moonbeam/utils"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func HandleListRandomizer(session *discordgo.Session, message *discordgo.MessageCreate) {
	sliceOfStrings := splitEachString(message.Content)
	randShuffle(sliceOfStrings)

	// only show the first result of the shuffled slice
	utils.SendMessage(session, message.ChannelID, "Randomized Result!", strings.Trim(fmt.Sprintf("%s", sliceOfStrings[:1]), "[ ]"))
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
