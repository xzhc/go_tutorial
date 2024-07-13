package greetings

import "fmt"

//Hello returns a greeting for the named person
func Hello(name string) string {
	//Return a greeting tha embeds the name in a message
	message := fmt.Sprintf("Hi, %v. welcome!", name)
	return message
}
