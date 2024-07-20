package guild

import (
	"fmt"
	"log"
	"moonbeam/commands"
	"moonbeam/config"
	"moonbeam/utils"
	"sync"
	"time"
)

var currentMemberData []utils.Player
var newMemberData []utils.Player
var validMemberNames []string
var ticker *time.Ticker
var mu sync.Mutex

// Load all JSON member entries into newMemberInfo variable
// we also load all member names into validMemberName variable
func loadCurrentMemberData() error {
	cfg := config.ParseConfig()

	for _, member := range cfg.Guild.Members {
		player := config.ConvertJsonToPlayer(member)
		currentMemberData = append(currentMemberData, player)
		validMemberNames = append(validMemberNames, player.Name)
	}

	return nil
}

// generate new player data for all names included in validMemberNames
func loadNewMemberData() error {
	for _, name := range validMemberNames {
		player, err := utils.ParseCharacterJSON(name)
		if err != nil {
			fmt.Printf("Error parsing character JSON for %s: %s\n", name, err)
			return err
		}

		newMemberData = append(newMemberData, player)
	}

	return nil
}

func StartMemberUpdateTask() {
	ticker = time.NewTicker(15 * time.Minute)

	go func() {
		for range ticker.C {
			err := loadCurrentMemberData()
			if err != nil {
				log.Printf("error in loading current member data %s", err)
				clearData()
				continue
			}

			err1 := loadNewMemberData()
			if err1 != nil {
				log.Printf("error in generating new member data %s", err)
				clearData()
				continue
			}

			diff := compareMemberData()
			mu.Lock()
			if len(diff) != 0 {
				commands.CreateFeedPosts(diff)
				config.RefreshMemberList(newMemberData)
			}

			clearData()
			mu.Unlock()
		}
	}()
}

func compareMemberData() []string {
	var diffs []string = nil
	for _, currentData := range currentMemberData {
		for _, newData := range newMemberData {
			if currentData.Name == newData.Name {
				if (currentData.Cards / 50) < (newData.Cards / 50) {
					// display in multiples of 50
					diffs = append(diffs, fmt.Sprintf("%s has collected %v cards!", currentData.Name, (newData.Cards/50)*50))
				}
				if (currentData.Fame / 50) < (newData.Fame / 50) {
					// display in multiples of 50
					diffs = append(diffs, fmt.Sprintf("%s has reached %v fame!", currentData.Name, (newData.Fame/50)*50))
				}
				if currentData.Guild != newData.Guild {
					diffs = append(diffs, fmt.Sprintf("%s has left the guild!", currentData.Name))
					// remove member from list so we no longer get updates
					config.RemoveMember(currentData.Name)
				}
				if currentData.Job != newData.Job {
					diffs = append(diffs, fmt.Sprintf("%s has advanced to %s!", currentData.Name, newData.Job))
				}
				if currentData.Level != newData.Level {
					diffs = append(diffs, fmt.Sprintf("%s has reached level %v!", currentData.Name, newData.Level))

				}
				if (currentData.Quests / 50) < (newData.Quests / 50) {
					// display in multiples of 50
					diffs = append(diffs, fmt.Sprintf("%s has completed %v quests!", currentData.Name, (newData.Quests/50)*50))
				}
			}
		}
	}

	if len(diffs) != 0 {
		fmt.Printf("Differences to be posted %s", diffs)
	}
	return diffs
}

func clearData() {
	currentMemberData = nil
	newMemberData = nil
	validMemberNames = nil
}
