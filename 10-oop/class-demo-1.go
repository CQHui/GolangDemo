package main

import "fmt"

// 大写表示能够对外访问
type Hero struct {
	Name  string
	Age   int
	Level int
}

func (this *Hero) GetName() string {
	return this.Name
}

func (this *Hero) SetName(name string) {
	this.Name = name
}

func (this *Hero) Show() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Age = ", this.Name)
	fmt.Println("Level = ", this.Name)
}

func main() {
	hero := Hero{Name: "zhangsan", Age: 18, Level: 3}
	hero.Show()

	hero.SetName("lisi")
	fmt.Printf("hero.GetName(): %v\n", hero.GetName())
}
