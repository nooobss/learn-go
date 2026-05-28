package main
import (
	"fmt"
	"math"
)

const box string = "kotak"
func gokil() {
	fmt.Println("-- 3_const.go --")
	fmt.Println(box)
	const n = 3000

	const d = 3e3 / n
	fmt.Println("nilai d : ", d)

	fmt.Println("nilai d int64 : ", int64(d))

	fmt.Println(math.Sin(d))
}