package main

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/Archs/js/raf"
	"github.com/Archs/pixi"

	"github.com/gopherjs/gopherjs/js"
)

var (
	stage      = pixi.NewStage(0xFFFFFF)
	renderer   = pixi.AutoDetectRenderer(800, 600)
	explosions = make([]*pixi.Sprite, 0)
	count      = 0.0
)

func onAssetsLoaded() {
	textures := make([]*pixi.Texture, 0)
	for i := 0; i < 26; i++ {
		frame := fmt.Sprintf("Explosion_Sequence_A %d.png", i+1)
		textures = append(textures, pixi.TextureFromFrame(frame))
	}

	for i := 0; i < 50; i++ {
		explosion := pixi.NewMovieClip(textures)
		explosion.Position.Set(rand.Float64()*800, rand.Float64()*600)
		explosion.Anchor.SetTo(0.5)
		explosion.Rotation = rand.Float64() * math.Pi
		explosion.Scale.SetTo(0.75 + rand.Float64()*0.5)
		explosion.GotoAndPlay(rand.Int() % 27)
		stage.AddChild(explosion)
	}

	raf.RequestAnimationFrame(animate)
}

func animate(t float64) {
	renderer.Render(stage)
	raf.RequestAnimationFrame(animate)
}

func main() {
	js.Global.Get("document").Get("body").Call("appendChild", renderer.View)

	pixi.Loader.Add("name", "SpriteSheet.json", onAssetsLoaded)
	pixi.Loader.Load()
}
