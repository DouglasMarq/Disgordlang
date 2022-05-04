package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"nitra/bot/handlers"

	"github.com/bwmarrin/discordgo"
)

var (
	TOKEN       string = os.Getenv("TOKEN")
	ENVIRONMENT string = os.Getenv("ENVIRONMENT")
)

func main() {
	discord, err := setup(TOKEN)
	defer discord.Close()

	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	if ENVIRONMENT == "dev" {
		fmt.Println("Initializing in DEV mode...")
	}

	err = handlers.Init(discord)

	if err != nil {
		discord.Close()
		return
	}

	fmt.Println("Bot is now running.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
}

func setup(token string) (*discordgo.Session, error) {
	return discordgo.New("Bot " + token)
}
