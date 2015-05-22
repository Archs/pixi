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
	stage    = pixi.NewContainer()
	render   = pixi.AutoDetectRenderer(300, 300, 0x32232)
	sprite   = pixi.SpriteFromImage("img/bunny.png", false, pixi.ScaleModes.Default)
	lastTime = float64(0)
	stepX    = float64(5)
	stepY    = float64(3)
	counter  = 0
)

func animate(t float64) {
	defer raf.RequestAnimationFrame(animate)
	counter += 1
	if counter%2 != 0 {
		return
	}
	// println(counter, sprite.Position.X, sprite.Position.Y)
	if sprite.Position.X >= 290 || sprite.Position.X <= 10 {
		stepX = -1 * stepX
	}
	if sprite.Position.Y >= 290 || sprite.Position.Y <= 10 {
		stepY = -1 * stepY
	}
	sprite.Position.Set(sprite.Position.X+rand.Float64()*stepX,
		sprite.Position.Y+rand.Float64()*stepY)
	sprite.Rotation = math.Mod(t/100, 2*math.Pi)
	render.Render(stage)
}

func run() {
	sprite.Anchor.SetTo(0.5)
	sprite.Position.SetTo(20)
	stage.AddChild(sprite)
	render.Render(stage)
	el := dom.Wrap(render.View)
	dom.Body().AppendChild(el)
	raf.RequestAnimationFrame(animate)
}

func main() {
	stage.Interactive = true
	stage.HitArea = pixi.NewRectangle(0, 0, 300, 300)
	dom.OnDOMContentLoaded(func() {
		print(stage, stage.Height, stage.Width)
		stage.On(pixi.EventClick, func(ed *pixi.InteractionEvent) {
			println("EventClick", ed.Data.Global)
		}).On(pixi.EventMouseClick, func(ed *pixi.InteractionEvent) {
			println("EventMouseClick", ed.Data.Global)
			// }).On(pixi.EventMouseMove, func(ed *pixi.InteractionEvent) {
			// 	println("EventMouseMove", ed.Data.Global)
		}).On(pixi.EventMouseUp, func(ed *pixi.InteractionEvent) {
			println("EventMouseUp", ed.Data.Global)
		})
		run()
	})
}
