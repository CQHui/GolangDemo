package main

import "fmt"

func func1() {
	fmt.Println("A")
}

func func2() {
	fmt.Println("B")
}

func func3() {
	fmt.Println("C")
}

func func4() {
	fmt.Println("D")
}

func deferFunc() int {
	fmt.Println("defer func called")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called")
	return 0
}

func deferAndReturn() int {
	//先return 后 defer, 有点类似Java finally
	defer deferFunc()
	return returnFunc()
}

func main() {
	//defer打印顺序遵循栈的顺序
	defer func1()
	defer func2()
	defer func3()
	func4()

	deferAndReturn()
}
