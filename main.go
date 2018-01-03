package main

import "fmt"


var (
	n int
	a int
	z float64
)
func main() {
	fmt.Println("jadwal sholat")

	rad := 3.14159 / 180
	city := "Jakarta"

	lamd := -5.7759362
	phi := 106.1174847

	//time difference
	td := 7

	lamd = lamd / 360 * 24
	phi = phi * rad

	bulan := map[string]int{

		"januari":   31,
		"februari":  28,
		"maret":     31,
		"april":     30,
		"mei":       31,
		"juni":      30,
		"juli":      31,
		"agustus":   31,
		"september": 30,
		"oktober":   31,
		"november":  30,
		"desember":  31,
	}

	n0 := 0
	for key, value := range bulan {
		for k:=1;k<=value;k++ {
			n := n0+k
			a := 6
			z := 110 * rad 
			goSub()
		}
	}
}

func goSub() {
	t := n + (a+)
}