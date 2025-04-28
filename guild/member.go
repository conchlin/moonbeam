package guild

import (
	"fmt"
	"log"
	"moonbeam/config"
	"moonbeam/utils"
	"sync"
	"time"
)

// reduce the amount of global variables with MemberData struct
type MemberData struct {
	CurrentData []utils.Player
	NewData     []utils.Player
	ValidNames  []string
	mu          sync.Mutex
}

var ticker *time.Ticker

// Load all JSON member entries into currentMemberInfo variable
// we also load all member names into validMemberName variable
func (md *MemberData) loadCurrentMemberData() error {
	cfg := config.ParseConfig()

	// use single guilds slice to reduce repetition
	guilds := [][]config.MemberInfo{
		cfg.Guild.Moonbeam,
		cfg.Guild.Lefay,
		cfg.Guild.Basement,
	}

	for _, guild := range guilds {
		for _, member := range guild {
			player := config.ConvertJsonToPlayer(member)
			md.CurrentData = append(md.CurrentData, player)
			md.ValidNames = append(md.ValidNames, player.Name)
		}
	}

	return nil
}

// generate new player data for all names included in validMemberNames
func (md *MemberData) loadNewMemberData() error {
	var foundError bool
	for _, name := range md.ValidNames {
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

		md.NewData = append(md.NewData, player)
	}

	return nil
}

func StartMemberUpdateTask() {
	ticker = time.NewTicker(15 * time.Minute)
	memberData := &MemberData{}

	go func() {
		for range ticker.C {
			memberData.mu.Lock()

			if err := memberData.loadCurrentMemberData(); err != nil {
				log.Printf("error in loading current member data: %s", err)
				memberData.mu.Unlock()
				continue
			}

			if err := memberData.loadNewMemberData(); err != nil {
				log.Printf("error in generating new member data: %s", err)
				memberData.mu.Unlock()
				continue
			}

			diff, updatedPlayers := memberData.compareMemberData()
			if len(diff) != 0 {
				CreateFeedPosts(diff)
				FlagForUpdate(updatedPlayers)
				config.RefreshMemberList(memberData.NewData)
			}

			memberData.clearData()
			memberData.mu.Unlock()
		}
	}()
}

func (md *MemberData) compareMemberData() ([]Event, []string) {
	var diffs []Event
	var updatedPlayer []string = nil
	for _, currentData := range md.CurrentData {
		for _, newData := range md.NewData {
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

func (md *MemberData) clearData() {
	md.CurrentData = nil
	md.NewData = nil
	md.ValidNames = nil
}
