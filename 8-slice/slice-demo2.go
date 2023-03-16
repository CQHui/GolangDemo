package main

import (
	"fmt"
)

func main() {
	slice1 := []int{1, 2, 3}
	fmt.Printf("len = %d, slice1 = %v\n", len(slice1), slice1)

	slice2 := make([]int, 3)

	fmt.Printf("len = %d, slice2 = %v\n", len(slice2), slice2)

	slice3 := make([]int, 3, 5)
	fmt.Printf("len = %d, cap = %d, slice3 = %v\n", len(slice3), cap(slice3), slice3)

	slice3 = append(slice3, 10)
	slice3 = append(slice3, 20)
	fmt.Printf("len = %d, cap = %d, slice3 = %v\n", len(slice3), cap(slice3), slice3)

	fmt.Println("-----2倍扩容-----")

	slice3 = append(slice3, 30)
	fmt.Printf("len = %d, cap = %d, slice3 = %v\n", len(slice3), cap(slice3), slice3)

	slice4 := slice3[0:4]
	slice5 := make([]int, 4)
	copy(slice5, slice3)
	slice3[1] = 100
	fmt.Printf("len = %d, cap = %d, slice3 = %v\n", len(slice3), cap(slice3), slice3)
	fmt.Printf("len = %d, cap = %d, slice4 = %v\n", len(slice4), cap(slice4), slice4)
	fmt.Printf("len = %d, cap = %d, slice5 = %v\n", len(slice5), cap(slice5), slice5)

}
