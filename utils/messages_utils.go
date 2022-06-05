package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func GetMessageParameter(message string) (int, error) {
	splittedMessage := strings.Split(message, " ")

	if len(splittedMessage) < 2 {
		return 0, fmt.Errorf("please, provide a parameter")
	}

	messageParameter, err := strconv.Atoi(splittedMessage[1])
	if err != nil {
		return 0, fmt.Errorf("invalid parameter")
	}

	return messageParameter, nil
}

func GetChannelMessagesID(s *discordgo.Session, channelId string, numberOfMessages int, messageId string) []string {
	messages, err := s.ChannelMessages(channelId, numberOfMessages, messageId, "", "")
	if err != nil {
		fmt.Printf("Can't delete messages")
		return nil
	}

	var messagesId []string
	for _, message := range messages {
		messagesId = append(messagesId, message.ID)
	}

	return messagesId
}
