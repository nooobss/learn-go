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

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("sebelum 12 siang")
	default :
		fmt.Println("setelah 12 siang")
	}

	whatAmI := func(i interface{}) { // interface{} adalah tipe data kosong, bisa menampung semua tipe data
		switch t := i.(type) {
		case bool:
			fmt.Println("ini boolean")
		case int:
			fmt.Println("ini integer")
		default:
			fmt.Printf("ini tipe data lain: %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(999)
	whatAmI("jokowi jianncuk")

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