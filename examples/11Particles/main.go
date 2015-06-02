package main

import (
	"github.com/Archs/js/dom"
	"github.com/Archs/js/raf"
	"github.com/Archs/pixi"
	"math"
	// "math"
	"math/rand"
)

const (
	MAX_PARTICLES = 280
)

var (
	stage    = pixi.NewContainer()
	renderer *pixi.Renderer
	// ctx      = pixi.NewGraphics()
	COLOURS = []float64{0x69D2E7, 0xA7DBD8, 0xE0E4CC, 0xF38630, 0xFA6900, 0xFF4E50, 0xF9D423}
	ps      = []*Particle{}
	pool    = []*Particle{}
	counter = 0
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
			Graphics: pixi.NewGraphics(),
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
}

func (g *Particle) draw(t float64) {
	g.Clear()
	// draw
	g.BeginFill(g.color, 1)
	g.DrawCircle(g.x, g.y, g.radius)
	g.EndFill()
}

func run(t float64) {
	raf.RequestAnimationFrame(run)
	counter += 1
	// frame control
	if counter%2 == 0 {
		return
	}
	// update & draw all the particles
	// ctx.Clear()
	for i := len(ps) - 1; i >= 0; i-- {
		p := ps[i]
		p.update()
		if !p.isAlive() {
			ps = append(ps[:i], ps[i+1:]...)
			stage.RemoveChild(p)
		} else {
			p.draw(t)
		}
	}
	renderer.Render(stage)
}

func makeParticles(x, y float64, n int) {
	for i := 0; i < rand.Intn(n); i++ {
		p := getParticle(x, y)
		p.BlendMode = pixi.BlendModes.Screen
		ps = append(ps, p)
		stage.AddChild(p)
	}
}

func main() {
	stage.Interactive = true
	// stage.AddChild(ctx)
	stage.MouseMove(func(ed *pixi.InteractionEvent) {
		id := ed.Data
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
