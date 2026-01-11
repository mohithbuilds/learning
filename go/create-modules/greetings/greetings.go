package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	// if no name is given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}
	// note how we don't throw an error, rather create and return a error object

	// if a name was given, return a greeting that embeds the name in a message
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
	// note we have to still return a nil in order to match the function signature
	// otherwise it'll fail to compile
}
