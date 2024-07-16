package main

import (
	"fmt"
	"reflect"
)

// 已知原有类型
func main() {
	var num float64 = 1.2345

	pointer := reflect.ValueOf(&num)
	value := reflect.ValueOf(num)

	fmt.Println(pointer)
	fmt.Println(value)
	// 可以理解为“强制转换”，但是需要注意的时候，转换的时候，如果转换的类型不完全符合，则直接panic
	// Golang 对类型要求非常严格，类型一定要完全符合
	// 如下两个，一个是*float64，一个是float64，如果弄混，则会panic

	convertPointer := pointer.Interface().(*float64)
	convertValue := value.Interface().(float64)
	fmt.Println(convertPointer)
	fmt.Println(convertValue)
}

/** explain: pointer,value都是reflect.ValueOf类型
convertPointer: *float64类型，convertValue: float64类型

question1:解释reflect.Value类型
answer:
reflect.Value 类型是 Go 语言反射包 reflect 中的一个核心类型，它代表了一个 Go 语言变量的 反射值。

反射值 存储了变量的各种信息，包括：

类型: 变量的数据类型，比如 int、string、struct 等等。
值: 变量存储的具体数据，比如 10、"hello" 等等。
地址: 变量在内存中的地址，只有当变量可寻址时（例如变量是可修改的）才存在地址信息。
reflect.Value 类型的用途：

获取变量信息: 使用 reflect.Value 类型可以获取变量的类型、值、地址等信息。
操作变量: 在某些情况下，可以使用 reflect.Value 类型对变量进行修改，例如修改变量的值。
动态处理数据: 可以使用 reflect.Value 类型来实现一些动态处理数据的功能，例如根据变量类型进行不同的操作。

question2:使用 reflect.Value 类型可以获取变量的类型那为什么还要有reflect.Typeof呢
answer:虽然 reflect.Value 类型可以获取变量的类型，但它和 reflect.TypeOf 还是有区别的。

1. reflect.Value 获取的是变量的 运行时 类型。

这意味着它所获取的类型是变量在当前执行时的类型，而这个类型可能与变量的静态类型不完全相同。

例如，如果一个变量被声明为 interface{} 类型，而实际上它存储的是一个 int 类型的数值，那么使用 reflect.Value 获取的类型就是 int。

2. reflect.TypeOf 获取的是变量的 静态 类型。

这意味着它获取的是变量在代码编译时的类型，而这个类型与变量在运行时的实际类型可能不一致。

例如，如果一个变量被声明为 interface{} 类型，那么使用 reflect.TypeOf 获取的类型就是 interface{}，即使它在运行时存储的是 int 类型的数值。

总结：

使用 reflect.Value 获取变量的类型时，需要考虑变量在运行时的实际类型，它可以反映变量的动态变化。
使用 reflect.TypeOf 获取变量的类型时，可以获取到变量的静态类型，它反映了变量的声明信息。
所以，在需要获取变量的运行时类型时，使用 reflect.Value 的 Type() 方法；而在需要获取变量的静态类型时，使用 reflect.TypeOf() 函数。
*/
