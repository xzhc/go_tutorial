package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("XZH")
	fmt.Println(message)
}
