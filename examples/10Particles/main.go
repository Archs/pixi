package main

import (
	"container/list"
	"github.com/Archs/js/dom"
	"github.com/Archs/js/raf"
	"github.com/Archs/pixi"
	"math"
	// "math"
	"math/rand"
)

var (
	stage    = pixi.NewStage(0x0)
	renderer pixi.Renderer
	ps       = list.New()
	ctx      = pixi.NewGraphics()
	COLOURS  = []float64{0x69D2E7, 0xA7DBD8, 0xE0E4CC, 0xF38630, 0xFA6900, 0xFF4E50, 0xF9D423}
)

type Particle struct {
	*pixi.Graphics
	x, y    float64
	vx, vy  float64
	radius  float64
	theta   float64
	force   float64
	damping float64
	color   float64
	el      *list.Element
}

func newParticle(x, y float64) *Particle {
	b := &Particle{
		Graphics: pixi.NewGraphics(),
		x:        x,
		y:        y,
		radius:   5 + 30*rand.Float64(),
		theta:    2 * math.Pi * rand.Float64(),
		force:    2 + 8*rand.Float64(),
		damping:  0.9 + 0.1*rand.Float64(),
		color:    COLOURS[rand.Intn(7)],
	}
	b.vx = b.force * math.Sin(b.theta)
	b.vy = b.force * math.Cos(b.theta)
	b.el = ps.PushBack(b)
	return b
}

func (g *Particle) draw(t float64) {
	g.Clear()
	// update
	g.radius *= 0.96
	if g.radius < 1 {
		g.remove()
		return
	}
	g.x += g.vx
	g.y += g.vy
	// g.theta = 2 * math.Pi * rand.Float64()
	g.vx *= g.damping
	g.vy *= g.damping
	// g.vx += math.Sin(g.theta) * 0.1
	// g.vy += math.Cos(g.theta) * 0.1
	// draw
	g.BeginFill(g.color, 1)
	g.DrawCircle(g.x, g.y, g.radius)
	g.EndFill()
}

func (p *Particle) remove() {
	ps.Remove(p.el)
}

func run(t float64) {
	defer raf.RequestAnimationFrame(run)
	n := int64(t)
	// frame control
	if n%3 != 0 {
		return
	}
	if ps.Len() > 0 {
		for e := ps.Front(); e != nil; e = e.Next() {
			p := e.Value.(*Particle)
			p.draw(t)
		}
	}
	renderer.Render(stage)
}

func makeParticles(x, y float64, n int) {
	for i := 0; i < rand.Intn(n); i++ {
		p := newParticle(x, y)
		stage.AddChild(p)
	}
}

func main() {
	stage.MouseMove(func(id *pixi.InteractionData) {
		makeParticles(id.Global.X, id.Global.Y, 4)
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
