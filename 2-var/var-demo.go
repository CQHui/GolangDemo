package main

/*
	四种变量的声明方式
*/

import (
	"fmt"
)

// 声明全局变量 方法一、方法二、方法三是可以的
var var1 int = 100
var var2 = 100

func main() {

	//方法1：
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	//方法2:
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of a = %T\n", b)

	var bb string = "stirng"

	fmt.Printf("bb = %s, type of bb = %T\n", bb, bb)

	//方法3
	var c = 200
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	//方法4
	e := 300
	fmt.Printf("type of e = %T\n", e)

	f := "this is a string"

	fmt.Println("f= ", f)

	//声明多个变量
	var xx, yy int = 100, 200
	fmt.Println(xx, yy)

	var kk, ll = 100, "string again"

	fmt.Println(kk, ll)

}
