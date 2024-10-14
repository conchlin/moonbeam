package guild

import (
	"moonbeam/config"
	"time"
)

var playersToUpdate []string

func FlagForUpdate(players []string) {
	playersToUpdate = append(playersToUpdate, players...)

	for _, player := range players {
		updateTimestamp(player)
	}
}

func updateTimestamp(flaggedChar string) {
	config := config.ParseConfig()

	for i, player := range config.Activity.Updated {
		if player.Name == flaggedChar {
			config.Activity.Updated[i].UpdateAt = time.Now()
		}
	}
}
