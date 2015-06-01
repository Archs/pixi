package main

import (
	"container/list"
	"github.com/Archs/js/dom"
	"github.com/Archs/js/raf"
	"github.com/Archs/pixi"
	// "math"
	"math/rand"
)

var (
	stage    = pixi.NewStage(0x040404)
	renderer pixi.Renderer
	ps       = list.New()
	COLOURS  = []float64{0x69D2E7, 0xA7DBD8, 0xE0E4CC, 0xF38630, 0xFA6900, 0xFF4E50, 0xF9D423}
)

type Particle struct {
	*pixi.Graphics
	x, y   float64
	radius float64
	color  float64
	el     *list.Element
}

func newParticle() *Particle {
	b := &Particle{
		Graphics: pixi.NewGraphics(),
		x:        0,
		y:        0,
		radius:   50,
		color:    COLOURS[rand.Intn(7)],
	}
	b.el = ps.PushBack(b)
	return b
}

func (g *Particle) draw(t float64) {
	g.Clear()
	if g.radius < 0 {
		return
		g.remove()
	}
	g.radius -= rand.Float64()
	g.BeginFill(g.color, 0.5)
	g.DrawCircle(g.x, g.y, g.radius)
	g.EndFill()
}

func (p *Particle) remove() {
	ps.Remove(p.el)
}

func run(t float64) {
	for e := ps.Front(); e != nil; e = e.Next() {
		p := e.Value.(*Particle)
		p.draw(t)
	}
	renderer.Render(stage)
	raf.RequestAnimationFrame(run)
}

func main() {
	stage.Click(func(id *pixi.InteractionData) {
		p := newParticle()
		p.x = id.Global.X
		p.y = id.Global.Y
		stage.AddChild(p)
	})
	dom.OnDOMContentLoaded(func() {
		dom.Body().Style.SetProperty("margin", "0")
		renderer = pixi.AutoDetectRenderer(dom.Window().InnerWidth, dom.Window().InnerHeight)
		v := dom.Wrap(renderer.View)
		v.Width = dom.Window().InnerWidth
		v.Height = dom.Window().InnerHeight
		dom.Body().AppendChild(v)
		raf.RequestAnimationFrame(run)
	})
}
