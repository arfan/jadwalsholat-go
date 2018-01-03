package main

import "fmt"
import "math"

var (
	n   int
	a   int
	z   float64
	rad float64

	lamd float64
	phi  float64

	t float64
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
		for k := 1; k <= value; k++ {
			n := n0 + k
			a := 6
			z := 110 * rad
			goSub()
		}
	}
}

func goSub() {
	t = n + (a-lamd)/24
	m = (.9856*t - 3.289) * RAD

	l = m + 1.916*rad*math.Sin(m) + .02*rad*math.Sin(2*m) + 282.634*rad
	lh = l / 3.14159 * 12
	ql = int(lh/6) + 1

	if int(ql/2)*2-ql != 0 {
		ql = ql - 1
	}

	ra = math.Atan(.91746*math.Tan(l)) / 3.14159 * 12

	ra = ra + ql*6

	sind = .39782 * math.Sin(l)

	cosd = Math.Sqrt(1 - sind*sind)

	dek = Math.Atan(sind / cosd)

	if a == 15 {
		z = math.Atan(math.Tan(zd) + 1)
	}

	x = (math.Cos(z) - sind*math.Sin(phi)) / (cosd * math.Cos(phi))

	if math.Abs(x) > 1 {
		goto label720
	}

	atnx = math.Atan(math.Sqrt(1-x*x)/x) / rad

	if atnx < 0 {
		atnx = atnx + 180
	}

	h = (360 - atnx) * 24 / 360

	if a == 18 {
		h = 24 - h
	}

	if a == 12 {
		h = 0
	}

label720:

	if a == 15 {
		h = 24 - h
	}

	tloc = h + ra - 0.06571*t - 6.622

	tloc = tloc + 24

	tloc = tloc - int(tloc/24)*24

	st = tloc - lamd + td

}
