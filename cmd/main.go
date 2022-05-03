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
	discord, err := discordgo.New("Bot " + TOKEN)
	defer discord.Close()
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	if ENVIRONMENT == "dev" {
		fmt.Println("Initializing in DEV mode...")
	}

	discord.UpdateGameStatus(int(discordgo.ActivityTypeGame), "Testing!")

	handlers.Init(discord)

	fmt.Println("Bot is now running.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	discord.Close()
}
