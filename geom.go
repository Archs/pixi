package pixi

import "github.com/gopherjs/gopherjs/js"

type Point struct {
	js.Object
	X float32 `js:"x"`
	Y float32 `js:"y"`
}

func (p Point) Set(x, y float32) {
	p.Call("set", x, y)
}

func (p Point) SetTo(v float32) {
	p.Call("set", v)
}
