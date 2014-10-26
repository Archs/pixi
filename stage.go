package pixi

import "github.com/gopherjs/gopherjs/js"

type Stage struct {
	js.Object
}

func (s Stage) AddChild(sprite Sprite) {
	s.Object.Call("addChild", sprite.Object)
}

func NewStage(background uint32) Stage {
	return Stage{Object: pkg.Get("Stage").New(background)}
}
