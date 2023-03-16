package main

import "fmt"

func myFunc(arg interface{}) {
	fmt.Println("func is called..")
	fmt.Println(arg)

	value, ok := arg.(string)

	if !ok {
		fmt.Println("is not string type")
	} else {
		fmt.Println("is a string type, value = ", value)
	}
}

type Book struct {
	title string
	auth  string
}

func main() {
	book := Book{title: "数学", auth: "人教"}

	myFunc(book)

	myFunc(100)
	myFunc("a string")
}
