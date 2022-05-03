package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func GetMessageParameter(message string) int {
	splittedMessage := strings.Split(message, " ")

	if len(splittedMessage) < 2 {
		return 0
	}

	messageParameter, err := strconv.Atoi(splittedMessage[1])
	if err != nil {
		fmt.Printf("Couldn't delete messages")
		return 0
	}

	return messageParameter
}

func GetChannelMessagesID(s *discordgo.Session, channelId string, numberOfMessages int, messageId string) []string {
	messages, err := s.ChannelMessages(channelId, numberOfMessages, messageId, "", "")
	if err != nil {
		fmt.Printf("Can't delete messages")
		return nil
	}

	var messagesId []string
	for index, message := range messages {
		fmt.Printf("%d", index)
		messagesId = append(messagesId, message.ID)
	}

	return messagesId
}
