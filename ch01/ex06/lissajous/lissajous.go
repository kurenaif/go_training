package lissajous

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var palette = []color.Color{color.Black}

const (
	backGroundIndex = 0
	lineIndex       = 1
	res             = 0.001
	size            = 100
	nframes         = 64
	delay           = 8
)

func init() {
	for t := 0; t <= 0xff; t += 2 {
		palette = append(palette, color.RGBA{uint8(0xff - t), 0x00, uint8(t), 0xff})
	}
}

func Lissajous(out io.Writer, cycles int) {
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		endTime := float64(cycles) * 2.0 * math.Pi
		for t := 0.0; t < endTime; t += res {
			// 全体のうち、どれくらいの時間をプロットしたか
			timeRate := t / endTime
			// 全体のうち、どれくらいのフレームをプロットしたか
			// frameRate := float64(i) / nframes
			// 切り捨てすると、最後の要素が使われなくなるので+1する
			index := uint8(timeRate * float64(len(palette)+1))
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), index)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
