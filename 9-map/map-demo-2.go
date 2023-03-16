package main

import "fmt"

func printMap(map1 map[string]string) {
	for key, value := range map1 {
		fmt.Println("key =", key)
		fmt.Println("value", value)
	}
}

func main() {
	cityMap := make(map[string]string)

	cityMap["China"] = "Beijing"
	cityMap["Japn"] = "Tokyo"
	cityMap["US"] = "NewYork"

	printMap(cityMap)

	delete(cityMap, "Japn")
	cityMap["US"] = "DC"

	fmt.Println("=======")

	printMap(cityMap)

}
