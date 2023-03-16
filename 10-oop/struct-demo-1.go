package main

import "fmt"

type myint int

type Book struct {
	title string
	auth  string
}

func main() {
	var book1 Book
	book1.title = "Golang"
	book1.auth = "Elliott"
	fmt.Printf("%v\n", book1)
}
