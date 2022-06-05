package handlers

import (
	"fmt"
	"nitra/bot/utils"
	"strings"

	"github.com/bwmarrin/discordgo"
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
		numberOfMessagesToDelete, err := utils.GetMessageParameter(m.Content)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, err.Error())
			return
		}

		if numberOfMessagesToDelete == 0 {
			s.ChannelMessageSend(m.ChannelID, "There's no messages to delete.")
			return
		}
		if numberOfMessagesToDelete > 100 {
			s.ChannelMessageSend(m.ChannelID, "Theres a max of 100 messages per clear.")
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
		message := fmt.Sprintf("Pong! %s", s.HeartbeatLatency())
		s.ChannelMessageSend(m.ChannelID, message)
	}

	if strings.Contains(m.Content, strings.ToLower("!pong")) {
		message := fmt.Sprintf("Pong! %s", s.HeartbeatLatency())
		s.ChannelMessageSend(m.ChannelID, message)
	}
}

func helpMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := utils.IsQA(m.GuildID)
	if err != nil {
		return
	}

	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.Contains(m.Content, strings.ToLower("!help")) {
		result, err := utils.FindCommand(strings.ToLower("!help"))
		if err != nil {
			fmt.Println("oh fuck")
		}
		s.ChannelMessageSend(m.ChannelID, result.Text)
	}
}
