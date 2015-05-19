package main

import (
	"github.com/Archs/js/raf"
	"github.com/Archs/pixi"

	"github.com/gopherjs/gopherjs/js"
)

var (
	stage    = pixi.NewStage(0x66FF99)
	renderer = pixi.AutoDetectRenderer(400, 300)
	texture  = pixi.TextureFromImage("bunny.png", true, pixi.ScaleModes.Default)
	bunny    = pixi.NewSprite(texture)
)

func animate(t float32) {
	raf.RequestAnimationFrame(animate)
	bunny.Rotation += 0.1
	renderer.Render(stage)
}

func main() {
	js.Global.Get("document").Get("body").Call("appendChild", renderer.View)

	bunny.Anchor.SetTo(0.5)
	bunny.Position.Set(200, 150)

	stage.AddChild(bunny)

	raf.RequestAnimationFrame(animate)
}
