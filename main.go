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
	t = N + (A - LAMD) / 24

   M = (.9856 * T - 3.289) * RAD

   L = M + 1.916 * RAD * SIN(M) + .02 * RAD * SIN(2 * M) + 282.634 * RAD

   LH = L / 3.14159 * 12: QL = INT(LH / 6) + 1

   IF INT(QL / 2) * 2 - QL <> 0 THEN QL = QL - 1

	RA = ATN(.91746 * TAN(L)) / 3.14159 * 12

   RA = RA + QL * 6

   SIND = .39782 * SIN(L)

   COSD = SQR(1 - SIND * SIND)

   DEK = ATN(SIND / COSD)

   IF A = 15 THEN Z = ATN(TAN(ZD) + 1)

   X = (COS(Z) - SIND * SIN(PHI)) / (COSD * COS(PHI))

   IF ABS(X) > 1 THEN GOTO 720

   ATNX = ATN(SQR(1 - X * X) / X) / RAD

   IF ATNX < 0 THEN ATNX = ATNX + 180

   H = (360 - ATNX) * 24 / 360

   IF A = 18 THEN H = 24 - H

   IF A = 12 THEN H = 0

   IF A = 15 THEN H = 24 - H

   TLOC = H + RA - .06571 * T - 6.622

   TLOC = TLOC + 24

   TLOC = TLOC - INT(TLOC / 24) * 24

   ST = TLOC - LAMD + TD

}