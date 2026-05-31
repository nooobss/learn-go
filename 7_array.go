package main

import "fmt"

func doArray() {
	fmt.Println("-- 7_array.go --")
	var aArray [5]int
	fmt.Println("array a:", aArray)

	aArray[4] = 999
	fmt.Println("set:", aArray)
	fmt.Println("get:", aArray[4])
	fmt.Println("len:", len(aArray))

	bArray := [5]int{1,2,3,4,5} // membuat array panjang 5 tipe int isinya 1,2,3,4,5 disimpan sbg b
	fmt.Println("array :", bArray)

	bLebarArray := [...]int{1,3,4,66}
	fmt.Println("array versi tanpa parameter len :", bLebarArray)

	cArray := [...]int{100, 3:400, 500}
	fmt.Println("array dengan inisiasi melompat :", cArray) // yang tidak disebutkan eksplisit akan diberi default 0

	var DuaD [2][3] int
	for i := range 2 {
		for j := range 3 {
			DuaD[i][j] = i + j
		}
	}
	fmt.Println(DuaD) // array 2dimensi
}
