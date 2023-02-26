package main

import "fmt"

func main() {
	p := struct {
		name   string
		age    int
		height float32
	}{
		name:   "John",
		age:    12,
		height: 172.2,
	}

	fmt.Println(p)
	p.age++
	fmt.Println(p.age)
	fmt.Println(p)
}
