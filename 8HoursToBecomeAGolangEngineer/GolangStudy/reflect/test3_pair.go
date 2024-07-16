package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Book struct{}

func (this *Book) ReadBook() {
	fmt.Println("Read a book")
}

func (this *Book) WriteBook() {
	fmt.Println("Write a book")
}
func main() {
	b := &Book{}

	var r Reader
	r = b
	r.ReadBook()

	var w Writer
	w = r.(Writer)
	w.WriteBook()
}

/** explain:
	1. 接口定义:

type Reader interface { ReadBook() }：定义了一个名为 Reader 的接口，它包含一个方法 ReadBook()。任何实现了 ReadBook() 方法的类型都可以被视为 Reader 接口的实现。
type Writer interface { WriteBook() }：类似地，定义了 Writer 接口，它包含方法 WriteBook()。
2. 结构体实现接口:

type Book struct{}：定义了一个名为 Book 的结构体。
func (this *Book) ReadBook() { ... } 和 func (this *Book) WriteBook() { ... }：这两个方法分别实现了 Reader 和 Writer 接口中定义的方法。这意味着 Book 结构体同时实现了 Reader 和 Writer 接口。
3. 接口的赋值和反射:

b := &Book{}：创建了一个 Book 结构体的指针并赋值给变量 b。
var r Reader：声明一个类型为 Reader 的变量 r。
r = b：将 b 赋值给 r，因为 b 实现了 Reader 接口。此时，r 变量存储的是 Book 结构体的指针，但 r 变量本身的类型是 Reader，它只知道 ReadBook() 方法，并不知道 Book 结构体的其他细节。
4. 类型断言:

var w Writer：声明一个类型为 Writer 的变量 w。
w = r.(Writer)：这是代码的关键部分，它使用了类型断言。r.(Writer) 表达式尝试将 r（类型为 Reader）转换为 Writer类型。因为r的底层类型是Book，并且 Book也实现了Writer` 接口，所以类型断言成功。
w.WriteBook()：现在 w 变量已经成功转换为 Writer 类型，可以调用 WriteBook() 方法了。
总结:

接口反射：r 变量虽然存储的是 Book 结构体的指针，但它的类型是 Reader。程序通过接口反射，仅能访问 Reader 接口定义的方法 ReadBook()。
类型断言：使用 r.(Writer) 将 r 转换为 Writer 类型，这使得程序可以访问 Writer 接口定义的方法 WriteBook()。
*/
