package main

import "fmt"

func doArray() {
	fmt.Println("-- 7_array.go --")
	var a [5]int
	fmt.Println("array a:", a)

	a[4] = 999
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a))
}