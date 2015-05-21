package main

import (
	"container/list"
	"github.com/Archs/js/canvas"
	"github.com/Archs/js/dom"
	"github.com/Archs/js/raf"
	"math"
	"math/rand"
)

var (
	ps      = list.New()
	COLOURS = []string{"#69D2E7", "#A7DBD8", "#E0E4CC", "#F38630", "#FA6900", "#FF4E50", "#F9D423"}
	ctx     *canvas.Context2D
	cw, ch  float64
)

type Particle struct {
	*canvas.Context2D
	x, y    float64
	vx, vy  float64
	radius  float64
	theta   float64
	force   float64
	damping float64
	color   string
	el      *list.Element
}

func newParticle(x, y float64) *Particle {
	b := &Particle{
		Context2D: ctx,
		x:         x,
		y:         y,
		radius:    5 + 30*rand.Float64(),
		theta:     2 * math.Pi * rand.Float64(),
		force:     2 + 8*rand.Float64(),
		damping:   0.92,
		color:     COLOURS[rand.Intn(7)],
	}
	b.vx = b.force * math.Sin(b.theta)
	b.vy = b.force * math.Cos(b.theta)
	b.el = ps.PushBack(b)
	return b
}

func (g *Particle) draw(t float64) {
	g.radius *= 0.9
	if g.radius < 1 {
		g.remove()
		return
	}
	g.x += g.vx
	g.y += g.vy
	g.theta = 2 * math.Pi * rand.Float64()
	g.vx *= g.damping
	g.vy *= g.damping
	g.vx += math.Sin(g.theta) * 0.1
	g.vy += math.Cos(g.theta) * 0.1
	// for canvas
	g.BeginPath()
	g.Arc(g.x, g.y, g.radius, 0, math.Pi*2, true)
	g.FillStyle = g.color
	g.Fill()
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
		ctx.ClearRect(0, 0, cw, ch)
		for e := ps.Front(); e != nil; e = e.Next() {
			p := e.Value.(*Particle)
			p.draw(t)
		}
	}
}

func makeParticles(x, y float64, n int) {
	for i := 0; i < rand.Intn(n); i++ {
		newParticle(x, y)
	}
}

func main() {
	dom.OnDOMContentLoaded(func() {
		dom.Body().Style.SetProperty("margin", "0")
		el := canvas.New(dom.CreateElement("canvas").Object)
		cw = float64(dom.Window().InnerWidth)
		ch = float64(dom.Window().InnerHeight)
		// cw = 800
		// ch = 600
		el.Width = int(cw)
		el.Height = int(ch)
		el.AddEventListener(dom.EvtMousemove, func(e *dom.Event) {
			x := float64(e.ClientX)
			y := float64(e.ClientY)
			makeParticles(x, y, 5)
		})
		ctx = el.GetContext2D()
		dom.Body().AppendChild(el.Element)
		raf.RequestAnimationFrame(run)
	})
}
