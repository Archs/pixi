package main

import (
	"github.com/Archs/js/dom"
	"github.com/Archs/js/raf"
	"github.com/Archs/pixi"
	"math"
)

var (
	stage        = pixi.NewContainer(0x1099bb)
	renderer     = pixi.AutoDetectRenderer(800, 600)
	texture      = pixi.TextureFromImage("img/p2.jpeg", false, pixi.ScaleModes.Default)
	tilingSprite = pixi.NewTilingSprite(texture, renderer.Width, renderer.Height)

	counter = 0.05
)

func run(t float64) {
	counter += 0.05
	raf.RequestAnimationFrame(run)
	tilingSprite.TileScale.Set(math.Sin(counter)+2, math.Cos(counter)+2)
	tilingSprite.TilePosition.X += 1
	tilingSprite.TilePosition.Y += 1
	renderer.Render(stage)
}

func main() {
	stage.AddChild(tilingSprite)
	dom.OnDOMContentLoaded(func() {
		el := dom.Wrap(renderer.View)
		dom.Body().AppendChild(el)
		raf.RequestAnimationFrame(run)
	})
}
