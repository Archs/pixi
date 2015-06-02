package pixi

import "github.com/gopherjs/gopherjs/js"

type Renderer struct {
	*js.Object
	View *js.Object `js:"view"`
	// 	width	number	800	optional
	// the width of the canvas view
	Width float64 `js:"width"`
	// height	number	600	optional
	// the height of the canvas view
	Height float64 `js:"height"`
}

func (r Renderer) Render(c *Container) {
	r.Call("render", c.Object)
}

func AutoDetectRenderer(width, height int) Renderer {
	return Renderer{Object: pkg.Call("autoDetectRenderer", width, height)}
}
