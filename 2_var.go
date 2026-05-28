package main // file 2_var.go disebut sebagai helper
import "fmt"

func variable() { //fungsi digunakan di file 1_hello.go, pintu masuk via func main()
	fmt.Println("-- 2_var.go --")
	var nama = "kanjeng dimas"
	fmt.Println(nama)
}

func deklarasi() {
	var a = "initialise" // deklarasi dengan keyword var
	fmt.Println(a)

	var b, c int = 1,2
	fmt.Println(b,c)

	var d = true
	fmt.Println(d)

	f := "salak" // shorthand hanya di dalam fungsi, tidak bisa di luar fungsi
	fmt.Println(f)
}