package main

import "fmt"

func ifelse() {
	fmt.Println("-- 5_ifelse.go --")
	if 7%2 == 0 {
		fmt.Println("7 : genap")
	} else {
		fmt.Println("7 : ganjil")
	}

	if 8%2 == 0 {
		fmt.Println("8 habis dibagi 2")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "adalah bil. negatif")
	} else if num < 10 {
		fmt.Println(num, "adalah bil. 1 digit")
	} else {
		fmt.Println(num, "punya multiple digit")
	}
}