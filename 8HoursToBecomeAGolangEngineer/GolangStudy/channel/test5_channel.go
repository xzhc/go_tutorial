package main

import "fmt"

//使用 range 来迭代不断操作channel
func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	for data := range c {
		fmt.Println(data)
	}
	fmt.Println("Finished!")
}
