package main

import "fmt"

//define a struct
type T struct {
	name string
}

func (t T) method1() {
	t.name = "newName1"
}

func (t *T) method2() {
	t.name = "newName2"
}

func main() {
	t := T{name: "newName"}

	fmt.Println("method1调用前", t.name)
	t.method1()
	fmt.Println("method1调用后", t.name)

	fmt.Println("method2调用前", t.name)
	t.method2()
	fmt.Println("method2调用后", t.name)

}

/**
当调用t.method1()时相当于method1(t)，实参和行参都是类型 T，可以接受。此时在method1()中的t只是参数t的值拷贝，所以method1()的修改影响不到main中的t变量。

当调用t.method2()=>method2(t)，这是将 T 类型传给了 *T 类型，go可能会取 t 的地址传进去：method2(&t)。所以 method1() 的修改可以影响 t。
*/
