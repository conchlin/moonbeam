package guild

import (
	"fmt"
	"moonbeam/commands"
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
	}
}

func StartMemberUpdateTask() {
	ticker = time.NewTicker(1 * time.Minute)

	fmt.Println("new update timer has started")
	go func() {
		for range ticker.C {
			loadNewMemberData()
			compareMemberData()
			// replace currentMemberData with newMemberData
			// write newMemberData to json
		}
	}()
}

func compareMemberData() {
	var diffs []string = nil
	for _, currentData := range currentMemberData {
		for _, newData := range newMemberData {
			if currentData.Name == newData.Name {
				if currentData.Cards != newData.Cards {
					diffs = append(diffs, fmt.Sprintf("%s has reached %v cards!", currentData.Name, newData.Cards))
				}
				if currentData.Fame != newData.Fame {
					diffs = append(diffs, fmt.Sprintf("%s has reached %v fame!", currentData.Name, newData.Fame))
				}
				if currentData.Guild != newData.Guild {
					diffs = append(diffs, fmt.Sprintf("%s has left the guild!", currentData.Name))
				}
				if currentData.Job != newData.Job {
					diffs = append(diffs, fmt.Sprintf("%s has advanced to %s!", currentData.Name, newData.Job))
				}
				if currentData.Level != newData.Level {
					diffs = append(diffs, fmt.Sprintf("%s has reached level %v!", currentData.Name, newData.Level))
				}
				if currentData.Quests != newData.Quests {
					diffs = append(diffs, fmt.Sprintf("%s has completed %v quests!", currentData.Name, newData.Quests))
				}
			}
		}
	}

	commands.CreateFeedPosts(diffs)
}
