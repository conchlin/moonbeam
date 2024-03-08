package guild

import (
	"fmt"
	"moonbeam/config"
	"moonbeam/utils"
	"time"
)

var currentMemberData []utils.Player
var newMemberData []utils.Player
var validMemberNames []string
var ticker *time.Ticker

// Load all JSON member entries into newMemberInfo variable
// we also load all member names into validMemberName variable
func LoadCurrentMemberData() ([]utils.Player, error) {
	cfg := config.ParseConfig()

	for _, member := range cfg.Guild.Members {
		player := config.ConvertJsonToPlayer(member)
		currentMemberData = append(currentMemberData, player)
		validMemberNames = append(validMemberNames, player.Name)
	}

	return currentMemberData, nil
}

// generate new player data for all names included in validMemberNames
func loadNewMemberData() {
	for _, name := range validMemberNames {
		player, err := utils.ParseCharacterJSON(name)
		if err != nil {
			return
		}

		newMemberData = append(newMemberData, player)
		// look for differences in memberdata
		// populate strings that describe differences
		// replace currentMemberData with newMemberData
		// write newMemberData to json
	}
}

func StartMemberUpdateTask() {
	ticker = time.NewTicker(15 * time.Minute)

	fmt.Println("new update timer has started")
	go func() {
		for range ticker.C {
			loadNewMemberData()
		}
	}()
}

/*func compareMemberData() {
	for _, currentData := range currentMemberData {
		for _, newData := range newMemberData {
			if currentData.Name == newData.Name {
				// compare here
			}
		}
	}
}*/
