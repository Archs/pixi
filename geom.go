package pixi

import "github.com/gopherjs/gopherjs/js"

// Point
type Point struct {
	*js.Object
	X float64 `js:"x"`
	Y float64 `js:"y"`
}

func NewPoint(x, y float64) Point {
	return Point{Object: pkg.Get("Point").New(x, y)}
}

// Sets the point to a new x and y position. If y is omitted, both x and y will be set to x.
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

func (r Rectangle) shape() *js.Object {
	return r.Object
}

var EmptyRectangle = Rectangle{Object: pkg.Get("EmptyRectangle")}
