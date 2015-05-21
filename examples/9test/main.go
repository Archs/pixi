package main

import (
	"github.com/Archs/js/dom"
	"github.com/Archs/js/raf"
	"github.com/Archs/pixi"
	_ "github.com/Archs/pixi/shim"
	"math"
	"math/rand"
)

var (
	stage    = pixi.NewStage(0x32232)
	render   = pixi.AutoDetectRenderer(300, 300)
	sprite   = pixi.SpriteFromImage("img/bunny.png", false, pixi.ScaleModes.Default)
	lastTime = float64(0)
	stepX    = float64(2)
	stepY    = float64(2)
)

func animate(t float64) {
	if sprite.Position.X > 300 || sprite.Position.X < 0 {
		stepX = -1 * stepX
	}
	if sprite.Position.Y > 300 || sprite.Position.Y < 0 {
		stepY = -1 * stepY
	}
	sprite.Position.Set(sprite.Position.X+rand.Float64()+stepX,
		sprite.Position.Y+rand.Float64()+stepY)
	sprite.Rotation = math.Mod(t/100, 2*math.Pi)
	render.Render(stage)
	raf.RequestAnimationFrame(animate)
}

func run() {
	sprite.Anchor.SetTo(0.5)
	sprite.Position.SetTo(0)
	stage.AddChild(sprite)
	render.Render(stage)
	el := dom.Wrap(render.View)
	dom.Body().AppendChild(el)
	raf.RequestAnimationFrame(animate)
}

func main() {
	dom.OnDOMContentLoaded(run)
}
