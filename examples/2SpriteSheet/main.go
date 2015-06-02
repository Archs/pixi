package main

import (
	"math"
	"math/rand"

	"github.com/Archs/js/raf"
	"github.com/Archs/pixi"

	"github.com/gopherjs/gopherjs/js"
)

var (
	stage    = pixi.NewContainer()
	renderer = pixi.AutoDetectRenderer(800, 600, 0xFFFFFF)
	group    = pixi.NewContainer()
	aliens   = make([]*pixi.Sprite, 0)
	count    = 0.0
)

func onAssetsLoaded() {
	frames := []string{
		"eggHead.png",
		"flowerTop.png",
		"helmlok.png",
		"skully.png",
	}

	for i := 0; i < 100; i++ {
		alien := pixi.SpriteFromFrame(frames[i%4])
		alien.Tint = rand.Uint32()
		alien.Position.X = rand.Float64()*800 - 400
		alien.Position.Y = rand.Float64()*600 - 300
		alien.Anchor.SetTo(0.5)
		aliens = append(aliens, alien)
		group.AddChild(alien)
	}

	raf.RequestAnimationFrame(animate)
}

func animate(t float64) {
	for i := 0; i < len(aliens); i++ {
		aliens[i].Rotation += 0.1
	}

	count += 0.01
	group.Scale.X = math.Sin(count)
	group.Scale.Y = math.Sin(count)
	group.Rotation += 0.01

	renderer.Render(stage)
	raf.RequestAnimationFrame(animate)
}

func main() {
	js.Global.Get("document").Get("body").Call("appendChild", renderer.View)

	pixi.Loader.Add("SpriteSheet", "SpriteSheet.json", onAssetsLoaded)
	pixi.Loader.Load(func(res *js.Object) {
		println(res)
	})

	group.Position.Set(400, 300)
	stage.AddChild(group)
}
