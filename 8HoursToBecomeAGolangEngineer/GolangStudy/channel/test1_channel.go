package main

//定义channel变量
import "fmt"

func main() {
	c := make(chan int)
	go func() {
		defer fmt.Println("子go程结束")
		fmt.Println("子go程正在运行")

		c <- 666 //666发送到c
	}()

	num := <-c //从c中接受数据并赋值给num

	fmt.Println("num = ", num)
	fmt.Println("main go程结束")
}
