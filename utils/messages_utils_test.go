package utils

import "testing"

func TestMessageParemeter(t *testing.T) {

	m := "!clear 50"

	result := GetMessageParameter(m)

	if result < 1 {
		t.Errorf("Expected number of messages to not be 0, but got %v", result)
	}
}
