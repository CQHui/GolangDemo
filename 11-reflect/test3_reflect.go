package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this *User) Call() {
	fmt.Println("user is called..")
	fmt.Printf("%v\n", this)
}

func DoFiledAndMethod(input interface{}) {
	inputType := reflect.TypeOf(input)
	fmt.Println("inputType is :", inputType.Name())

	inputValue := reflect.ValueOf(input)
	fmt.Println("input value is:", inputValue)

	//获取每个field

	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field, value)
	}

	//获取每个method
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v n", m.Name, m.Type)
	}
}

func main() {
	user := User{1, "AceId", 18}
	DoFiledAndMethod(user)
}
