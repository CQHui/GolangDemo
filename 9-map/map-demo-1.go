package main

import "fmt"

func main() {
	var map1 map[string]string
	if map1 == nil {
		fmt.Println("map1 is nil")
	}

	map1 = make(map[string]string, 10)

	map1["one"] = "china"
	map1["two"] = "US"
	map1["three"] = "UK"

	fmt.Println(map1)

	map2 := map[string]string{
		"one":   "java",
		"two":   "python",
		"three": "golang",
	}
	fmt.Println(map2)

}
