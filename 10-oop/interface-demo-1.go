package main

import "fmt"

type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is sleeping")
}

func (this *Cat) GetColor() string {
	return this.color
}

func (this *Cat) GetType() string {
	return "Cat"
}

// 具体的类
type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is Sleep")
}

func (this *Dog) GetColor() string {
	return this.color
}

func (this *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal AnimalIF) {
	animal.Sleep() //多态
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("kind = ", animal.GetType())
}

func main() {

	var animal AnimalIF
	animal = &Cat{"Green"}

	animal.Sleep() //调用的就是Cat的Sleep()方法 , 多态的现象

	animal = &Dog{"Yellow"}

	animal.Sleep() // 调用Dog的Sleep方法，多态的现象

	cat := Cat{"Green"}
	dog := Dog{"Yellow"}

	showAnimal(&cat)
	showAnimal(&dog)
}
