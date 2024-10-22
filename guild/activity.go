package guild

import (
	"moonbeam/config"
)

var playersToUpdate []string

func FlagForUpdate(players []string) {
	playersToUpdate = append(playersToUpdate, players...)

	for _, player := range players {
		config.UpdateTimestamp(player)
	}
}
