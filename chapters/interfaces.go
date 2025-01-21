package chapters

import (
	"fmt"
	"math"
)

type Geometry interface {
	Area() float64
	Perim() float64
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	R float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perim() float64 {
	return 2*r.Width + 2*r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.R * c.R
}

func (c Circle) Perim() float64 {
	return 2 * math.Pi * c.R
}

func Measure(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.Area())
	fmt.Println(g.Perim())
}

func DetectCicle(g Geometry) {
	if c, ok := g.(Circle); ok {
		fmt.Println("circle with radius:", c.R)
	}
}
