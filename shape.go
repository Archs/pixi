package pixi

import (
	"github.com/gopherjs/gopherjs/js"
)

// Polygon, Circle, Rectangle, Ellipse implements this interface
type Shape interface {
	shape() *js.Object
}

type basicShape struct {
	o *js.Object
}

func wrapBasicShape(o *js.Object) *basicShape {
	return &basicShape{o: o}
}

var (
	// Name    Type    Description
	// SHAPES  object
	// Properties

	// Name    Type    Default Description
	// POLY    object  0
	// RECT    object  1
	// CIRC    object  2
	// ELIP    object  3
	// RREC    object  4
	SHAPES = shapes{Object: pkg.Get("SHAPES")}
)

type shapes struct {
	*js.Object
	POLY int `js:"POLY"` //  0
	RECT int `js:"RECT"` //  1
	CIRC int `js:"CIRC"` //  2
	ELIP int `js:"ELIP"` //  3
	RREC int `js:"RREC"` //  4
}

func (b *basicShape) Type() int {
	return b.o.Get("type").Int()
}

func (s *basicShape) shape() *js.Object {
	return s.o
}

// contains(x, y){boolean}

// core/math/shapes/Circle.js, line 61
// Checks whether the x and y coordinates given are contained within this circle

// Name    Type    Description
// x   number
// The X coordinate of the point to test
// y   number
// The Y coordinate of the point to test
// Returns:

// Type    Description
// boolean Whether the x/y coordinates are within this Circle
func (b *basicShape) Contains(x, y float64) bool {
	return b.o.Call("contains", x, y).Bool()
}

type Circle struct {
	*basicShape
	//     x   number
	// The X coordinate of the center of this circle
	X float64 `js:"x"`
	// y   number
	// The Y coordinate of the center of this circle
	Y float64 `js:"y"`
	// radius  number
	// The radius of the circle
	Radius float64 `js:"radius"`
}

func NewCircle(x, y, radius float64) *Circle {
	return &Circle{
		basicShape: wrapBasicShape(pkg.Get("Circle").New(x, y, radius)),
	}
}

func (c *Circle) Clone() *Circle {
	return &Circle{
		basicShape: wrapBasicShape(c.o.Call("clone")),
	}
}

// getBounds(){Rectangle}
//
// core/math/shapes/Circle.js, line 83
// Returns the framing rectangle of the circle as a Rectangle object
//
//  Returns:
//
//  Type    Description
//  Rectangle   the framing rectangle
func (c *Circle) Bounds() Rectangle {
	o := c.o.Call("getBounds")
	return Rectangle{Object: o}
}

type Ellipse struct {
	Rectangle
}

func NewEllipse(x, y, width, height float64) *Ellipse {
	return &Ellipse{
		Rectangle{Object: pkg.Get("Ellipse").New(x, y, width, height)},
	}
}

func (e *Ellipse) Clone() *Ellipse {
	return &Ellipse{
		Rectangle{Object: e.Object.Call("clone")},
	}
}

type Polygon struct {
	*basicShape
	// points Array.<number>
	// An array of the points of this polygon
	Points []Point `js:"points"`
}

// points  Array.<Point> | Array.<number> | Point | number
//
// This can be an array of Points that form the polygon,
// a flat array of numbers that will be interpreted as [x,y, x,y, ...],
// or the arguments passed can be all the points of the polygon e.g.
//
//  new PIXI.Polygon(new PIXI.Point(), new PIXI.Point(), ...),
//
// or the arguments passed can be flat x,y values e.g.
//
//  new Polygon(x,y, x,y, x,y, ...) where x and y are Numbers.
func NewPolygon(points ...interface{}) *Polygon {
	return &Polygon{
		basicShape: wrapBasicShape(pkg.Get("Polygon").New(points...)),
	}
}

func (p *Polygon) Clone() *Polygon {
	return &Polygon{
		basicShape: wrapBasicShape(p.o.Call("clone")),
	}
}
