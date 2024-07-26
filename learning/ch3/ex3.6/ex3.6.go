// png格式的mandelbrot 图像,
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

			xn := []float64{x - (xmax - xmin) / width / 4, x + (xmax - xmin) / width / 4}
			yn := []float64{y - (ymax - ymin) / height / 4, y + (ymax - ymin) / height / 4}
			var rnow, gnow, bnow, anow uint32 // 因为有相加操作， 所以要大一点
			//fmt.Fprintf(os.Stderr, "%g\n", xn[0])
			for _, xnow := range xn {
				for _, ynow := range yn {
					rtmp, gtmp, btmp, atmp := mandelbrot(complex(xnow, ynow)).RGBA()
					//fmt.Fprintf(os.Stderr, "%d\n", atmp)
					rnow += rtmp >> 8
					gnow += gtmp >> 8
					bnow += btmp >> 8
					anow += atmp >> 8
				}
			}
			rnow /= 4
			gnow /= 4
			bnow /= 4
			anow /= 4
			//fmt.Fprintf(os.Stderr, "%d\n", anow)
			img.SetRGBA(px, py, color.RGBA{uint8(rnow), uint8(gnow), uint8(bnow), uint8(anow)})
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
			return color.Gray{255 - contrast * n}
		}
	}
	return color.Black
}
