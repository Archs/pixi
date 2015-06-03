package pixi

import "github.com/gopherjs/gopherjs/js"

const (
	EventClick           = "click"
	EventMouseDown       = "mousedown"
	EventMouseOut        = "mouseout"
	EventMouseOver       = "mouseover"
	EventMouseUp         = "mouseup"
	EventMouseClick      = "mouseclick"
	EventMouseUpOutside  = "mouseupoutside"
	EventMouseMove       = "mousemove"
	EventRightClick      = "rightclick"
	EventRightDown       = "rightdown"
	EventRightUp         = "rightup"
	EventRightUpoutside  = "rightupoutside"
	EventTap             = "tap"
	EventTouchEnd        = "touchend"
	EventTouchEndOutside = "touchendoutside"
	EventTouchMove       = "touchmove"
	EventTouchStart      = "touchstart"
)

type displayObject interface {
	displayer() *js.Object
}

type DisplayObject struct {
	*js.Object
	Name string `js:"name"`
	// The coordinate of the object relative to the local coordinates of the parent.
	Position *Point `js:"position"`
	Scale    *Point `js:"scale"`
	// The pivot point of the displayObject that it rotates around
	Pivot               *Point  `js:"pivot"`
	Rotation            float64 `js:"rotation"`
	Alpha               float64 `js:"alpha"`
	Visible             bool    `js:"visible"`
	ButtonMode          bool    `js:"buttonMode"`
	Renderable          bool    `js:"renderable"`
	Interactive         bool    `js:"interactive"`
	InteractiveChildren bool    `js:"interactiveChildren"`
	DefaultCursor       string  `js:"defaultCursor"`

	// Set CacheAsBitmap to true if you want this display object to be cached as a bitmap.
	// * This basically takes a snap shot of the display object as it is at that moment.
	//   It can provide a performance benefit for complex static displayObjects.
	// * When cacheAsBitmap is set to true the graphics object will be rendered as if it was a sprite.
	// * This is useful if your graphics element does not change often, as it will speed up the rendering
	// * of the object in exchange for taking up texture memory. It is also useful if you need the graphics
	// * object to be anti-aliased, because it will be rendered using canvas. This is not recommended if
	// * you are constantly redrawing the graphics element.
	CacheAsBitmap bool    `js:"cacheAsBitmap"`
	X             float64 `js:"x"`
	Y             float64 `js:"y"`

	// for pure containers only, indicates the mouse/touch event response area
	// Sprite Graphics don't need this
	HitArea *Rectangle `js:"hitArea"`

	// The area the filter is applied to. This is used as more of an optimisation
	// rather than figuring out the dimensions of the displayObject each frame
	// you can set this rectangle
	FilterArea Rectangle `js:"filterArea"`
	//
	// 	filters Array.<Filter>
	//
	// Sets the filters for the displayObject.
	// IMPORTANT: This is a webGL only feature and will be ignored by the canvas renderer. To remove filters simply set this property to 'null'
	filters []*js.Object `js:"filters"`

	// Sets a mask for the displayObject.
	// A mask is an object that limits the visibility of an object to the shape of the mask applied to it.
	// In PIXI a regular mask must be a PIXI.Graphics object. This allows for much faster masking in canvas as it utilises shape clipping.
	// To remove a mask, set this property to null.
	Mask *Graphics `js:"mask"`

	// a gopherjs specific DisplayObject update function
	// which can be used in the container.Children/ChildAt/GetChildByName returned instance
	//
	// Must be set before invoke
	Update func(float64) `js:"gopherjsSpecificUpdater"`
}

func wrapDisplayObject(object *js.Object) *DisplayObject {
	d := &DisplayObject{
		Object: object,
	}
	d.Update = func(float64) { println("no updater defined for", d.Name) }
	return d
}

// displayer satisfies the displayObject interface.
func (d *DisplayObject) displayer() *js.Object {
	return d.Object
}

