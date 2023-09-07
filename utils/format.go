package utils

import (
	"fmt"
	"wanderlust/config"

	"github.com/bwmarrin/discordgo"
)

// GetNickname retrieves the nickname of a user in a Discord server based on UserId and guildId
func GetNickname(session *discordgo.Session, userId string) string {
	config := config.ParseConfig()
	member, err := session.GuildMember(config.Discord.GuildID, userId)
	if err != nil {
		fmt.Println("Error fetching member data:", err)
		return ""
	}

	nickname := member.Nick
	if nickname == "" {
		nickname = member.User.Username
	}

	return nickname
}
