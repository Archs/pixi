package main

import (
	"math/rand"
	"time"

	"github.com/Archs/js/raf"
	"github.com/Archs/pixi"

	"github.com/gopherjs/gopherjs/js"
)

var (
	width  = 1024.0
	height = 728.0
	slideX = width / 2
	slideY = height / 2
	sx     = 1.0
	sy     = 1.0

	stage    = pixi.NewContainer(0x000000)
	renderer = pixi.AutoDetectRenderer(int(width), int(height))
	texture  = pixi.TextureFromImage("bubble_32x32.png", false, pixi.ScaleModes.Default)
	balls    = make([]*Ball, 0)
)

type Ball struct {
	sprite *pixi.Sprite
	x, y   float64
}

func animate(t float64) {
	for i := 0; i < len(balls); i++ {
		ball := balls[i]
		ball.sprite.Position.X = ball.x + slideX
		ball.sprite.Position.Y = ball.y + slideY
		ball.x *= sx
		ball.y *= sy

		if ball.x > width {
			ball.x -= width
		} else if ball.x < -width {
			ball.x += width
		}

		if ball.y > height {
			ball.y -= height
		} else if ball.y < -height {
			ball.y += height
		}
	}

	renderer.Render(stage)
	raf.RequestAnimationFrame(animate)
}

func main() {
	js.Global.Get("document").Get("body").Call("appendChild", renderer.View)

	rand.Seed(int64(time.Now().Nanosecond()))
	sx = 1.0 + rand.Float64()/20
	sy = 1.0 + rand.Float64()/20

	for i := 0; i < 2500; i++ {
		ball := pixi.NewSprite(texture)
		ball.Position.Set(rand.Float64()*width-slideX, rand.Float64()*height-slideY)
		ball.Anchor.SetTo(0.5)
		balls = append(balls, &Ball{ball, ball.Position.X, ball.Position.Y})
		stage.AddChild(ball)
	}

	raf.RequestAnimationFrame(animate)
}
