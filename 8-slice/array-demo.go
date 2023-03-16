package main

import (
	"fmt"
)

func main() {
	var array1 [10]int

	array2 := [10]int{1, 2, 3, 4}

	array3 := [4]int{1, 2, 3, 4}

	for i := 0; i < len(array1); i++ {
		fmt.Println(array1[i])
	}

	for index, value := range array2 {
		fmt.Println("index = ", index, "value = ", value)
	}

	fmt.Printf("array1, type = %T\n", array1)
	fmt.Printf("array2, type = %T\n", array2)
	fmt.Printf("array3, type = %T\n", array3)

}
