package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2.0, -2.0, +2.0, +2.0
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, newton(z))
		}
	}

	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func averageColors(colors []color.Color) color.Color {
	var rsum, gsum, bsum uint32

	for _, color := range colors {
		r, g, b, _ := color.RGBA()
		rsum += r
		gsum += g
		bsum += b
	}

	r := rsum / uint32(len(colors))
	g := gsum / uint32(len(colors))
	b := bsum / uint32(len(colors))
	return color.NRGBA{uint8(r), uint8(g), uint8(b), 255}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{0, 0, 255 - contrast*n, 255}
		}
	}
	return color.Black
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			c := acos(z)
			r, g, b, a := c.RGBA()
			return color.RGBA{uint8(r) / (i + 1), uint8(g) / (i + 1), uint8(b) / (i + 1), uint8(a)}
		}
	}
	return color.Black
}
