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
	stage    = pixi.NewContainer(0xf2f2f2)
	renderer *pixi.Renderer
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
	renderer.Render(stage)
	raf.RequestAnimationFrame(run)
}

func main() {
	stage.Interactive = true
	stage.Click(func(dd *pixi.EventData) {
		id := dd.Data
		println(id.Global.X, id.Global.Y)
		b := newBall()
		b.x = id.Global.X
		b.y = id.Global.Y
		stage.AddChild(b)
	})
	g := newBall()
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
