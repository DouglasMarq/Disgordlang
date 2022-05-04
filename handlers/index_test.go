package handlers

import (
	"os"
	"testing"

	"github.com/bwmarrin/discordgo"
)

var token string = os.Getenv("TOKEN")

func TestHandlersInitialization(t *testing.T) {
	discord, err := discordgo.New("Bot " + token)
	defer discord.Close()

	if err != nil {
		t.Errorf("Error: %s", err)
	}

	err = Init(discord)

	if err != nil {
		t.Errorf("Error opening connection: %s", err)
	}
}
