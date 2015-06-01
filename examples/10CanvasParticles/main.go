package main

import (
	"github.com/Archs/js/canvas"
	"github.com/Archs/js/dom"
	"github.com/Archs/js/raf"
	"math"
	"math/rand"
)

const (
	MAX_PARTICLES = 280
)

var (
	COLOURS = []string{"#69D2E7", "#A7DBD8", "#E0E4CC", "#F38630", "#FA6900", "#FF4E50", "#F9D423"}
	ctx     *canvas.Context2D
	cw, ch  float64
	counter = 0
	ps      = []*Particle{}
	pool    = []*Particle{}
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
}

func getParticle(x, y float64) *Particle {
	var p *Particle
	if len(ps) >= MAX_PARTICLES {
		pool = append(pool, ps[0])
		ps = ps[1:]
	}
	if len(pool) > 0 {
		p = pool[0]
		pool = pool[1:]
	} else {
		p = &Particle{
			Context2D: ctx,
		}
	}
	p.init(x, y)
	return p
}

func (p *Particle) init(x, y float64) {
	p.x = x
	p.y = y
	p.radius = 5 + 35*rand.Float64()
	p.theta = 2 * math.Pi * rand.Float64()
	p.force = 2 + 8*rand.Float64()
	p.damping = 0.9 + 0.1*rand.Float64()
	p.color = COLOURS[rand.Intn(7)]
	p.vx = p.force * math.Sin(p.theta)
	p.vy = p.force * math.Cos(p.theta)
}

func (p *Particle) isAlive() bool {
	return p.radius > 0.5
}

func (g *Particle) update() {
	// update
	g.radius *= 0.93
	g.x += g.vx
	g.y += g.vy
	// g.theta += (0.5 - rand.Float64()) * 0.15
	g.vx *= g.damping
	g.vy *= g.damping
	g.vx += (math.Sin(g.theta) * 0.1)
	g.vy += (math.Cos(g.theta) * 0.1)
	// g.vx += (rand.Float64() * 0.1)
	// g.vy += (rand.Float64() * 0.1)
}

func (g *Particle) draw(t float64) {
	// for canvas
	g.BeginPath()
	g.Arc(g.x, g.y, g.radius, 0, math.Pi*2, true)
	g.FillStyle = g.color
	g.Fill()
}

func run(t float64) {
	raf.RequestAnimationFrame(run)
	counter += 1
	// frame control
	if counter%2 == 0 {
		return
	}
	// update & draw all the particles
	ctx.ClearRect(0, 0, cw, ch)
	for i := len(ps) - 1; i >= 0; i-- {
		p := ps[i]
		p.update()
		if !p.isAlive() {
			ps = append(ps[:i], ps[i+1:]...)
		} else {
			p.draw(t)
		}
	}
}

func makeParticles(x, y float64, n int) {
	println("makeParticles:", x, y, n)
	for i := 0; i < rand.Intn(n); i++ {
		p := getParticle(x, y)
		ps = append(ps, p)
	}
}

func main() {
	dom.OnDOMContentLoaded(func() {
		s := dom.Body()
		// set full window and black background
		s.Style.SetProperty("margin", "0")
		s.Style.SetProperty("background", "#222")

		el := canvas.New(dom.CreateElement("canvas").Object)
		cw = float64(dom.Window().InnerWidth)
		ch = float64(dom.Window().InnerHeight)
		el.Width = int(cw)
		el.Height = int(ch)
		el.AddEventListener(dom.EvtMousemove, func(e *dom.Event) {
			e.PreventDefault()
			x := float64(e.ClientX)
			y := float64(e.ClientY)
			makeParticles(x, y, 5)
		})
		ctx = el.GetContext2D()
		ctx.GlobalCompositeOperation = canvas.CompositeLighter
		dom.Body().AppendChild(el.Element)
		raf.RequestAnimationFrame(run)
	})
}
