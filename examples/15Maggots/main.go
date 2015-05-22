package main

import (
	"github.com/Archs/js/dom"
	"github.com/Archs/js/raf"
	"github.com/Archs/pixi"
	"math"
	"math/rand"
	"time"
)

var (
	stage    = pixi.NewContainer()
	sprites  = pixi.NewParticleContainer(10000)
	renderer = pixi.AutoDetectRenderer(800, 600)

	maggotNum = 100
)

type Maggot struct {
	*pixi.Sprite
	step      float64
	direction float64
}

func newMaggot() *Maggot {
	s := new(Maggot)
	s.Sprite = pixi.SpriteFromImage("img/tinyMaggot.png", false, pixi.ScaleModes.Default)
	s.direction = 2 * math.Pi * rand.Float64()
	s.Rotation = math.Pi - s.direction
	s.Scale.SetTo(0.4)
	s.Anchor.SetTo(0.5)
	s.Position.X = renderer.Width * rand.Float64()
	s.Position.Y = renderer.Height * rand.Float64()
	s.step = 2 + 5*rand.Float64()
	s.Update = s.run
	return s
}

func (s *Maggot) run(float64) {
	s.Position.X += (s.step * math.Sin(s.direction))
	s.Position.Y += (s.step * math.Cos(s.direction))
	if s.Position.X < 0 {
		s.Position.X += renderer.Width
	} else if s.Position.X > renderer.Width {
		s.Position.X -= renderer.Width
	}
	if s.Position.Y < 0 {
		s.Position.Y += renderer.Height
	} else if s.Position.Y > renderer.Height {
		s.Position.Y -= renderer.Height
	}
}

func animate(t float64) {
	raf.RequestAnimationFrame(animate)
	for _, s := range sprites.Children {
		s.Update(t)
	}
	renderer.Render(stage)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < maggotNum; i++ {
		sprites.AddChild(newMaggot())
	}
	stage.AddChild(sprites)
	dom.OnDOMContentLoaded(func() {
		el := dom.Wrap(renderer.View)
		dom.Body().AppendChild(el)
		raf.RequestAnimationFrame(animate)
	})
}
