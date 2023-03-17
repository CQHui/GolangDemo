package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"大话西游", 1998, 10, []string{"zhouxinchi"}}

	jsonStr, err := json.Marshal(movie)

	if err != nil {
		fmt.Println("json marshal error")
		return
	}

	fmt.Printf("json = %s\n", jsonStr)

}
