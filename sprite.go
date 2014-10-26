package pixi

import "github.com/gopherjs/gopherjs/js"

type Sprite struct {
	js.Object
	Rotation float32 `js:"rotation"`
	Anchor   Point
	Position Point
}

func NewSprite(texture Texture) Sprite {
	object := pkg.Get("Sprite").New(texture.Object)
	return Sprite{
		Object:   object,
		Anchor:   Point{Object: object.Get("anchor")},
		Position: Point{Object: object.Get("position")},
	}
}
