package main

import (
	"fmt"
)

func foo1(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100

	return c
}

// 返回多个返回值, 匿名的
func foo2(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("a = ", a)

	return 100, 200
}

// 返回多个返回值, 有形参名称
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("a = ", a)
	fmt.Println("a = ", a)

	//默认初始值0
	return 1000, r2
}

func main() {
	c := foo1("abc", 1)
	fmt.Println("c = ", c)

	r1, r2 := foo2("edf", 2)
	fmt.Println("r1 = ", r1, "r2 = ", r2)

	res1, res2 := foo3("edffafe", 3)
	fmt.Println("res1 = ", res1, "res2 = ", res2)

}
