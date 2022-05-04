package main

import (
	"os"
	"testing"
)

var (
	TESTTOKEN string = os.Getenv("TOKEN")
)

func TestDiscordInitialization(t *testing.T) {
	result, err := setup(TESTTOKEN)

	if err != nil {
		t.Errorf("Error while instantiating a discord session: %s", err)
	}

	if result == nil {
		t.Errorf("Error while initializing discord bot")
	}
}
