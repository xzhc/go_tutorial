package main

import "fmt"

func Demo(i int) {
	// defer fmt.Println("1")
	// defer fmt.Println("2")
	// defer fmt.Println("3")
	// defer fmt.Println("4")
	var arr [10]int
	//错误拦截在产生错误前
	defer func() {
		//设置Recover拦截错误信息
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	arr[i] = 10
}

func main() {
	Demo(10)
}
