package main

import (
	"fmt"
	"time"
)

func switching() {
	fmt.Println("-- 6_switch.go --")
	i := 2
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("siji")
	case 2:
		fmt.Println("loro")
	case 3:
		fmt.Println("telu")
	}

}

var hariIni = time.Now().Weekday()
func switchinghariIni() {
	fmt.Print(hariIni)
	switch hariIni {
	case time.Saturday, time.Sunday:
		fmt.Println(" : ini weekend")
	default:
		fmt.Println(" : waktunya kerja")
	}
}