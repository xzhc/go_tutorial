package main

import "fmt"

//单向channel及其应用
//   chan<- //只写
func counter(out chan<- int) {
	defer close(out)
	for i := 0; i < 5; i++ {
		out <- i //如果对方不读 会阻塞
	}
}

//   <-chan //只读
func printer(in <-chan int) {
	for num := range in {
		fmt.Println(num)
	}
}

func main() {
	c := make(chan int) //   chan   //读写

	go counter(c) //生产者
	printer(c)    //消费者

	fmt.Println("done")
}
