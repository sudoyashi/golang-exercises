package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("error: empty name")
	}

	// If a name was received, return a value that embeds the name
	// in a greeting message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
		"LET'S GO, %v!",
		"Time to cook, chef %v.",
	}

	// Values are indexed per line, so line 1 is the first value in the array.
	// Line 2 is the second value (meaning [1], not [2]!)
	// Introducing the math/rand module willl allow us to randomly choose values
	// Use math/rand and randomFormat together to produce random string text!
	return formats[rand.Intn(len(formats))]
}
