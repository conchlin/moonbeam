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
func loadCurrentMemberData() ([]utils.Player, error) {
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
	ticker = time.NewTicker(15 * time.Minute)

	go func() {
		for range ticker.C {
			loadCurrentMemberData()
			loadNewMemberData()
			diff := compareMemberData()
			commands.CreateFeedPosts(diff)
			config.RefreshMemberList(newMemberData)
			// clear data for next iteration
			currentMemberData = nil
			newMemberData = nil
			validMemberNames = nil
		}
	}()
}

func compareMemberData() []string {
	var diffs []string = nil
	for _, currentData := range currentMemberData {
		for _, newData := range newMemberData {
			if currentData.Name == newData.Name {
				// card notifications should happen every 50 cards
				/*if currentData.Cards != newData.Cards && newData.Cards-currentData.Cards >= 50 {
					remainder := newData.Cards % 50
					roundedNumber := newData.Cards - remainder
					diffs = append(diffs, fmt.Sprintf("%s has reached %v cards!", currentData.Name, roundedNumber))
				}*/
				// fame notifications should happen every 25 fame
				/*if currentData.Fame != newData.Fame && newData.Fame-currentData.Fame >= 25 {
					remainder := newData.Fame % 25
					roundedNumber := newData.Fame - remainder
					diffs = append(diffs, fmt.Sprintf("%s has reached %v fame!", currentData.Name, roundedNumber))
				}*/
				if currentData.Guild != newData.Guild {
					diffs = append(diffs, fmt.Sprintf("%s has left the guild!", currentData.Name))
				}
				if currentData.Job != newData.Job {
					diffs = append(diffs, fmt.Sprintf("%s has advanced to %s!", currentData.Name, newData.Job))
				}
				if currentData.Level != newData.Level {
					diffs = append(diffs, fmt.Sprintf("%s has reached level %v!", currentData.Name, newData.Level))

				}
				/*if currentData.Quests != newData.Quests && newData.Fame-currentData.Fame >= 25 {
					remainder := newData.Quests % 25
					roundedNumber := newData.Quests - remainder
					diffs = append(diffs, fmt.Sprintf("%s has completed %v quests!", currentData.Name, roundedNumber))
				}*/
			}
		}
	}

	return diffs
}
