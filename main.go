package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"time"

	svg "github.com/ajstarks/svgo"
	hsl "github.com/misterunix/colorworks/hsl"
)

type point struct {
	X float64
	Y float64
	R float64
}

var Points []point
var filename string

/*
var img *image.RGBA
var cyan color.RGBA
*/

func main() {

	Points = make([]point, 0)

	/*
		width := 5000
		height := 5000
		upLeft := image.Point{0, 0}
		lowRight := image.Point{width, height}

		img = image.NewRGBA(image.Rectangle{upLeft, lowRight})
		cyan = color.RGBA{100, 200, 200, 0xff}
	*/

	seventhousand()

	//saveImage()

}

/*
func plot(x, y int) {
	img.Set(x, y, cyan)
}
*/

/*
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
*/

func saveImage() {

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	width := 2500
	height := 2500
	wh := width / 2
	hh := height / 2
	canvas := svg.New(f)
	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height, "fill:white")
	//l := len(Points)

	for ii, p := range Points {
		x := int(p.X*500.0) + wh
		y := int(p.Y*500.0) + hh
		r := int(p.R * 500)
		h := float64(ii) / 7000.0 * 360.0
		sa := 80.0 //(p.Y * 100.0) + 100
		l := 80.0  // (p.R * 100.0) + 100
		red, green, blue := hsl.HSLtoRGB(h, sa, l)
		rgb := fmt.Sprintf("#%02x%02x%02x", red, green, blue)
		s := "fill:none;stroke:" + rgb + ";stroke-width:1;stroke-opacity:0.3"
		canvas.Circle(x, y, r, s)
	}

	canvas.End()

}

/*
func circle(x, y, r int) {
	for i := 0; i < 360; i++ {
		plot(x+int(math.Cos(float64(i))*float64(r)), y+int(math.Sin(float64(i))*float64(r)))
	}
}
*/

func seventhousand() {

	rseed := time.Now().UnixNano()

	rand.Seed(rseed)

	for tmp := 1; tmp <= 40; tmp++ {
		tmpf := float64(tmp)
		filename = fmt.Sprintf("SVG/7000-%02d-5.svg", tmp)
		filename1 := fmt.Sprintf("SVG/7000-%02d-5-desc.tmp", tmp)
		f, err := os.OpenFile(filename1, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		q := fmt.Sprintf("x := math.Sin(math.Pi*6*float64(k)/7000.0) * (1 + math.Pow(math.Cos(math.Pi*20*float64(k)/7000.0), 2) - 1.0/%f*math.Pow(math.Cos(16.0*math.Pi*float64(k)/7000.0), 2))\n", tmpf)
		f.WriteString(q)
		q = fmt.Sprintf("y := math.Cos(math.Pi*6*float64(k)/7000.0) * (1 + math.Pow(math.Cos(math.Pi*20*float64(k)/7000.0), 2) - 1.0/%f*math.Pow(math.Cos(16.0*math.Pi*float64(k)/7000.0), 2))\n", tmpf)
		f.WriteString(q)
		q = fmt.Sprintf("r := 1.0/200.0 + 1.0/7.0*math.Pow(math.Cos(28*math.Pi*float64(k)/7000.0), 2) + 1.0/7.0*math.Pow(math.Sin(24.0*math.Pi*float64(k)/7000.0), 6) + 1.0/20.0*math.Pow(math.Cos(84.0*math.Pi*float64(k)/7000.0), 4)\n")
		f.WriteString(q)

		for k := 1; k <= 7000; k++ {

			//x := math.Sin(math.Pi*6*float64(k)/7000.0) * (1 + math.Pow(math.Cos(math.Pi*20*float64(k)/7000.0), 2) - 1.0/5.0*math.Pow(math.Cos(16.0*math.Pi*float64(k)/7000.0), 2))
			//y := math.Cos(math.Pi*6*float64(k)/7000.0) * (1 + math.Pow(math.Cos(math.Pi*20*float64(k)/7000.0), 2) - 1.0/5.0*math.Pow(math.Cos(16.0*math.Pi*float64(k)/7000.0), 2))
			x := math.Sin(math.Pi*6*float64(k)/7000.0) * (1 + math.Pow(math.Cos(math.Pi*20*float64(k)/7000.0), 2) - 1.0/tmpf*math.Pow(math.Cos(16.0*math.Pi*float64(k)/7000.0), 2))
			y := math.Cos(math.Pi*6*float64(k)/7000.0) * (1 + math.Pow(math.Cos(math.Pi*20*float64(k)/7000.0), 2) - 1.0/tmpf*math.Pow(math.Cos(16.0*math.Pi*float64(k)/7000.0), 2))
			r := 1.0/200.0 + 1.0/7.0*math.Pow(math.Cos(28*math.Pi*float64(k)/7000.0), 2) + 1.0/7.0*math.Pow(math.Sin(24.0*math.Pi*float64(k)/7000.0), 6) + 1.0/20.0*math.Pow(math.Cos(84.0*math.Pi*float64(k)/7000.0), 4)

			p := point{X: x, Y: y, R: r}
			Points = append(Points, p)
		}

		saveImage()
		Points = Points[:0]

	}

}
