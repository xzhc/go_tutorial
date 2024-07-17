package main

//关闭channel
import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()

	for {
		//ok为true说明channel没有关闭，为false说明管道已经关闭
		if data, ok := <-c; ok {
			fmt.Println(data)
			fmt.Println(ok)
		} else {
			break
		}
	}

	fmt.Println("Finished")
}
