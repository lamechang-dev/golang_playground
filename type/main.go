package main

import (
	"fmt"
	"io"
	"time"
)

func main() {
	// 組み込み型を基にする
	type MyInt int
	// 他のパッケージの型を基にする
	type MyWriter io.Writer
	// 型リテラルを基にする
	type Person struct {
		Name string
	}

	var n int = 100
	m := MyInt(n)
	n = int(m)

	d := 10 * time.Second
	fmt.Println(d)
	a := int(d)

	fmt.Println(a)
}
