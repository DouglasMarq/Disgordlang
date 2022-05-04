package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Init(s *discordgo.Session) error {
	s.AddHandler(deleteMessages)
	s.AddHandler(messagePing)

	s.Identify.Intents = discordgo.IntentsGuildMessages

	err := s.Open()

	if err != nil {
		fmt.Print("Error opening connection", err)
		return err
	}

	return err
}
