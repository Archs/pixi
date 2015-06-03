package main

import (
	"github.com/Archs/js/dom"
	"github.com/Archs/js/raf"
	"github.com/Archs/pixi"
	"math"
	"math/rand"
)

var (
	stage          = pixi.NewContainer()
	sprites        = pixi.NewContainer()
	render         = pixi.AutoDetectRenderer(800, 600)
	renderTexture  = pixi.NewRenderTexture(render, render.Width, render.Height, pixi.ScaleModes.Default, 1)
	renderTexture2 = pixi.NewRenderTexture(render, render.Width, render.Height, pixi.ScaleModes.Default, 1)
	renderSprit    = pixi.NewSprite(renderTexture.Texture)

	count = 0.0
)

var fruits = []string{
	"assets/spinObj_01.png",
	"assets/spinObj_02.png",
	"assets/spinObj_03.png",
	"assets/spinObj_04.png",
	"assets/spinObj_05.png",
	"assets/spinObj_06.png",
	"assets/spinObj_07.png",
	"assets/spinObj_08.png",
}

type Fruit struct {
	*pixi.Sprite
	url string
}

func (f *Fruit) update(t float64) {
	f.Rotation += 0.1
}

func newFruit(url string) *Fruit {
	f := new(Fruit)
	f.url = url
	f.Sprite = pixi.SpriteFromImage(f.url, false, pixi.ScaleModes.Default)
	f.Anchor.SetTo(0.5)
	f.Position.X = rand.Float64()*400 - 200
	f.Position.Y = rand.Float64()*400 - 200
	f.Update = f.update
	return f
}

func loadSprites() {
	for i := 0; i < 20; i++ {
		f := newFruit(fruits[i%8])
		sprites.AddChild(f)
	}
}

func run(t float64) {
	count += 0.01
	defer raf.RequestAnimationFrame(run)
	for _, s := range sprites.Children {
		s.Update(t)
	}

	temp := renderTexture
	renderTexture = renderTexture2
	renderTexture2 = temp
	// set the new texture
	renderSprit.Texture = renderTexture.Texture
	// twist this up!
	sprites.Rotation -= 0.01
	renderSprit.Scale.SetTo(1 + math.Sin(count)*0.2)
	renderTexture2.Render(stage, false, true)
	// and finally render the stage
	render.Render(stage)
}

func main() {
	// setup
	// ops := pixi.NewRendererOptions()
	// ops.Transparent = true
	// render = pixi.AutoDetectRenderer(800, 600, ops)
	render.BackgroundColor = 0x070202 // transparent works first
	sprites.Position.Set(400, 300)
	renderSprit.Anchor.SetTo(0.5)
	renderSprit.Position.Set(400, 300)
	stage.AddChild(sprites)
	stage.AddChild(renderSprit)
	stage.Interactive = true
	stage.HitArea = pixi.NewRectangle(0, 0, 800, 600)
	// load
	loadSprites()
	dom.OnDOMContentLoaded(func() {
		el := dom.Wrap(render.View)
		dom.Body().AppendChild(el)
		stage.On(pixi.EventClick, func(ed *pixi.InteractionEvent) {
			println("EventClick", ed.Data.Global)
		}).On(pixi.EventMouseClick, func(ed *pixi.InteractionEvent) {
			println("EventMouseClick", ed.Data.Global)
		}).On(pixi.EventMouseUp, func(ed *pixi.InteractionEvent) {
			println("EventMouseUp", ed.Data.Global)
		})
		run(0)
	})
}
