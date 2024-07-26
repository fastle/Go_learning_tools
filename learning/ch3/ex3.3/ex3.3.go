package main

import (
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)
var sin30 = math.Sin(angle) // Go的常量是在编译之前就能确定的常量
var cos30 = math.Cos(angle)



func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
        "style='stroke: grey; fill: white; stroke-width: 0.7' "+
        "width='%d' height='%d'>", width, height)
	for i:= 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i + 1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j + 1)
			dx, dy, dz := corner(i + 1, j + 1)
			if math.IsNaN(ax) || math.IsNaN(ay) || math.IsNaN(bx) || math.IsNaN(by) || math.IsNaN(cx) || math.IsNaN(cy) || math.IsNaN(dx) || math.IsNaN(dy) {
				fmt.Fprintf(os.Stderr, "NAN")
			} else {
				//将z映射到一个较大范围

				fmt.Printf("<polygon style='fill: ")
				
				avgz := int((az + bz + cz + dz) * 10.0 + 8.0) * 18
				
				redv, bluev := 0, 0 
				if avgz <= 255 {
					redv = 0
					bluev = 255 - avgz
				} else {
					redv = avgz - 255
					bluev = 0
				}
				if redv > 255 {
					redv = 255
				}
				if bluev > 255{
					bluev = 255
				}
				
				fmt.Printf("#%02X00", redv)
				fmt.Printf("%02X", bluev)	
				fmt.Printf("' points='%g,%g %g,%g %g,%g %g,%g'/>\n",ax, ay, bx, by, cx, cy, dx, dy)
				
			}
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, float64) {
    x := xyrange * (float64(i)/cells - 0.5)
    y := xyrange * (float64(j)/cells - 0.5)

    z := f(x, y)
    sx := width/2 + (x-y)*cos30*xyscale
    sy := height/2 + (x+y)*sin30*xyscale - z*zscale
    return sx, sy, z
}

func f(x, y float64) float64 {
    r := math.Hypot(x, y) 
    return math.Sin(r) / r
}