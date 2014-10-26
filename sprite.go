package pixi

type Sprite struct {
	*DisplayObjectContainer
	Anchor Point
}

func NewSprite(texture Texture) Sprite {
	object := pkg.Get("Sprite").New(texture.Object)
	return Sprite{
		DisplayObjectContainer: wrapDisplayObjectContainer(object),
		Anchor:                 Point{Object: object.Get("anchor")},
	}
}
