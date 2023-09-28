package main

import (
	"fmt"
	"log"

	// External
	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting messsage
	message, err := greetings.Hello("Joshua")
	// If there is an error, print error to console and exit.
	if err != nil {
		log.Fatal(err)
	}

	// If no error, get a greeting message and print it.
	fmt.Println(message)
}
