package main

import "fmt"

func main() {
	ns := []int{1, 2, 3, 4, 5}
	fmt.Println(ns)
	slicedNs := ns[:3]
	fmt.Println(slicedNs)
	fmt.Println(len(slicedNs))
	fmt.Println(cap(slicedNs))

	println(len(ns), cap(ns))
	// 容量が足りない場合
	// 元のおよそ2倍の容量の配列を確保しなおす
	ns = append(ns, 6, 7)
	fmt.Println(ns)
	println(len(ns), cap(ns))

	println("==============")

	// range の例
	for _, v := range ns {
		fmt.Println(v)
	}
}
