package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num float64 = 1.2345
	fmt.Println("value:", reflect.ValueOf(num))
	fmt.Println("type:", reflect.TypeOf(num))
}
