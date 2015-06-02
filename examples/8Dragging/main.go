package main

import (
	"github.com/Archs/js/raf"
	"github.com/Archs/pixi"

	"github.com/gopherjs/gopherjs/js"
)

var (
	stage    = pixi.NewContainer()
	renderer = pixi.AutoDetectRenderer(800, 600, 0x97C56E)
)

func animate(t float64) {
	raf.RequestAnimationFrame(animate)
	renderer.Render(stage)
}

type DraggableSprite struct {
	*pixi.Sprite
	dragging bool
	data     *pixi.InteractionData
}

func NewDraggableSprite(url string, x, y float64) *DraggableSprite {
	sprite := pixi.SpriteFromImage(url, false, pixi.ScaleModes.Nearest)

	sprite.Interactive = true
	sprite.ButtonMode = true

	sprite.Position.Set(x, y)
	sprite.Anchor.SetTo(0.5)
	sprite.Scale.SetTo(3)

	ds := &DraggableSprite{Sprite: sprite}

	ds.MouseDown(func(ed *pixi.InteractionEvent) { ds.down(ed.Data) })
	ds.TouchStart(func(ed *pixi.InteractionEvent) { ds.down(ed.Data) })

	ds.MouseUp(func(ed *pixi.InteractionEvent) { ds.up(ed.Data) })
	ds.MouseUpOutside(func(ed *pixi.InteractionEvent) { ds.up(ed.Data) })
	ds.TouchEnd(func(ed *pixi.InteractionEvent) { ds.up(ed.Data) })
	ds.TouchEndOutside(func(ed *pixi.InteractionEvent) { ds.up(ed.Data) })

	ds.MouseMove(func(ed *pixi.InteractionEvent) { ds.move(ed.Data) })

	return ds
}

func (ds *DraggableSprite) down(data *pixi.InteractionData) {
	ds.Alpha = 0.75
	ds.dragging = true
	ds.data = data
}

func (ds *DraggableSprite) up(data *pixi.InteractionData) {
	ds.Alpha = 1
	ds.dragging = false
	ds.data = nil
}

func (ds *DraggableSprite) move(data *pixi.InteractionData) {
	if ds.dragging {
		position := ds.data.LocalPosition(ds.Parent())
		ds.Position.Set(position.X, position.Y)
	}
}

func main() {
	js.Global.Get("document").Get("body").Call("appendChild", renderer.View)

	stage.AddChild(NewDraggableSprite("bunny.png", 400, 300))

	raf.RequestAnimationFrame(animate)
}
