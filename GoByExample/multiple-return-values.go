package main

import "fmt"

func vals() (int, int) {
	return 3, 8
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)
	d, c := vals()
	fmt.Println(d, c)
}
