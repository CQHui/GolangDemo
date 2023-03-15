package main

import (
	"fmt"
)

// const 定义枚举类型
const (
	//每行iota累加1, 默认值0
	BEIJING = iota
	SHANGHAI
	HANGZHOU
)

func main() {
	//常量, 只读
	const length = 100

	fmt.Println(length)

	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("HANGZHOU = ", HANGZHOU)

}
