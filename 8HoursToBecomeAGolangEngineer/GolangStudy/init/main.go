package main

import (
	_ "GolangStudy/init/InitLib1"
	_ "GolangStudy/init/InitLib2"
	"fmt"
)

func init() {
	fmt.Println("libmain init")
}

func main() {
	fmt.Println("libmain main")
}
