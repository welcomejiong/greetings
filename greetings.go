package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// Return a greeting that embeds the name in a message.
	if name == "" {
		return "", errors.New("empty name please enter your name.eg:lingxiaoxiao")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
