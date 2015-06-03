package main

import (
	"github.com/Archs/js/dom"
	"github.com/Archs/js/raf"
	"gopkg.in/Archs/pixi.v3"
)

var (
	stage    = pixi.NewContainer()
	renderer = pixi.AutoDetectRenderer(800, 600)
)

func animate(t float64) {
	raf.RequestAnimationFrame(animate)
	renderer.Render(stage)
}

func main() {
	dom.OnDOMContentLoaded(func() {
		el := dom.Wrap(renderer.View)
		dom.Body().AppendChild(el)
		raf.RequestAnimationFrame(animate)
	})
}
