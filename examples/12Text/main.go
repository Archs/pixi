package main

import (
	"github.com/Archs/js/dom"
	"github.com/Archs/js/raf"
	"github.com/Archs/pixi"
	"github.com/gopherjs/gopherjs/js"
	"math"
)

var (
	stage    = pixi.NewContainer(0x1099bb)
	renderer = pixi.AutoDetectRenderer(800, 600)
)

func run(t float64) {
	raf.RequestAnimationFrame(run)
	renderer.Render(stage)
}

func main() {
	txt := pixi.NewText("A Text Object will create a line or multiple lines of text. \nTo split a line you can use '\\n' in your text string, or add a wordWrap property set to true and and wordWrapWidth property with a value in the style object.", js.M{
		"font":               "36px Arial bold italic",
		"fill":               "#F7EDCA",
		"stroke":             "#4a1850",
		"strokeThickness":    5,
		"dropShadow":         true,
		"dropShadowColor":    "#000000",
		"dropShadowAngle":    math.Pi / 6,
		"dropShadowDistance": 6,
		"wordWrap":           true,
		"wordWrapWidth":      500,
	})
	txt.Anchor.SetTo(0)
	txt.Position.Set(20, 20)
	stage.AddChild(txt)
	dom.OnDOMContentLoaded(func() {
		el := dom.Wrap(renderer.View)
		dom.Body().AppendChild(el)
		raf.RequestAnimationFrame(run)
	})
}
