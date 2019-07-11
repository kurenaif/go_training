package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/big"
	"math/cmplx"
	"os"
)

const (
	xleft                  = 0.443009947545415
	offset                 = 0.00000000000001
	xmin, ymin, xmax, ymax = xleft, 0.405, xleft + offset, +0.405 + offset
	width, height          = 256, 256
)

func main() {

	img128 := image.NewRGBA(image.Rect(0, 0, width, height))
	img64 := image.NewRGBA(image.Rect(0, 0, width, height))
	imgBigFloat := image.NewRGBA(image.Rect(0, 0, width, height))
	imgBigRat := image.NewRGBA(image.Rect(0, 0, width, height))

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			// Image point (px, py) represents complex value z.
			img128.Set(px, py, mandelbrot128(x, y))
			img64.Set(px, py, mandelbrot64(x, y))
			imgBigFloat.Set(px, py, mandelbrotBigFloat(x, y))
			imgBigRat.Set(px, py, mandelbrotBigRat(x, y))
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
	fileBigFloat, err := os.Create("bigFloat.png")
	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
	fileBigRat, err := os.Create("bigRat.png")
	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}

	png.Encode(file128, img128)           // NOTE: ignoring errors
	png.Encode(file64, img64)             // NOTE: ignoring errors
	png.Encode(fileBigFloat, imgBigFloat) // NOTE: ignoring errors
	png.Encode(fileBigRat, imgBigRat)     // NOTE: ignoring errors

	imgDiff := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			_, _, b64, _ := imgBigFloat.At(px, py).RGBA()
			_, _, b128, _ := img128.At(px, py).RGBA()
			var diff uint8
			if b128 < b64 {
				diff = uint8(b64) - uint8(b128)
				// diff = uint8(b64) - uint8(b128)
				if diff != 0 {
					diff = 255
				}
			} else {
				diff = uint8(b128) - uint8(b64)
				if diff != 0 {
					diff = 255
				}
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
	const iterations = 10
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
	const iterations = 10
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

func mandelbrotBigFloat(x float64, y float64) color.Color {
	const iterations = 10
	const contrast = 15

	zx := new(big.Float).SetFloat64(x)
	zy := new(big.Float).SetFloat64(y)
	vx := new(big.Float).SetFloat64(0.0)
	vy := new(big.Float).SetFloat64(0.0)
	bigTwo := new(big.Float).SetFloat64(2.0)
	bigFour := new(big.Float).SetFloat64(4.0)

	for n := uint8(0); n < iterations; n++ {
		nvx := new(big.Float).Sub(new(big.Float).Mul(vx, vx), new(big.Float).Mul(vy, vy))
		nvy := new(big.Float).Mul(new(big.Float).Mul(bigTwo, vx), vy)
		nvx = new(big.Float).Add(nvx, zx)
		nvy = new(big.Float).Add(nvy, zy)
		vx = nvx
		vy = nvy
		// x^2+y^2 > 4
		// sqrtを外した形
		doubleLength := new(big.Float).Add(new(big.Float).Mul(vx, vx), new(big.Float).Mul(vy, vy))
		if doubleLength.Cmp(bigFour) == 1 {
			return color.RGBA{0, 0, 255 - contrast*n, 255}
		}
	}
	return color.Black
}

func mandelbrotBigRat(x float64, y float64) color.Color {
	const iterations = 10
	const contrast = 15

	zx := new(big.Rat).SetFloat64(x)
	zy := new(big.Rat).SetFloat64(y)
	vx := new(big.Rat).SetFloat64(0.0)
	vy := new(big.Rat).SetFloat64(0.0)
	bigTwo := new(big.Rat).SetFloat64(2.0)
	bigFour := new(big.Rat).SetFloat64(4.0)

	for n := uint8(0); n < iterations; n++ {
		nvx := new(big.Rat).Sub(new(big.Rat).Mul(vx, vx), new(big.Rat).Mul(vy, vy))
		nvy := new(big.Rat).Mul(new(big.Rat).Mul(bigTwo, vx), vy)
		nvx = new(big.Rat).Add(nvx, zx)
		nvy = new(big.Rat).Add(nvy, zy)
		vx = nvx
		vy = nvy
		// x^2+y^2 > 4
		// sqrtを外した形
		doubleLength := new(big.Rat).Add(new(big.Rat).Mul(vx, vx), new(big.Rat).Mul(vy, vy))
		if doubleLength.Cmp(bigFour) == 1 {
			return color.RGBA{0, 0, 255 - contrast*n, 255}
		}
	}
	return color.Black
}
