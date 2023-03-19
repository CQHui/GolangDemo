package main

import "fmt"

func main() {
	//定义channel
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine is ending")

		fmt.Println("goroutine running")

		c <- 666
	}()

	num := <-c
	fmt.Println("num", num)
	fmt.Println("main goroutine is ending")
}
