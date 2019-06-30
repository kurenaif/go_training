package main

import (
	"image/color"
	"math/rand"
	"os"
	"time"

	"./lissajous"
)

var palette = []color.Color{color.Black}

const (
	backGroundIndex = 0
	lineIndex       = 1
	cycles          = 5
	res             = 0.001
	size            = 100
	nframes         = 64
	delay           = 8
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	for t := 0; t <= 0xff; t += 2 {
		palette = append(palette, color.RGBA{uint8(0xff - t), 0x00, uint8(t), 0xff})
	}
	lissajous.Lissajous(os.Stdout, cycles)
}

// func lissajous(out io.Writer) {
// 	freq := rand.Float64() * 3.0
// 	anim := gif.GIF{LoopCount: nframes}
// 	phase := 0.0
// 	for i := 0; i < nframes; i++ {
// 		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
// 		img := image.NewPaletted(rect, palette)
// 		for t := 0.0; t < cycles*2.0*math.Pi; t += res {
// 			rate := t / (cycles * 2.0 * math.Pi)
// 			frameRate := float64(i) / nframes * float64(len(palette)-1)
// 			index := uint8((int(float64(len(palette)-1)*rate)+int(frameRate))%len(palette) + 1)
// 			// fmt.Println(index)
// 			x := math.Sin(t)
// 			y := math.Sin(t*freq + phase)
// 			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), index)
// 		}
// 		phase += 0.1
// 		anim.Delay = append(anim.Delay, delay)
// 		anim.Image = append(anim.Image, img)
// 	}
// 	gif.EncodeAll(out, &anim)
// }
