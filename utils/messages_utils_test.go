package utils

import "testing"

var (
	correctMessage     string = "!clear 50"
	incorrectMessage   string = "!clear dsadsa"
	noArgumentsMessage string = "!clear"
)

func TestCorrectlyMessageParameter(t *testing.T) {
	result := GetMessageParameter(correctMessage)

	if result < 1 {
		t.Errorf("Expected number of messages to not be 0, but got %v", result)
	}
}

func TestIncorrectlyMessageParameter(t *testing.T) {
	result := GetMessageParameter(incorrectMessage)

	if result != 0 {
		t.Errorf("Expected number of messages to be 0, but got %v", result)
	}
}

func TestMessageWithNoArguments(t *testing.T) {
	result := GetMessageParameter(noArgumentsMessage)

	if result != 0 {
		t.Errorf("Expected number of messages to be 0, but got %v", result)
	}
}
