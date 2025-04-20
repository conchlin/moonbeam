package guild

import (
	"fmt"
	"log"
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

// Load all JSON member entries into currentMemberInfo variable
// we also load all member names into validMemberName variable
func loadCurrentMemberData() error {
	cfg := config.ParseConfig()

	// use single guilds slice to reduce repetition
	guilds := [][]config.MemberInfo{
		cfg.Guild.Moonbeam,
		cfg.Guild.Lefay,
		cfg.Guild.Basement,
		cfg.Guild.Torrent,
	}

	for _, guild := range guilds {
		for _, member := range guild {
			player := config.ConvertJsonToPlayer(member)
			currentMemberData = append(currentMemberData, player)
			validMemberNames = append(validMemberNames, player.Name)
		}
	}

	return nil
}

// generate new player data for all names included in validMemberNames
func loadNewMemberData() error {
	var foundError bool
	for _, name := range validMemberNames {
		player, err := utils.ParseCharacterJSON(name)
		if err != nil {
			fmt.Printf("Error parsing character JSON for %s: %s\n", name, err)
			foundError = true
			// go through all names before returning error
			continue
		}

		if foundError {
			return fmt.Errorf("one or more players failed to load")
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

			diff, updatedPlayers := compareMemberData()
			mu.Lock()
			if len(diff) != 0 {
				CreateFeedPosts(diff)
				FlagForUpdate(updatedPlayers)
				config.RefreshMemberList(newMemberData)
			}

			clearData()
			mu.Unlock()
		}
	}()
}

func compareMemberData() ([]Event, []string) {
	var diffs []Event
	var updatedPlayer []string = nil
	for _, currentData := range currentMemberData {
		for _, newData := range newMemberData {
			if currentData.Name == newData.Name {
				if (currentData.Cards / 50) < (newData.Cards / 50) {
					// display in multiples of 50
					diffs = append(diffs, Event{
						Achievement: fmt.Sprintf("%s has collected %v cards!", currentData.Name, (newData.Cards/50)*50),
						Guild:       newData.Guild,
					})
					updatedPlayer = append(updatedPlayer, currentData.Name)
				}
				if (currentData.Fame / 50) < (newData.Fame / 50) {
					// display in multiples of 50
					diffs = append(diffs, Event{
						Achievement: fmt.Sprintf("%s has reached %v fame!", currentData.Name, (newData.Fame/50)*50),
						Guild:       newData.Guild,
					})
					updatedPlayer = append(updatedPlayer, currentData.Name)
				}
				if currentData.Guild != newData.Guild {
					// log to console instead of discord
					fmt.Printf("%s has left the guild", currentData.Name)
					config.RemoveMember(currentData.Name, currentData.Guild)
				}
				if currentData.Job != newData.Job {
					diffs = append(diffs, Event{
						Achievement: fmt.Sprintf("%s has advanced to %s!", currentData.Name, newData.Job),
						Guild:       newData.Guild,
					})
					updatedPlayer = append(updatedPlayer, currentData.Name)
				}
				if currentData.Level != newData.Level && newData.Level >= 30 {
					// only send level up posts if the player is at least level 30
					diffs = append(diffs, Event{
						Achievement: fmt.Sprintf("%s has reached level %v!", currentData.Name, newData.Level),
						Guild:       newData.Guild,
					})
					updatedPlayer = append(updatedPlayer, currentData.Name)
				}
				if (currentData.Quests / 50) < (newData.Quests / 50) {
					// display in multiples of 50
					diffs = append(diffs, Event{
						Achievement: fmt.Sprintf("%s has completed %v quests!", currentData.Name, (newData.Quests/50)*50),
						Guild:       newData.Guild,
					})
					updatedPlayer = append(updatedPlayer, currentData.Name)
				}
			}
		}
	}

	if len(diffs) != 0 {
		fmt.Printf("Differences to be posted %s", diffs)
	}
	return diffs, updatedPlayer
}

func clearData() {
	currentMemberData = nil
	newMemberData = nil
	validMemberNames = nil
}
