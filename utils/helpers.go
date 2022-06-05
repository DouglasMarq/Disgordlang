package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type commands struct {
	Commands []command `json:"commands"`
}
type command struct {
	Command     string `json:"command"`
	Description string `json:"description"`
	Text        string `json:"text"`
	UserLevel   string `json:"userLevel"`
}

var (
	ENVIRONMENT string = os.Getenv("ENVIRONMENT")
	GUILDID     string = os.Getenv("GUILDID")
	jsonFile    *os.File
	err         error
	botCommands commands
	commandsMap = make(map[string]int)
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

func FindCommand(c string) (command, error) {
	if jsonFile == nil {
		loadJsonFile()
	}
	defer jsonFile.Close()

	if element, ok := commandsMap[c]; ok {
		return botCommands.Commands[element], nil
	}

	byteValue, err := ioutil.ReadAll(jsonFile)
	// TODO - Do something about this error
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(byteValue, &botCommands)

	//caching command into a map
	for i, element := range botCommands.Commands {
		commandsMap[element.Command] = i
	}

	for _, element := range botCommands.Commands {
		if element.Command == c {
			return element, nil
		}
	}

	return command{}, fmt.Errorf("command %s not found", c)
}

func loadJsonFile() *os.File {
	currentPath, _ := os.Getwd()
	jsonFile, err = os.Open(fmt.Sprintf("%s/handlers/commands.json", currentPath))
	// TODO - Do something about this error
	if err != nil {
		fmt.Println(err)
	}

	return jsonFile
}
