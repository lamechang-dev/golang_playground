package main

import "fmt"

func main() {
	var ns = [...]int{1, 2, 3, 4, 5}
	fmt.Println(ns)
	sliced_ns := ns[:3]
	fmt.Println(sliced_ns)
}
