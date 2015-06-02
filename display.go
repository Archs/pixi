package pixi

import "github.com/gopherjs/gopherjs/js"

type displayObject interface {
	displayer() *js.Object
}

type DisplayObject struct {
	*js.Object
	// The coordinate of the object relative to the local coordinates of the parent.
	Position Point `js:"position"`
	Scale    Point `js:"scale"`
	// The pivot point of the displayObject that it rotates around
	Pivot         Point   `js:"pivot"`
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
	// filterArea Rectangle
	//
	// The area the filter is applied to. This is used as more of an optimisation rather than figuring out the dimensions of the displayObject each frame you can set this rectangle
	FilterArea Rectangle `js:"filterArea"`
	//
	// 	filters Array.<Filter>
	//
	// Sets the filters for the displayObject.
	// IMPORTANT: This is a webGL only feature and will be ignored by the canvas renderer. To remove filters simply set this property to 'null'
	filters []*js.Object `js:"filters"`

	// mask Graphics
	//
	// Sets a mask for the displayObject.
	// A mask is an object that limits the visibility of an object to the shape of the mask applied to it.
	// In PIXI a regular mask must be a PIXI.Graphics object. This allows for much faster masking in canvas as it utilises shape clipping.
	// To remove a mask, set this property to null.
	Mask *Graphics `js:"mask"`
}

func wrapDisplayObject(object *js.Object) *DisplayObject {
	return &DisplayObject{
		Object: object,
	}
}

// displayer satisfies the displayObject interface.
func (d *DisplayObject) displayer() *js.Object {
	return d.Object
}

