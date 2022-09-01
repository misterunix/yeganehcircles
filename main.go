package main

import (
	"fmt"
	"math"
)

func main() {

	for k := 1; k < 14000; k++ {

		x := math.Cos(10.0*math.Pi*float64(k)/14000.0) * (1.0 - 0.5*(math.Cos(16.0*math.Pi*float64(k)/14000.0))*(math.Cos(16.0*math.Pi*float64(k)/14000.0)))
		y := math.Sin(10.0*math.Pi*float64(k)/14000.0) * (1.0 - 0.5*(math.Cos(16.0*math.Pi*float64(k)/14000.0))*(math.Cos(16.0*math.Pi*float64(k)/14000.0)))
		r := 1.0/200.0 + 1.0/10.0*(math.Sin(52.0*math.Pi*float64(k)/14000.0))*(math.Sin(52.0*math.Pi*float64(k)/14000.0))*(math.Sin(52.0*math.Pi*float64(k)/14000.0))*(math.Sin(52.0*math.Pi*float64(k)/14000.0))

		x = (x * 10000.0) + 10000.0
		y = (y * 10000.0) + 10000.0
		r = r * 10000.0

		fmt.Println(x, y, r)

		/*


			term1 := math.Cos(10.0 * math.Pi * float64(k) / 14000.0)
			term2 := 16.0 * math.Pi * float64(k) / 14000.0

			term4 := math.Cos(term2)
			term5 := term4 * term4
			term6 := 1 - 0.5*term5
			X := (term1 * term6) * 10000.0

			term1 = math.Sin(10.0 * math.Pi * float64(k) / 14000.0)
			term2 = 16.0 * math.Pi * float64(k) / 14000.0

			term4 = math.Cos(term2)
			term5 = term4 * term4
			term6 = 1 - 0.5*term5
			Y := (term1 * term6) * 10000.0

			term1 = 1.0/200.0 + 1.0/10.0
			term2 = math.Sin(52.0 * math.Pi * float64(k) / 14000.0)
			term3 := term2 * term2 * term2 * term2
			R := (term1 * term3) * 10000.0

			X = X + 10000.0
			Y = Y + 10000.0
			fmt.Println(X, Y, R)
		*/
	}
}
