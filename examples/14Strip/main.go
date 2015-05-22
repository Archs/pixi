package main

import (
	"github.com/Archs/js/dom"
	"github.com/Archs/js/raf"
	"github.com/Archs/pixi"
	"math"
)

var (
	stage    = pixi.NewContainer()
	renderer = pixi.AutoDetectRenderer(800, 600, 0xffffff)

	snakeLen = 918.0

	snake *Snake
)

type Snake struct {
	*pixi.Rope
	Points  []*pixi.Point
	Texture *pixi.Texture
	counter float64
	ropeLen float64
}

func newSnake() *Snake {
	s := new(Snake)
	s.ropeLen = snakeLen / 20.0
	s.counter = 0
	s.Points = make([]*pixi.Point, 0)
	for i := 0; i < 20; i++ {
		s.Points = append(s.Points, pixi.NewPoint(float64(i)*s.ropeLen, 0))
	}
	s.Texture = pixi.TextureFromImage("img/snake.png", false, pixi.ScaleModes.Default)
	s.Rope = pixi.NewRope(s.Texture, s.Points)
	s.Scale.SetTo(0.2)
	s.Update = s.roll
	return s
}

func (s *Snake) roll() {
	s.counter += 0.2
	for i, p := range s.Points {
		p.Y = math.Sin(s.counter+float64(i)*0.5) * 30
		p.X = float64(i)*s.ropeLen + math.Sin(s.counter+float64(i)*0.5)*30
	}
	s.Position.X += 4
	if s.Position.X > renderer.Width+s.Width/2 {
		s.Position.X = -1*s.Width + 10
	}
}

func animate(t float64) {
	raf.RequestAnimationFrame(animate)
	for _, s := range stage.Children {
		s.Update()
	}
	// move the snake
	renderer.Render(stage)
}

func main() {
	snake = newSnake()
	snake.Position.Set(120, 300)
	stage.AddChild(snake)
	dom.OnDOMContentLoaded(func() {
		el := dom.Wrap(renderer.View)
		dom.Body().AppendChild(el)
		raf.RequestAnimationFrame(animate)
	})
}
