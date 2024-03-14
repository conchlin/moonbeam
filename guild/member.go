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
func loadNewMemberData() {
	for _, name := range validMemberNames {
		player, err := utils.ParseCharacterJSON(name)
		if err != nil {
			fmt.Printf("Error parsing character JSON for %s: %s\n", name, err)
			continue
		}

		newMemberData = append(newMemberData, player)
	}
}

func StartMemberUpdateTask() {
	ticker = time.NewTicker(15 * time.Minute)

	go func() {
		for range ticker.C {
			err := loadCurrentMemberData()
			if err != nil {
				log.Printf("error in loading current member data %s", err)
				continue
			}
			loadNewMemberData()
			diff := compareMemberData()

			mu.Lock()
			commands.CreateFeedPosts(diff)
			config.RefreshMemberList(newMemberData)
			// clear data for next iteration
			currentMemberData = nil
			newMemberData = nil
			validMemberNames = nil
			mu.Unlock()
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
					// remove member from list so we no longer get updates
					config.RemoveMember(currentData.Name)
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

	fmt.Printf("Differences to be posted %s", diffs)
	return diffs
}
