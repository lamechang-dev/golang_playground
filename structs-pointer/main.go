package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	(*p).X = 200 // このように、ポインターを使って構造体のフィールドにアクセスすることもできる
	fmt.Println(v)
}
