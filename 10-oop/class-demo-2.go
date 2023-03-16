package main

import "fmt"

type Human struct {
	name string
	sex  string
}

type SuperMan struct {
	//继承 HUman
	Human

	level int
}

func (this *Human) Eat() {
	fmt.Println("Human eat...")
}

func (this *Human) Walk() {
	fmt.Println("Human walk...")
}

//------------

// 重定义
func (this *SuperMan) Eat() {
	fmt.Println("Super man eat...")
}

func (this *SuperMan) Fly() {
	fmt.Println("Super man fly...")
}

func main() {
	h := Human{"zhangsan", "male"}

	h.Eat()
	h.Walk()

	s := SuperMan{Human{"lisi", "male"}, 50}
	s.Eat()
	s.Walk()
	s.Fly()
}