// Parent is the display object container that contains this display object.
func (d *DisplayObject) Parent() *Container {
	return wrapContainer(d.Get("parent"))
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

// destroy()
// Base destroy method for generic display objects
func (d *DisplayObject) Destroy() {
	d.Call("destroy")
}

// // SetStageReference sets the object's stage reference.
// func (d *DisplayObject) SetStageReference(stage *Stage) {
// 	d.Call("setStageReference", stage.Object)
// }

// // RemoveStageReference removes the object's stage reference.
// func (d *DisplayObject) RemoveStageReference() {
// 	d.Call("removeStageReference")
// }

// // SetFilterArea sets the area the filter is applied to.
// func (d *DisplayObject) SetFilterArea(rectangle Rectangle) {
// 	d.Set("filterArea", rectangle.Object)
// }

// toGlobal(position){Point}
//
// Calculates the global position of the display object
//
//  Name	Type	Description
//  position	Point
//  The world origin to calculate from
//  Returns:
//
//  Type	Description
//  Point	A point object representing the position of this object
func (d *DisplayObject) ToGlobal(position Point) *Point {
	o := d.Call("toGlobal", position)
	return &Point{Object: o}
}

// getGlobalPosition(point: Point): Point;
func (d *DisplayObject) GetGlobalPosition(position Point) *Point {
	o := d.Call("getGlobalPosition", position)
	return &Point{Object: o}
}

// toLocal(position, from){Point}
//
// Calculates the local position of the display object relative to another point
//
//	 Name	Type	Description
//	 position	Point
//	 The world origin to calculate from
//	 from	DisplayObject	optional
//	 The DisplayObject to calculate the global position from
//	 Returns:
//
//	 Type	Description
//	 Point	A point object representing the position of this object
func (d *DisplayObject) ToLocal(position, from Point) *Point {
	o := d.Call("toLocal", position, from)
	return &Point{Object: o}
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

// on(event: 'click', fn: (event: interaction.InteractionEvent) => void, context?: any): EventEmitter;
// on(event: 'mousedown', fn: (event: interaction.InteractionEvent) => void, context?: any): EventEmitter;
// on(event: 'mouseout', fn: (event: interaction.InteractionEvent) => void, context?: any): EventEmitter;
// on(event: 'mouseover', fn: (event: interaction.InteractionEvent) => void, context?: any): EventEmitter;
// on(event: 'mouseup', fn: (event: interaction.InteractionEvent) => void, context?: any): EventEmitter;
// on(event: 'mouseclick', fn: (event: interaction.InteractionEvent) => void, context?: any): EventEmitter;
// on(event: 'mouseupoutside', fn: (event: interaction.InteractionEvent) => void, context?: any): EventEmitter;
// on(event: 'rightclick', fn: (event: interaction.InteractionEvent) => void, context?: any): EventEmitter;
// on(event: 'rightdown', fn: (event: interaction.InteractionEvent) => void, context?: any): EventEmitter;
// on(event: 'rightup', fn: (event: interaction.InteractionEvent) => void, context?: any): EventEmitter;
// on(event: 'rightupoutside', fn: (event: interaction.InteractionEvent) => void, context?: any): EventEmitter;
// on(event: 'tap', fn: (event: interaction.InteractionEvent) => void, context?: any): EventEmitter;
// on(event: 'touchend', fn: (event: interaction.InteractionEvent) => void, context?: any): EventEmitter;
// on(event: 'touchendoutside', fn: (event: interaction.InteractionEvent) => void, context?: any): EventEmitter;
// on(event: 'touchmove', fn: (event: interaction.InteractionEvent) => void, context?: any): EventEmitter;
// on(event: 'touchstart', fn: (event: interaction.InteractionEvent) => void, context?: any): EventEmitter;
// on(event: string, fn: Function, context?: any): EventEmitter;
func (d *DisplayObject) On(eventName string, cb func(*InteractionEvent)) *DisplayObject {
	d.Call("on", eventName, cb)
	return d
}

func (d *DisplayObject) Once(eventName string, cb func(*InteractionEvent)) *DisplayObject {
	d.Call("once", eventName, cb)
	return d
}

// A Container represents a collection of display objects.
type Container struct {
	*DisplayObject
	Children []*DisplayObject `js:"children"`
	Width    float64          `js:"width"`
	Height   float64          `js:"height"`
}

func NewContainer() *Container {
	return wrapContainer(pkg.Get("Container").New())
}

func wrapContainer(object *js.Object) *Container {
	return &Container{DisplayObject: wrapDisplayObject(object)}
}

// AddChild adds a child to the container.
func (d Container) AddChild(do ...displayObject) {
	for _, v := range do {
		d.Call("addChild", v.displayer())
	}
}

// AddChildAt adds a child at the specified index.
func (d Container) AddChildAt(do displayObject, index int) {
	d.Call("addChildAt", do.displayer(), index)
}

// getChildByName(name: string): DisplayObject;
func (d Container) GetChildByName(name string) *DisplayObject {
	return wrapDisplayObject(d.Call("getChildByName", name))
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
	// Name	Type	Description
	// texture	Texture
	// the texture of the tiling sprite
	Texture *Texture `js:"texture"`
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

func SpriteFromFrame(frameId string) *Sprite {
	return wrapSprite(pkg.Get("Sprite").Call("fromFrame", frameId))
}

func SpriteFromImage(imageId string, crossOrigin bool, scaleMode int) *Sprite {
	return wrapSprite(pkg.Get("Sprite").Call("fromImage", imageId, crossOrigin, scaleMode))
}

// export interface ParticleContainerProperties {
//
//     scale?: any;
//     position?: any;
//     rotation?: number;
//     uvs?: any;
//     alpha?: number;
// }
type ParticleContainerProperties struct {
	*js.Object
	// scale	boolean	false	optional
	// When true, scale be uploaded and applied.
	Scale bool `js:"scale"`
	// position	boolean	true	optional
	// When true, position be uploaded and applied.
	Position bool `js:"position"`
	// rotation	boolean	false	optional
	// When true, rotation be uploaded and applied.
	Rotation bool `js:"rotation"`
	// uvs	boolean	false	optional
	// When true, uvs be uploaded and applied.
	Uvs bool `js:"uvs"`
	// alpha	boolean	false	optional
	// When true, alpha be uploaded and applied.
	Alpha bool `js:"alpha"`
}

func NewParticleContainerProperties() *ParticleContainerProperties {
	return &ParticleContainerProperties{
		Object: js.Global.Get("Object").New(),
	}
}

// export class ParticleContainer extends Container {
//
//         constructor(size?: number, properties?: ParticleContainerProperties);
//
//         interactiveChildren: boolean;
//         blendMode: number;
//         roundPixels: boolean;
//
//         setProperties(properties: ParticleContainerProperties): void;
//         addChildAt(child: DisplayObject, index: number): DisplayObject;
//         removeChildAt(index: number): DisplayObject;
//
//     }

// The ParticleContainer class is a really fast version of the Container built solely for speed,
// so use when you need a lot of sprites or particles.
// The tradeoff of the ParticleContainer is that advanced functionality will not work.
// ParticleContainer implements only the basic object transform (position, scale, rotation).
// Any other functionality like tinting, masking, etc will not work on sprites in this batch.
//
// It's extremely easy to use :
//
//	 var container = new ParticleContainer();
//
//	 for (var i = 0; i < 100; ++i)
//	 {
//	     var sprite = new PIXI.Sprite.fromImage("myImage.png");
//	     container.addChild(sprite);
//	 }
type ParticleContainer struct {
	*Container
	// roundPixels: boolean;
	RoundPixels bool `js:"roundPixels"`
	// setProperties(properties: ParticleContainerProperties): void;
	SetProperties func(*ParticleContainerProperties) `js:"setProperties"`
}

// [size=15000] {number} The number of images in the ParticleContainer before it flushes.
func NewParticleContainer(size int, ps ...*ParticleContainerProperties) *ParticleContainer {
	if len(ps) > 0 {
		return &ParticleContainer{
			Container: wrapContainer(pkg.Get("ParticleContainer").New(size, ps[0])),
		}
	}
	return &ParticleContainer{
		Container: wrapContainer(pkg.Get("ParticleContainer").New(size)),
	}
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
		Sprite: wrapSprite(pkg.Get("extras").Get("MovieClip").New(objs)),
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
