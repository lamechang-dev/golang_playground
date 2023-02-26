package main

import "fmt"

func main() {

	const break_target = 9

	for i := 0; i <= 10; i++ {

		if i == break_target {
			break
		}

		fmt.Println(i)
	}
}
