package pixi

import "github.com/gopherjs/gopherjs/js"

type Renderer struct {
	*js.Object
	View *js.Object `js:"view"`
}

func (r Renderer) Render(stage *Stage) {
	r.Call("render", stage.Object)
}

func AutoDetectRenderer(width, height int) Renderer {
	return Renderer{Object: pkg.Call("autoDetectRenderer", width, height)}
}
