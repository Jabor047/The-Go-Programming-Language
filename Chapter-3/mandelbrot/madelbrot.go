
package main

import (
	"math"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main(){
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height = 1024, 1024
		epsX = (xmax - xmin) / width
		epsY = (ymax - ymin) / height
	)

	offX := []float64{-epsX, epsX}
	offY := []float64{-epsY, epsY}

	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py) / height * (ymax - ymin) + ymin
		for px := 0; px < width; px++{
			x := float64(px) / width * (xmax - xmin) + xmin
			// z := complex(x, y) 

			//Supersampling
			SubPixels := make([]color.Color, 0)
			for i := 0; i < 2; i++ {
				for j := 0; j < 2; j++ {
					z := complex(x + offX[i], y + offY[j])
					SubPixels = append(SubPixels, netwon(z))
				}
			}	

			// image point(px, py) represents complex value of z
			img.Set(px, py, avg(SubPixels))
		}
	}
	f, _ := os.Create("mandelbrot.png") // Ignoring errors
	png.Encode(f, img) //NOTE Ignoring errors
}

func avg(colors []color.Color) color.Color{
	var r, b, g, a uint16
	n := len(colors)
	for _, c := range colors{
		r_, g_, b_, a_ := c.RGBA()
		r += uint16(r_ / uint32(n))
		g += uint16(g_ / uint32(n))
		b += uint16(b_ / uint32(n))
		a += uint16(a_ / uint32(n))
	}
	return color.RGBA64{r, g, b, a}
}

func mandelbrot(z complex128) color.Color{
	const iterations = 200
	const contrast = 15

	var v complex128

	for n := uint8(0); n < iterations; n++{
		v = v * v + z
		if cmplx.Abs(v) > 2 {
			// return color.Gray{255 - contrast * n}
			switch {
			case n > 50:
				return color.RGBA{100, 0, 0, 255}
			default:
				//logarithmic blue gradient to show the small differences on the 
				// periphery of the fractal
				logscale := math.Log(float64(n)) / math.Log(float64(iterations))
				return color.RGBA{0, 0 , 255 - uint8(logscale * 255), 255} 
			}
		}
	}

	return color.Black
}

func netwon(z complex128) color.Color{
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++{
		z -= (z - 1/(z * z * z)) / 4
		if cmplx.Abs(z * z * z * z ) < 1e-6 {
			return color.Gray{255 - contrast * i}
		} 
	}
	return color.Black
}