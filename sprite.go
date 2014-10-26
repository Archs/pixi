package pixi

import "github.com/gopherjs/gopherjs/js"

type Sprite struct {
	*DisplayObjectContainer
	Anchor    Point
	Tint      uint32 `js:"tint"`
	BlendMode int    `js:"blendMode"`
}

func NewSprite(texture *Texture) *Sprite {
	object := pkg.Get("Sprite").New(texture.Object)
	return wrapSprite(object)
}

func wrapSprite(object js.Object) *Sprite {
	return &Sprite{
		DisplayObjectContainer: wrapDisplayObjectContainer(object),
		Anchor:                 Point{Object: object.Get("anchor")},
	}
}

// SetTexture sets the texture of the sprite.
func (s *Sprite) SetTexture(texture *Texture) {
	s.Call("setTexture", texture.Object)
}

func SpriteFromFrame(frameId string) *Sprite {
	return wrapSprite(pkg.Get("Sprite").Call("fromFrame", frameId))
}

func SpriteFromImage(imageId string, crossOrigin bool, scaleMode int) *Sprite {
	return wrapSprite(pkg.Get("Sprite").Call("fromImage", imageId, crossOrigin, scaleMode))
}
