package lottery

import (
	"moonbeam/utils"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type Lottery struct {
	BuyInAmount int
	Contestants []*Contestant
}

var activeLotto *Lottery

func CreateLottery(mesos int) *Lottery {
	lotto := &Lottery{
		BuyInAmount: mesos,
	}

	activeLotto = lotto
	return lotto
}

func GetActiveLottery() *Lottery {
	return activeLotto
}

func (lotto *Lottery) AddContestant(player *Contestant) {
	lotto.Contestants = append(lotto.Contestants, player)
}

// returns list of all contestants in the active lottery
func ShowLottery(session *discordgo.Session, message *discordgo.MessageCreate) {
	var builder strings.Builder
	activeLotto := GetActiveLottery()
	buyin := strconv.Itoa(activeLotto.BuyInAmount)

	builder.WriteString("Buy-In Amount: " + buyin)
	for _, players := range activeLotto.Contestants {
		builder.WriteString("```" + players.Name + "```")
	}

	utils.SendMessage(session, message.ChannelID, "Current Lottery Contestants", builder.String())
}
