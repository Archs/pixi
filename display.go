package pixi

import "github.com/gopherjs/gopherjs/js"

type displayObject interface {
	displayer() js.Object
}

type DisplayObject struct {
	js.Object
	Position      Point
	Scale         Point
	Pivot         Point
	Rotation      float64 `js:"rotation"`
	Alpha         float64 `js:"alpha"`
	Visible       bool    `js:"visible"`
	ButtonMode    bool    `js:"buttonMode"`
	Renderable    bool    `js:"renderable"`
	Interactive   bool    `js:"interactive"`
	DefaultCursor string  `js:"defaultCursor"`
	CacheAsBitmap bool    `js:"cacheAsBitmap"`
	X             float64 `js:"x"`
	Y             float64 `js:"y"`
}

func wrapDisplayObject(object js.Object) *DisplayObject {
	return &DisplayObject{
		Object:   object,
		Position: Point{Object: object.Get("position")},
		Scale:    Point{Object: object.Get("scale")},
		Pivot:    Point{Object: object.Get("pivot")},
	}
}

// displayer satisfies the displayObject interface.
func (d *DisplayObject) displayer() js.Object {
	return d.Object
}

// Parent is the display object container that contains this display object.
func (d *DisplayObject) Parent() *DisplayObjectContainer {
	return wrapDisplayObjectContainer(d.Get("parent"))
}

// Stage the display object is connected to.
func (d *DisplayObject) Stage() *Stage {
	return wrapStage(d.Get("stage"))
}

// WorldAlpha is the multiplied alpha of the DisplayObject.
func (d *DisplayObject) WorldAlpha() float64 {
	return d.Get("worldAlpha").Float()
}

// WorldVisible indicates if the sprite is globaly visible.
func (d *DisplayObject) WorldVisible() bool {
	return d.Get("worldVisible").Bool()
}

// Bounds is the bounds of the DisplayObject as a rectangle object.
func (d *DisplayObject) Bounds() Rectangle {
	return Rectangle{Object: d.Call("getBounds")}
}

// LocalBounds is the local bounds of the DisplayObject as a rectangle object.
func (d *DisplayObject) LocalBounds() Rectangle {
	return Rectangle{Object: d.Call("getLocalBounds")}
}

// SetStageReference sets the object's stage reference.
func (d *DisplayObject) SetStageReference(stage *Stage) {
	d.Call("setStageReference", stage.Object)
}

// RemoveStageReference removes the object's stage reference.
func (d *DisplayObject) RemoveStageReference() {
	d.Call("removeStageReference")
}

// SetFilterArea sets the area the filter is applied to.
func (d *DisplayObject) SetFilterArea(rectangle Rectangle) {
	d.Set("filterArea", rectangle.Object)
}

// TODO: mask
// TODO: filters

// A DisplayObjectContainer represents a collection of display objects.
type DisplayObjectContainer struct {
	*DisplayObject
	Width  float64 `js:"width"`
	Height float64 `js:"height"`
}

func NewDisplayObjectContainer() *DisplayObjectContainer {
	return wrapDisplayObjectContainer(pkg.Get("DisplayObjectContainer").New())
}

func wrapDisplayObjectContainer(object js.Object) *DisplayObjectContainer {
	return &DisplayObjectContainer{DisplayObject: wrapDisplayObject(object)}
}

// AddChild adds a child to the container.
func (d DisplayObjectContainer) AddChild(do displayObject) {
	d.Call("addChild", do.displayer())
}

// AddChildAt adds a child at the specified index.
func (d DisplayObjectContainer) AddChildAt(do displayObject, index int) {
	d.Call("addChildAt", do.displayer(), index)
}

// ChildAt returns the child at the specified index.
func (d DisplayObjectContainer) ChildAt(index int) *DisplayObject {
	return wrapDisplayObject(d.Call("getChildAt", index))
}

// RemoveChild removes a child from the container.
func (d DisplayObjectContainer) RemoveChild(do displayObject) {
	d.Call("removeChild", do.displayer())
}

// RemoveChildAt removes a child at the specified index.
func (d DisplayObjectContainer) RemoveChildAt(index int) {
	d.Call("removeChildAt", index)
}

// RemoveChildren removes all child instances from the container.
func (d DisplayObjectContainer) RemoveChildren(start, end int) {
	d.Call("removeChildren", start, end)
}

// RemoveChildren removes all child instances from the container.
func (d DisplayObjectContainer) RemoveAllChildren() {
	d.Call("removeChildren")
}

type SpriteBatch struct {
	js.Object
}

func NewSpriteBatch() *SpriteBatch {
	return &SpriteBatch{wrapDisplayObjectContainer(pkg.Get("SpriteBatch").New())}
}
