// 建立方法样例

package geometry

import "math"

type Point struct {X, Y float64}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

func (p Point) Distance(q Point) float64 { // p 是方法的接收器
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

