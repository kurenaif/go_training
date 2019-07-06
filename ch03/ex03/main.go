package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 20.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.1        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func average(vs []float64) float64 {
	sum := 0.0
	for _, v := range vs {
		sum += v
	}
	return sum / float64(len(vs))
}

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	points := [][]float64{}
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az, err := corner(i+1, j)
			if err != nil {
				continue
			}
			bx, by, bz, err := corner(i, j)
			if err != nil {
				continue
			}
			cx, cy, cz, err := corner(i, j+1)
			if err != nil {
				continue
			}
			dx, dy, dz, err := corner(i+1, j+1)
			if err != nil {
				continue
			}
			z := average([]float64{az, bz, cz, dz})
			points = append(points, []float64{ax, ay, bx, by, cx, cy, dx, dy, z})
		}
	}

	// zの最大値、最小値を求める
	zmin := math.MaxFloat64
	zmax := -math.MaxFloat64
	for _, point := range points {
		z := point[8]
		zmin = math.Min(zmin, z)
		zmax = math.Max(zmax, z)
	}

	// 描画
	for _, point := range points {
		ax := point[0]
		ay := point[1]
		bx := point[2]
		by := point[3]
		cx := point[4]
		cy := point[5]
		dx := point[6]
		dy := point[7]
		z := point[8]
		println((z - zmin) / (zmax - zmin))
		fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill=%s/>\n",
			ax, ay, bx, by, cx, cy, dx, dy, getColor((z-zmin)/(zmax-zmin)))
	}

	fmt.Println("</svg>")
}

// 変換後のx, 変換後のy, 変換前のz, エラーを返す
// zは二次元投影する前の座標なので注意！(色付け用)
func corner(i, j int) (float64, float64, float64, error) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)
	if math.IsNaN(z) {
		return 0, 0.0, 0.0, errors.New("Divide by zero")
		os.Exit(1)
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z, nil
}

func f(x, y float64) float64 {
	return math.Sin(x) + math.Cos(y)
}

func getColor(s float64) string {
	r := int64(255.99 * s)
	b := 255 - int64(255.99*s)
	return fmt.Sprintf("'rgb(%d,0,%d)'", r, b)
}
