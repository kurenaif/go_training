package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
)

var (
	width, height = 600, 320                       // canvas size in pixels
	cells         = 100                            // number of grid cells
	xyrange       = 20.0                           // axis ranges (-xyrange..+xyrange)
	xyscale       = float64(width) / 2.0 / xyrange // pixels per x or y unit
	zscale        = float64(height) * 0.1          // pixels per z unit
	angle         = math.Pi / 6                    // angle of x, y axes (=30°)
)

func average(vs []float64) float64 {
	sum := 0.0
	for _, v := range vs {
		sum += v
	}
	return sum / float64(len(vs))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var color = ""

		// これ必要！！！忘れない！！
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "query parameter is invalid")
			log.Print(err)
			return
		}

		w.Header().Set("Content-Type", "image/svg+xml")

		if vals, isExist := r.Form["width"]; isExist {
			if len(vals) != 1 {
				log.Printf("Invalid number of values ​​for query parameter \"width\" (expected 1 but given %d)", len(vals))
				fmt.Fprintf(w, "query parameter is invalid")
				return
			}
			val := vals[0]
			wid, err := strconv.Atoi(val)
			if err != nil {
				log.Print(err)
				fmt.Fprint(w, err)
				return
			}
			width = wid
			print(width)
		}

		if vals, isExist := r.Form["height"]; isExist {
			if len(vals) != 1 {
				log.Printf("Invalid number of values ​​for query parameter \"height\" (expected 1 but given %d)", len(vals))
				fmt.Fprintf(w, "query parameter is invalid")
				return
			}
			val := vals[0]
			hei, err := strconv.Atoi(val)
			if err != nil {
				log.Print(err)
				fmt.Fprint(w, err)
				return
			}
			height = hei
		}

		if vals, isExist := r.Form["color"]; isExist {
			if len(vals) != 1 {
				log.Printf("Invalid number of values ​​for query parameter \"color\" (expected 1 but given %d)", len(vals))
				fmt.Fprintf(w, "query parameter is invalid")
				return
			}
			color = vals[0]
		}
		plot3d(w, color)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func plot3d(out io.Writer, color string) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
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
		if color == "" {
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, getColor((z-zmin)/(zmax-zmin)))
		} else {
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}

	fmt.Fprintln(out, "</svg>")

}

// 変換後のx, 変換後のy, 変換前のz, エラーを返す
// zは二次元投影する前の座標なので注意！(色付け用)
func corner(i, j int) (float64, float64, float64, error) {
	sinAngle, cosAngle := math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)

	// Compute surface height z.
	z := f(x, y)
	if math.IsNaN(z) {
		return 0, 0.0, 0.0, errors.New("Divide by zero")
		os.Exit(1)
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width)/2 + float64(x-y)*cosAngle*xyscale
	sy := float64(height)/2 + float64(x+y)*sinAngle*xyscale - z*zscale
	return sx, sy, z, nil
}

func f(x, y float64) float64 {
	return math.Sin(x) + math.Cos(y)
}

func getColor(s float64) string {
	r := int64(255.99 * s)
	b := 255 - int64(255.99*s)
	return fmt.Sprintf("rgb(%d,0,%d)", r, b)
}
