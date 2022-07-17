package main

import (
	"fmt"
	"math"
)

type point struct {
	x float64
	y float64
}

// newPoint - конструктор возвращает точку с заданными координатами
func newPoint(x, y float64) point {
	return point{x, y}
}

//
func distance(p1, p2 point) float64 {
	return math.Sqrt(math.Abs((math.Pow(p2.x, 2) - math.Pow(p1.x, 2)) + (math.Pow(p2.y, 2) - math.Pow(p1.y, 2))))
}

func main() {
	p1 := newPoint(50.5, 29.3)
	p2 := newPoint(20.1, 10.5)
	d := distance(p2, p1)
	fmt.Println(d)
}
