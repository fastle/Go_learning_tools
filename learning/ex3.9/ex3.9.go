// 实现http传参， 先处理参数再绘制

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main(){
	const (
	//	xmin, ymin, xmax, ymax = -2, -2,  2, 2
		width, height = 1024, 1024
	)
	params := map[string] float64 { // 使用map直接存， 要是在之后使用多个if判断， 耗时反而会更久
		"xmin": -2, 
		"xmax": 2,
		"ymin": -2,
		"ymax": 2,
		"zoom": 1,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
			for name := range params {
				s := r.FormValue(name)
				if s == "" {
					continue
				}
				f, err := strconv.ParseFloat(s, 64)
				if err != nil {
					http.Error(w, fmt.Sprintf("query param %s: %s", name, err), http.StatusBadRequest)
					return
				}
				params[name] = f  // 读取信息的方式， 来自Gopl-homework
			}
			if params["xmax"] <= params["xmin"] || params["ymax"] <= params["ymin"] {
				http.Error(w, fmt.Sprintf("min coordinate greater than max"), http.StatusBadRequest)
					return 
			}
			xmin, xmax, ymin, ymax, zoom := params["xmin"],params["xmax"],params["ymin"],params["ymax"],params["zoom"]
			lenX, lenY := xmax - xmin, ymax - ymin
			midX, midY := xmin + lenX / 2, ymin + lenY / 2
			xmin, xmax, ymin, ymax = midX - lenX / 2 / zoom, midX + lenX / zoom / 2, midY - lenY / 2 / zoom, midY + lenY / 2 / zoom
			//fmt.Fprintf(os.Stderr, "%g %g %g %g\n", xmin, xmax, ymin, ymax)
			img := image.NewRGBA(image.Rect(0,0,width, height))
			for py := 0; py < height; py++ {
				y := float64(py) / height * (ymax - ymin) + ymin
				for px := 0; px < width; px++ {
					x := float64(px) / width * (xmax - xmin) + xmin 
					z := complex(x, y)
					img.Set(px, py, mandelbrot(z))
				}
			}
			png.Encode(w, img)
	})
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