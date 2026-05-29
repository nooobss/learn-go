package main

import (
	"fmt"
	"time"
)
// hanya deklarasi package dan import yang boleh di luar fungsi main() atau fungsi lainnya
// menjalankan fungsi wajib didalam fungsi. Contoh fmt.Println() harus di dalam fungsi
func switching() {
	fmt.Println("-- 6_switch.go --")
	i := 2 // deklarasi := hanya di dalam fungsi, 
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

/*catatan
time adalah package
.Now() adalah fungsi atau pembuat objek yang ada di package time
.Weekday() adalah method yang ada di objek yang dibuat oleh Now()
*/