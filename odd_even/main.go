package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(strconv.Itoa(i) + "-偶数")
		} else {
			fmt.Println(strconv.Itoa(i) + "-奇数")
		}
	}
}
