package handlers

import (
	"nitra/bot/utils"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	ENVIRONMENT string = os.Getenv("ENVIRONMENT")
	GUILDID     string = os.Getenv("GUILDID")
)

func deleteMessages(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := utils.IsQA(m.GuildID)
	if err != nil {
		return
	}

	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.Contains(m.Content, strings.ToLower("!clear")) {

		numberOfMessagesToDelete := utils.GetMessageParameter(m.Content)

		if numberOfMessagesToDelete == 0 {
			s.ChannelMessageSend(m.ChannelID, "There's no messages to delete.")
			return
		}

		messagesId := utils.GetChannelMessagesID(s, m.ChannelID, numberOfMessagesToDelete, m.ID)

		s.ChannelMessageDelete(m.ChannelID, m.ID)
		s.ChannelMessagesBulkDelete(m.ChannelID, messagesId)
	}
}

func messagePing(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := utils.IsQA(m.GuildID)
	if err != nil {
		return
	}

	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.Contains(m.Content, strings.ToLower("!ping")) {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	if strings.Contains(m.Content, strings.ToLower("!pong")) {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
