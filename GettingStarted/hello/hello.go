package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slices of names
	names := []string{"xzh", "cyh", "ls", "tq", "syy"}
	//Requesting a greeting message
	messages, err := greetings.Hellos(names)
	//If an error was returned, print it to the console and exit the program
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
