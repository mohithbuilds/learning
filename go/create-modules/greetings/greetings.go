package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
	// if no name is given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}
	// note how we don't throw an error, rather create and return a error object

	// if a name was given, return a greeting that embeds the name in a message
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
	// note we have to still return a nil in order to match the function signature
	// otherwise it'll fail to compile
}

// Hellos returns a map that associates each of the named people with a
// greeting message
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages
	messages := make(map[string]string)

	// loop through the slice of names and map a message
	for _, name := range names {
		message, err := Hello(name)
		// handle errs
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with the name
		messages[name] = message
	}
	return messages, nil
}

// randomFormat returns one of a set of greeting messages
// The returned message is selected randomly
func randomFormat() string {
	// A slice of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hey %v! How's it going?",
	}

	// return randomly selected message format by specifying a random index for
	// the formats slice
	return formats[rand.Intn(len(formats))]
}
