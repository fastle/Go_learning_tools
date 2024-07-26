// 练习3.5 实现彩色效果
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
)

func main(){
	const (
		xmin, ymin, xmax, ymax = -2, -2,  2, 2
		width, height = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0,0,width, height))
	for py := 0; py < height; py++ {
		y := float64(py) / height * (ymax - ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px) / width * (xmax - xmin) + xmin 
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	handler := func(w http.ResponseWriter, r *http.Request) {
		png.Encode(w, img)
	}
	http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func mandelbrot(z complex128) color.Color{
	const iterations = 200
	const contrast = 15
	var v complex128 
	for n := uint8(0); n < iterations; n++ {
		v = v * v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{50, 100, 100 + n, 255 - contrast * n}
		}
	}
	return color.RGBA{200, 200, 100, 0}
}
