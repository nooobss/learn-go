package main
import (
	"fmt"
	"time"
	"math/rand"
	"math"
)

func main() {
	fmt.Println("-- 1_hello.go --")
	fmt.Println("jam :", time.Now())
	fmt.Println("kode hari ini : ", rand.Intn(10))
	fmt.Printf("nilai pi : %v\n", math.Pi) // perlu placeholder karna Printf()
	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
	variable() //memanggil fungsi variable() yang ada di file 2_var.go
	deklarasi() //memanggil fungsi deklarasi() yang ada di file 2_var.go
	gokil()
	gogo()
	ifelse()
	switching()
	switchinghariIni()
	doArray() //memanggil fungsi doArray() yang ada di file 7_array.go
}

/*placeholder(verb)
integer :
%d, %b, %x

float :
%f, %g, %.2f, 

string :
%s, %q

#note :
verb terbaik : %v otomatis pilih verb terbaik sesuai input-an

*/

/*catatan

untuk menjalankan file : $ go run 1_hello.go
untuk buat file executable(build) : $ go build 1_hello.go
setelah build, muncul 2 file : 1_hello.go dan 1_hello
menjalankan file executable : $ ./1_hello
*/