package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	xmin, ymin, xmax, ymax = 0.353, 0.353 + 1e-128, +0.351, +0.351 + 1e-128
	width, height          = 2048, 2048
)

func main() {

	img128 := image.NewRGBA(image.Rect(0, 0, width, height))
	img64 := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			// Image point (px, py) represents complex value z.
			img128.Set(px, py, mandelbrot128(x, y))
			img64.Set(px, py, mandelbrot64(x, y))
		}
	}

	file128, err := os.Create("128.png")
	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
	file64, err := os.Create("64.png")
	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}

	png.Encode(file128, img128) // NOTE: ignoring errors
	png.Encode(file64, img64)   // NOTE: ignoring errors

	imgDiff := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			_, _, b64, _ := img64.At(px, py).RGBA()
			_, _, b128, _ := img128.At(px, py).RGBA()
			var diff uint8
			if b128 < b64 {
				diff = uint8(b64) - uint8(b128)
			} else {
				diff = uint8(b128) - uint8(b64)
			}
			imgDiff.Set(px, py, color.Gray{diff})
		}
	}

	fileDiff, err := os.Create("diff.png")
	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
	png.Encode(fileDiff, imgDiff) // NOTE: ignoring errors
}

func mandelbrot128(x float64, y float64) color.Color {
	z := complex(x, y)
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

func mandelbrot64(x float64, y float64) color.Color {
	z := complex64(complex(x, y))
	const iterations = 200
	const contrast = 15

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(complex128(v)) > 2 {
			return color.RGBA{0, 0, 255 - contrast*n, 255}
		}
	}
	return color.Black
}
