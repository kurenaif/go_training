package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// これ必要！！！忘れない！！
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "query parameter is invalid")
			log.Print(err)
			return
		}

		x := 0.0
		if vals, isExist := r.Form["x"]; isExist {
			if len(vals) != 1 {
				log.Printf("Invalid number of values ​​for query parameter \"x\" (expected 1 but given %d)", len(vals))
				fmt.Fprintf(w, "query parameter is invalid")
				return
			}
			val := vals[0]
			temp, err := strconv.ParseFloat(val, 64)
			if err != nil {
				log.Print(err)
				fmt.Fprint(w, err)
				return
			}
			x = temp
		}

		y := 0.0
		if vals, isExist := r.Form["y"]; isExist {
			if len(vals) != 1 {
				log.Printf("Invalid number of values ​​for query parameter \"y\" (expected 1 but given %d)", len(vals))
				fmt.Fprintf(w, "query parameter is invalid")
				return
			}
			val := vals[0]
			temp, err := strconv.ParseFloat(val, 64)
			if err != nil {
				log.Print(err)
				fmt.Fprint(w, err)
				return
			}
			y = temp
		}

		scale := 0.0
		if vals, isExist := r.Form["scale"]; isExist {
			if len(vals) != 1 {
				log.Printf("Invalid number of values ​​for query parameter \"y\" (expected 1 but given %d)", len(vals))
				fmt.Fprintf(w, "query parameter is invalid")
				return
			}
			val := vals[0]
			temp, err := strconv.ParseFloat(val, 64)
			if err != nil {
				log.Print(err)
				fmt.Fprint(w, err)
				return
			}
			scale = temp
		}

		makeImage(w, x, y, 1.0/scale)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func makeImage(out io.Writer, xOffset float64, yOffset float64, scale float64) {
	xmin, ymin, xmax, ymax := -2.0*scale, -2.0*scale, +2.0*scale, +2.0*scale
	width, height := 1024, 1024

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin + yOffset
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin + xOffset
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(out, img) // NOTE: ignoring errors
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
