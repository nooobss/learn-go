package main

import "fmt"

/*note :
deklarasi var dengan tipe data tidak sah di dalam for loop
deklarasi dengan var dan tipe data hanya bisa diluar for loop
di dalam for loop hanya bisa dengan :=

*/

func gogo() {
	fmt.Println("-- 4_for.go --")

	var i int = 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 1; j < 3; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("range", i)
	}

	for { // for tanpa kondisi akan menjadi infinite loop
		fmt.Println("l o o p")
		break // perlu diberi break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println("ganjil", n)
	}

}