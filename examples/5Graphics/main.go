package main

import (
	"github.com/Archs/js/dom"
	"github.com/Archs/js/raf"
	"github.com/Archs/pixi"
	"math"
	// "math"
	"math/rand"
)

var (
	stage    = pixi.NewStage(0x872324)
	renderer pixi.Renderer
	bs       = []*Ball{}
)

type Ball struct {
	*pixi.Graphics
	x, y float64
}

func newBall() *Ball {
	b := &Ball{
		Graphics: pixi.NewGraphics(),
		x:        0,
		y:        0,
	}
	b.LineStyle(3, 0xff9387*rand.Float64(), 1)
	bs = append(bs, b)
	return b
}

func (g *Ball) draw(t float64) {
	g.Clear()
	// g.BeginFill(0xFF3300*rand.Float64(), rand.Float64())
	g.LineStyle(3, 0xff9387*rand.Float64(), 1)
	// g.DrawCircle(50+rand.Float64()*20, 50+rand.Float64()*20, 40+rand.Float64()*5)
	g.DrawCircle(g.x, g.y, math.Mod(t, 50))
	// g.EndFill()
}

func run(t float64) {
	for _, b := range bs {
		b.draw(t)
	}
	defer raf.RequestAnimationFrame(run)
	renderer.Render(stage)
}

func main() {
	stage.Click(func(id *pixi.InteractionData) {
		println(id.Global.X, id.Global.Y)
		b := newBall()
		b.x = id.Global.X
		b.y = id.Global.Y
		stage.AddChild(b)
	})
	g := newBall()
	// elps := pixi.NewEllipse(100, 200, 100, 50)
	// g.DrawShape(elps)
	// g.DrawShape(pixi.NewCircle(100, 200, 80))
	// g.DrawShape(pixi.NewRectangle(100, 200, 40, 40))
	// g.DrawShape(pixi.NewPolygon(2, 50, 80, 90, 120, 70))
	// g.DrawShape(pixi.NewPolygon(pixi.NewPoint(90, 90),
	// 	pixi.NewPoint(120, 120),
	// 	pixi.NewPoint(90, 110),
	// 	pixi.NewPoint(90, 90),
	// ))
	stage.AddChild(g)
	dom.OnDOMContentLoaded(func() {
		println("dom loaded")
		renderer = pixi.AutoDetectRenderer(dom.Window().InnerWidth, dom.Window().InnerHeight)
		v := dom.Wrap(renderer.View)
		v.Width = dom.Window().InnerWidth
		v.Height = dom.Window().InnerHeight

		dom.Body().AppendChild(v)
		raf.RequestAnimationFrame(run)
	})
}
