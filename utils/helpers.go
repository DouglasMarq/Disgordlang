package utils

import (
	"fmt"
	"os"
	"strings"
)

var (
	ENVIRONMENT string = os.Getenv("ENVIRONMENT")
	GUILDID     string = os.Getenv("GUILDID")
)

func IsQA(dGuildId string) (bool, error) {
	if ENVIRONMENT == "dev" &&
		strings.TrimSpace(GUILDID) == "" {
		fmt.Printf("ENVIRONMENT is %s, expected GUILDID: %s", ENVIRONMENT, GUILDID)
		return true, fmt.Errorf("ENVIRONMENT is %s, expected GUILDID: %s", ENVIRONMENT, GUILDID)
	}

	if dGuildId != GUILDID && ENVIRONMENT == "dev" {
		return true, nil
	}

	return false, nil
}
