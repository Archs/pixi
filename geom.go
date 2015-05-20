package pixi

import "github.com/gopherjs/gopherjs/js"

// Point
type Point struct {
	*js.Object
	X float64 `js:"x"`
	Y float64 `js:"y"`
}

func NewPoint(x, y, width, height float64) Point {
	return Point{Object: pkg.Get("Point").New(x, y)}
}

func (p Point) Set(x, y float64) {
	p.Call("set", x, y)
}

func (p Point) SetTo(v float64) {
	p.Call("set", v)
}

func (p Point) Clone() Point {
	return Point{Object: p.Call("clone")}
}

// Rectangle
type Rectangle struct {
	*js.Object
	X      float64 `js:"x"`
	Y      float64 `js:"y"`
	Width  float64 `js:"width"`
	Height float64 `js:"height"`
}

func NewRectangle(x, y, width, height float64) Rectangle {
	return Rectangle{Object: pkg.Get("Rectangle").New(x, y, width, height)}
}

func (r Rectangle) Clone() Rectangle {
	return Rectangle{Object: r.Call("clone")}
}

func (r Rectangle) Contains(x, y float64) bool {
	return r.Call("contains", x, y).Bool()
}

var EmptyRectangle = Rectangle{Object: pkg.Get("EmptyRectangle")}
