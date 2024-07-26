// png格式的mandelbrot 图像
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
			img.Set(px, py, newton(z))
		}
	}
	handler := func(w http.ResponseWriter, r *http.Request) {
		png.Encode(w, img)
	}
	http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func newton(z complex128) color.Color{
	const iterations = 200
	const contrast = 15
	var v complex128 
	v = z
	eps := 1e-8
	ans1, ans2, ans3, ans4 := complex(1, 0), complex(-1, 0), complex(0, 1), complex(0, -1)
	for n := uint8(0); n < iterations; n++ {
		v = v - f(v) / diff(v)
		if cmplx.Abs(v - ans1) < eps || cmplx.Abs(v - ans2) < eps || cmplx.Abs(v - ans3) < eps || cmplx.Abs(v - ans4) < eps {
			return color.Gray{255 - contrast * n}
		}
	}
	return color.Black
}

func f(z complex128) complex128 {
	return z * z * z * z - complex(1,0)
}

func diff(z complex128) complex128 {
	return 4 * z * z * z 
}