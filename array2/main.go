package main

import "fmt"

func main() {
	start := 1
	end := 4

	ns := []int{1, 2, 3, 4, 5}

	fmt.Println(ns[start:end])

	for i, v := range []string{"foo", "bar", "baz"} {
		fmt.Println(i, v)
		// 0 foo
		// 1 bar
		// 2 baz
	}
}
