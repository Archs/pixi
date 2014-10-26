package pixi

import "github.com/gopherjs/gopherjs/js"

type displayObject interface {
	displayer() js.Object
}

type DisplayObject struct {
	js.Object
	Position Point
	Scale    Point
	Pivot    Point
	Rotation float32 `js:"rotation"`
	Alpha    float32 `js:"alpha"`
	Visible  bool    `js:"visible"`
}

func wrapDisplayObject(object js.Object) *DisplayObject {
	return &DisplayObject{
		Object:   object,
		Position: Point{Object: object.Get("position")},
		Scale:    Point{Object: object.Get("scale")},
		Pivot:    Point{Object: object.Get("pivot")},
	}
}

func (d *DisplayObject) displayer() js.Object {
	return d.Object
}

type DisplayObjectContainer struct {
	*DisplayObject
}

func wrapDisplayObjectContainer(object js.Object) *DisplayObjectContainer {
	return &DisplayObjectContainer{DisplayObject: wrapDisplayObject(object)}
}

func (d DisplayObjectContainer) AddChild(do displayObject) {
	d.Call("addChild", do.displayer())
}
