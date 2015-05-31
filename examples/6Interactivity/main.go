package main

import (
	"github.com/Archs/js/raf"
	"github.com/Archs/pixi"

	"github.com/gopherjs/gopherjs/js"
)

var (
	stage      = pixi.NewStage(0x000000)
	renderer   = pixi.AutoDetectRenderer(620, 400)
	background = pixi.SpriteFromImage("bg.jpg", true, pixi.ScaleModes.Default)
	buttons    = make([]*Button, 0)
)

type Button struct {
	*pixi.Sprite
	isDown  bool
	isOver  bool
	upTex   *pixi.Texture
	downTex *pixi.Texture
	overTex *pixi.Texture
}

func NewButton(x, y float64, upTex, downTex, overTex *pixi.Texture) *Button {
	sprite := pixi.NewSprite(upTex)
	sprite.ButtonMode = true

	sprite.Anchor.SetTo(0.5)
	sprite.Position.Set(x, y)

	sprite.Interactive = true

	button := &Button{
		Sprite:  sprite,
		isDown:  false,
		isOver:  false,
		upTex:   upTex,
		downTex: downTex,
		overTex: overTex,
	}

	button.MouseUp(button.up)
	button.TouchEnd(button.up)
	button.MouseUpOutside(button.up)
	button.TouchEndOutside(button.up)

	button.MouseDown(button.down)
	button.TouchStart(button.down)

	button.MouseOver(button.over)
	button.MouseOut(button.out)

	return button
}

func (button *Button) OnTap(cb func(*pixi.InteractionData)) {
	button.Click(cb)
	button.Tap(cb)
}

func (button *Button) up(data *pixi.InteractionData) {
	button.isDown = false
	if button.isOver {
		button.SetTexture(button.overTex)
	} else {
		button.SetTexture(button.upTex)
	}
}

func (button *Button) down(data *pixi.InteractionData) {
	button.isDown = true
	button.SetTexture(button.downTex)
}

func (button *Button) over(data *pixi.InteractionData) {
	button.isOver = true
	if !button.isDown {
		button.SetTexture(button.overTex)
	}
}

func (button *Button) out(data *pixi.InteractionData) {
	button.isOver = false
	if !button.isDown {
		button.SetTexture(button.upTex)
	}
}

func animate(t float64) {
	raf.RequestAnimationFrame(animate)
	renderer.Render(stage)
}

func main() {
	js.Global.Get("document").Get("body").Call("appendChild", renderer.View)

	stage.AddChild(background)

	upTex := pixi.TextureFromImage("buttonUp.png", true, pixi.ScaleModes.Default)
	downTex := pixi.TextureFromImage("buttonDown.png", true, pixi.ScaleModes.Default)
	overTex := pixi.TextureFromImage("buttonOver.png", true, pixi.ScaleModes.Default)

	coords := []float64{
		175.0, 75.0,
		600 - 145, 75,
		600/2 - 20, 400/2 + 10,
		175, 400 - 75,
		600 - 115, 400 - 95,
	}

	for i := 0; i < len(coords)/2; i++ {
		button := NewButton(coords[i*2], coords[i*2+1], upTex, downTex, overTex)

		button.OnTap(func(data *pixi.InteractionData) {
			println("CLICK")
		})

		stage.AddChild(button)
		buttons = append(buttons, button)
	}

	buttons[0].Scale.X = 1.2
	buttons[1].Scale.Y = 1.2
	buttons[2].Rotation = 0.314159
	buttons[3].Scale.SetTo(0.8)
	buttons[4].Scale.Set(0.8, 1.2)
	buttons[4].Rotation = 3.14159

	raf.RequestAnimationFrame(animate)
}
