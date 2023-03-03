package main

import "fmt"

func main() {
	array := [3]int{1, 2, 3}

	for _, value := range array {
		fmt.Println(value)
	}
}
