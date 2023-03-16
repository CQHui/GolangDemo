package main

import (
	"fmt"
)

func printArray(myArray []int) {
	for index, value := range myArray {
		fmt.Println("index = ", index, "value = ", value)
	}
	myArray[0] = 100
}

func main() {
	//动态数组
	myArray := []int{1, 2, 3, 4}

	fmt.Printf("array2, type = %T\n", myArray)

	printArray(myArray)

	fmt.Println(myArray)

}