// Parent is the display object container that contains this display object.
func (d *DisplayObject) Parent() *Container {
	return wrapContainer(d.Get("parent"))
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
func (d *DisplayObject) GetBounds() Rectangle {
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

// generateTexture(renderer, resolution, scaleMode){Texture}
//
// Useful function that returns a texture of the display object that can then be used to create sprites This can be quite useful if your displayObject is static / complicated and needs to be reused multiple times.
//
// Name	Type	Description
// renderer	CanvasRenderer | WebGLRenderer
// The renderer used to generate the texture.
// resolution	Number
// The resolution of the texture being generated
// scaleMode	Number
// See SCALE_MODES for possible values
// Returns:
//
// Type	Description
// Texture	a texture of the display object
func (d *DisplayObject) GenerateTexture(renderer Renderer, resolution int, scaleMode int) *Texture {
	o := d.Call("generateTexture", renderer, resolution, scaleMode)
	return &Texture{Object: o}
}

func (d *DisplayObject) MouseDown(cb func(*InteractionData)) {
	d.Set("mousedown", func(data *js.Object) {
		cb(wrapInteractionData(data))
	})
}

func (d *DisplayObject) MouseUp(cb func(*InteractionData)) {
	d.Set("mouseup", func(data *js.Object) {
		cb(wrapInteractionData(data))
	})
}

func (d *DisplayObject) MouseUpOutside(cb func(*InteractionData)) {
	d.Set("mouseupoutside", func(data *js.Object) {
		cb(wrapInteractionData(data))
	})
}

func (d *DisplayObject) MouseOver(cb func(*InteractionData)) {
	d.Set("mouseover", func(data *js.Object) {
		cb(wrapInteractionData(data))
	})
}

func (d *DisplayObject) MouseOut(cb func(*InteractionData)) {
	d.Set("mouseout", func(data *js.Object) {
		cb(wrapInteractionData(data))
	})
}

func (d *DisplayObject) MouseMove(cb func(*InteractionData)) {
	d.Set("mousemove", func(data *js.Object) {
		cb(wrapInteractionData(data))
	})
}

func (d *DisplayObject) TouchStart(cb func(*InteractionData)) {
	d.Set("touchstart", func(data *js.Object) {
		cb(wrapInteractionData(data))
	})
}

func (d *DisplayObject) TouchEnd(cb func(*InteractionData)) {
	d.Set("touchend", func(data *js.Object) {
		cb(wrapInteractionData(data))
	})
}

func (d *DisplayObject) TouchEndOutside(cb func(*InteractionData)) {
	d.Set("touchendoutside", func(data *js.Object) {
		cb(wrapInteractionData(data))
	})
}

func (d *DisplayObject) TouchMove(cb func(*InteractionData)) {
	d.Set("touchmove", func(data *js.Object) {
		cb(wrapInteractionData(data))
	})
}

func (d *DisplayObject) Tap(cb func(*InteractionData)) {
	d.Set("tap", func(data *js.Object) {
		cb(wrapInteractionData(data))
	})
}

func (d *DisplayObject) Click(cb func(*InteractionData)) {
	d.Set("click", func(data *js.Object) {
		cb(wrapInteractionData(data))
	})
}

// TODO: mask
// TODO: filters

// A Container represents a collection of display objects.
type Container struct {
	*DisplayObject
	Width  float64 `js:"width"`
	Height float64 `js:"height"`
}

func NewContainer() *Container {
	return wrapContainer(pkg.Get("Container").New())
}

func wrapContainer(object *js.Object) *Container {
	return &Container{DisplayObject: wrapDisplayObject(object)}
}

// AddChild adds a child to the container.
func (d Container) AddChild(do displayObject) {
	d.Call("addChild", do.displayer())
}

// AddChildAt adds a child at the specified index.
func (d Container) AddChildAt(do displayObject, index int) {
	d.Call("addChildAt", do.displayer(), index)
}

// ChildAt returns the child at the specified index.
func (d Container) ChildAt(index int) *DisplayObject {
	return wrapDisplayObject(d.Call("getChildAt", index))
}

// RemoveChild removes a child from the container.
func (d Container) RemoveChild(do displayObject) {
	d.Call("removeChild", do.displayer())
}

// RemoveChildAt removes a child at the specified index.
func (d Container) RemoveChildAt(index int) {
	d.Call("removeChildAt", index)
}

// RemoveChildren removes all child instances from the container.
func (d Container) RemoveChildren(start, end int) {
	d.Call("removeChildren", start, end)
}

// RemoveChildren removes all child instances from the container.
func (d Container) RemoveAllChildren() {
	d.Call("removeChildren")
}

type Sprite struct {
	*Container
	// The anchor sets the origin point of the texture.
	// The default is 0,0 this means the texture's origin is the top left
	// Setting the anchor to 0.5,0.5 means the texture's origin is centered
	// Setting the anchor to 1,1 would mean the texture's origin point will be the bottom right corner
	Anchor Point `js:"anchor"`
	// The tint applied to the sprite.
	// This is a hex value. A value of 0xFFFFFF will remove any tint effect.
	Tint      uint32 `js:"tint"`
	BlendMode int    `js:"blendMode"`
}

func NewSprite(texture *Texture) *Sprite {
	object := pkg.Get("Sprite").New(texture.Object)
	return wrapSprite(object)
}

func wrapSprite(object *js.Object) *Sprite {
	return &Sprite{
		Container: wrapContainer(object),
		// Anchor:                 Point{Object: object.Get("anchor")},
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

type SpriteBatch struct {
	*js.Object
}

func NewSpriteBatch() *SpriteBatch {
	return &SpriteBatch{wrapContainer(pkg.Get("SpriteBatch").New()).Object}
}

type Stage struct {
	*Container
}

func NewStage(background uint32) *Stage {
	return wrapStage(pkg.Get("Stage").New(background))
}

func wrapStage(object *js.Object) *Stage {
	return &Stage{Container: wrapContainer(object)}
}

type MovieClip struct {
	*Sprite
	AnimationSpeed int  `js:"animationSpeed"`
	Loop           bool `js:"loop"`
}

func NewMovieClip(textures []*Texture) *MovieClip {
	objs := make([]*js.Object, 0, len(textures))
	for _, t := range textures {
		objs = append(objs, t.Object)
	}

	return &MovieClip{
		Sprite: wrapSprite(pkg.Get("MovieClip").New(objs)),
	}
}

func (m *MovieClip) OnComplete(cb func()) {
	m.Set("onComplete", cb)
}

func (m *MovieClip) CurrentFrame() float64 {
	return m.Get("currentFrame").Float()
}

func (m *MovieClip) Playing() bool {
	return m.Get("playing").Bool()
}

func (m *MovieClip) TotalFrames() int {
	return m.Get("totalFrames").Int()
}

// Goes to a specific frame and begins playing the MovieClip
func (m *MovieClip) GotoAndPlay(frameNumber int) {
	m.Call("gotoAndPlay", frameNumber)
}

func MovieClipFromImages(urls []string) *MovieClip {
	return &MovieClip{
		Sprite: wrapSprite(pkg.Get("MovieClip").Call("fromImages", urls)),
	}
}
