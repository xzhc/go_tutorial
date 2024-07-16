package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new goroutine: i = %d\n", i)
		time.Sleep(1 * time.Second) //延时1s
	}
}
func main() {
	//创建一个goroutine,启动另外一个任务
	go newTask()
	// i := 0
	// //main goroutine循环打印
	// for {
	// 	i++
	// 	fmt.Printf("main goroutine: i = %d\n", i)
	// 	time.Sleep(1 * time.Second) //延时1s
	// }
	fmt.Println("main goroutine exit")
}
