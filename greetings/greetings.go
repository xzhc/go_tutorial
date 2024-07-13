package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	//If not name was given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}
	//Return a greeting tha embeds the name in a message
	message := fmt.Sprintf("Hi, %v. welcome!", name)
	return message, nil
}
