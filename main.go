package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

var img *image.RGBA
var cyan color.RGBA

func main() {

	width := 10000
	height := 10000
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img = image.NewRGBA(image.Rectangle{upLeft, lowRight})
	cyan = color.RGBA{100, 200, 200, 0xff}

	for k := 1; k < 14000; k++ {

		x := math.Cos(10.0*math.Pi*float64(k)/14000.0) * (1.0 - 0.5*(math.Cos(16.0*math.Pi*float64(k)/14000.0))*(math.Cos(16.0*math.Pi*float64(k)/14000.0)))
		y := math.Sin(10.0*math.Pi*float64(k)/14000.0) * (1.0 - 0.5*(math.Cos(16.0*math.Pi*float64(k)/14000.0))*(math.Cos(16.0*math.Pi*float64(k)/14000.0)))
		r := 1.0/200.0 + 1.0/10.0*(math.Sin(52.0*math.Pi*float64(k)/14000.0))*(math.Sin(52.0*math.Pi*float64(k)/14000.0))*(math.Sin(52.0*math.Pi*float64(k)/14000.0))*(math.Sin(52.0*math.Pi*float64(k)/14000.0))

		x = (x * 10000.0) + 10000.0
		y = (y * 10000.0) + 10000.0
		r = r * 10000.0

		//fmt.Println(x, y, r)
		plotLine(int(x), int(y), int(x+r), int(y+r))

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

	saveImage()

}

func plot(x, y int) {
	img.Set(x, y, cyan)
}

func plotLine(x0, y0, x1, y1 int) {
	dx := math.Abs(float64(x1) - float64(x0))
	var sx, sy float64
	if x0 < x1 {
		sx = 1
	} else {
		sx = -1
	}
	//sx := x0 < x1 ? 1 : -1
	dy := -math.Abs(float64(y1) - float64(y0))
	if y0 < y1 {
		sy = 1
	} else {
		sy = -1
	}
	//sy = y0 < y1 ? 1 : -1
	errorterm := dx + dy

	for {
		plot(x0, y0)
		if x0 == x1 && y0 == y1 {
			break
		}
		e2 := 2 * errorterm
		if e2 >= dy {
			if x0 == x1 {
				break
			}
			errorterm = errorterm + dy
			x0 = x0 + int(sx)
		}
		if e2 <= dx {
			if y0 == y1 {
				break
			}
			errorterm = errorterm + dx
			y0 = y0 + int(sy)
		}
	}
}

func saveImage() {
	file, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, img)
}
