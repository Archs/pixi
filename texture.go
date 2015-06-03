package pixi

import (
	"github.com/Archs/js/dom"

	"github.com/gopherjs/gopherjs/js"
)

type Texture struct {
	*js.Object
	// for base texture

	Uuid       float64 `js:"uuid"`
	Resolution float64 `js:"resolution"`
	// The width of the base texture set when the image has loaded
	Width float64 `js:"width"`
	// The height of the base texture set when the image has loaded
	Height float64 `js:"height"`
	// Used to store the actual width of the source of this texture
	RealWidth float64 `js:"realWidth"`
	// Used to store the actual height of the source of this texture
	RealHeight float64 `js:"realHeight"`
	// The scale mode to apply when scaling this texture
	ScaleMode float64 `js:"scaleMode"`
	// This is never true if the underlying source fails to load or has no texture data.
	HasLoaded bool `js:"hasLoaded"`
	// * Set to true if the source is currently loading.
	// *
	// * If an Image source is loading the 'loaded' or 'error' event will be
	// * dispatched when the operation ends. An underyling source that is
	// * immediately-available bypasses loading entirely.
	IsLoading bool `js:"isLoading"`
	// The image source that is used to create the texture. read only
	Source *js.Object `js:"source"` // HTMLImageElement | HTMLCanvasElement;
	// Controls if RGB channels should be pre-multiplied by Alpha  (WebGL only)
	PremultipliedAlpha bool   `js:"premultipliedAlpha"`
	ImageUrl           string `js:"imageUrl"`
	// Wether or not the texture is a power of two, try to use power of two textures as much as you can
	IsPowerOfTwo bool `js:"isPowerOfTwo"`
	//  Set this to true if a mipmap of this texture needs to be generated. This value needs to be set before the texture is used
	// * Also the texture must be a power of two size to work
	Mipmap bool `js:"mipmap"`

	// update(): void;
	//  Updates the texture on all the webgl renderers, this also assumes the src has changed.
	Update func() `js:"update"`
	// loadSource(source: HTMLImageElement | HTMLCanvasElement): void;
	// LoadSource func() `js:"loadSource"`

	// destroy(): void;
	// Destroys this base texture
	Destroy func(destroyBase bool) `js:"destroy"`
	// dispose(): void;
	// * Frees the texture from WebGL memory without destroying this texture object.
	// * This means you can still use the texture later which will upload it to GPU
	// * memory again.
	Dispose func() `js:"dispose"`
	// updateSourceImage(newSrc: string): void;
	// * Changes the source image of the texture.
	// * The original source must be an Image element.
	UpdateSourceImage func() `js:"updateSourceImage"`

	// for texture

	// Does this Texture have any frame data assigned to it?
	NoFrame bool `js"noFrame"`
	// The texture trim data.
	Trim *Rectangle `js"trim"`
	// This will let the renderer know if the texture is valid. If it's not then it cannot be rendered.
	Valid bool `js"valid"`
	// This will let a renderer know that a texture has been updated (used mainly for webGL uv updates)
	RequiresUpdate bool `js"requiresUpdate"`
	// This is the area of the BaseTexture image to actually copy to the Canvas / WebGL when rendering, irrespective of the actual frame size or placement (which can be influenced by trimmed texture atlases)
	// The area of original texture
	Crop *Rectangle `js"crop"`
	// Indicates whether the texture should be rotated by 90 degrees
	Rotate bool `js"rotate"`
	// The rectangle frame of the texture to show
	Frame *Rectangle      `js"frame"`
	Clone func() *Texture `js"clone"`
	// BaseTexture    BaseTexture     `js"baseTexture"`
}

type RenderTexture struct {
	*Texture
	Renderer *Renderer `js"renderer"` // CanvasRenderer | WebGLRenderer;

	// render(displayObject *DisplayObject, matrix Matrix, clear bool, updateTransform bool)
	Resize    func(width float64, height float64, updateBase bool) `js:"resize"`
	Clear     func()                                               `js:"clear"`
	GetImage  func() *js.Object                                    `js:"getImage"`
	GetPixels func() []float64                                     `js:"getPixels"`
	GetPixel  func(x float64, y float64) []float64                 `js:"getPixel"`
	GetBase64 func() string                                        `js:"getBase64"`
	GetCanvas func() *js.Object                                    `js:"getCanvas"`
}

// func (r RenderTexture) Clone() *RenderTexture {
// 	return &RenderTexture{
// 		Texture: *wrapTexture(r.Call("clone")),
// 	}
// }

// render(displayObject, matrix, clear, updateTransform)
// Draw/render the given DisplayObject onto the texture.
//
// The displayObject and descendents are transformed during this operation. If updateTransform is true then the transformations will be restored before the method returns. Otherwise it is up to the calling code to correctly use or reset the transformed display objects.
//
// The display object is always rendered with a worldAlpha value of 1.
//
//   Name    Type    Default Description
//   displayObject   DisplayObject
//   The display object to render this texture on
//   matrix  Matrix      optional
//   Optional matrix to apply to the display object before rendering.
//   clear   boolean false   optional
//   If true the texture will be cleared before the displayObject is drawn
//   updateTransform boolean true    optional
//   If true the displayObject's worldTransform/worldAlpha and all children transformations will be restored. Not restoring this information will be a little faster.
func (r *RenderTexture) Render(displayObject displayObject, clear bool, updateTransform bool) {
	r.Call("render", displayObject.displayer(), nil, clear, updateTransform)
}

func wrapTexture(o *js.Object) *Texture {
	return &Texture{Object: o}
}

func TextureFromImage(url string, crossOrigin bool, scaleMode int) *Texture {
	return &Texture{Object: pkg.Get("Texture").Call("fromImage", url, crossOrigin, scaleMode)}
}

func TextureFromFrame(frameId string) *Texture {
	return &Texture{Object: pkg.Get("Texture").Call("fromFrame", frameId)}
}

// PIXI.Texture.fromVideo(video, scaleMode){Texture}
//
// core/textures/Texture.js, line 350
// Helper function that creates a new Texture based on the given video element.
//
// 	Name	Type	Description
// 	video	HTMLVideoElement
// 	scaleMode	number
// 	See {{#crossLink "PIXI/scaleModes:property"}}scaleModes{{/crossLink}} for possible values
// 	Returns:
//
// 	Type	Description
// 	Texture	A Texture
func TextureFromVideo(video *dom.Element, scaleMode int) *Texture {
	return &Texture{Object: pkg.Get("Texture").Call("fromVideo", video.Object, scaleMode)}
}

// PIXI.Texture.fromVideoUrl(videoUrl, scaleMode){Texture}
//
// core/textures/Texture.js, line 370
// Helper function that creates a new Texture based on the video url.
//
//  Name	Type	Description
//  videoUrl	string
//  scaleMode	number
//  See {SCALE_MODES} for possible values
//  Returns:
//
//  Type	Description
//  Texture	A Texture
func TextureFromVideoUrl(videoUrl string, scaleMode int) *Texture {
	return &Texture{Object: pkg.Get("Texture").Call("fromVideoUrl", videoUrl, scaleMode)}
}

// PIXI.Texture.fromCanvas(canvas, scaleMode){Texture}
//
// core/textures/Texture.js, line 337
// Helper function that creates a new Texture based on the given canvas element.
//
//  Name	Type	Description
//  canvas	Canvas
//  The canvas element source of the texture
//  scaleMode	number
//  See {{#crossLink "PIXI/scaleModes:property"}}scaleModes{{/crossLink}} for possible values
//  Returns:
//
//  Type	Description
//  Texture
func TextureFromCanvas(canvas *js.Object, scaleMode int) *Texture {
	return &Texture{Object: pkg.Get("Texture").Call("fromCanvas", canvas, scaleMode)}
}

// new PIXI.RenderTexture(renderer, width, height, scaleMode, resolution)
//
// A RenderTexture is a special texture that allows any Pixi display object to be rendered to it.
//
// Hint: All DisplayObjects (i.e. Sprites) that render to a RenderTexture should be preloaded otherwise black rectangles will be drawn instead.
//
//      A RenderTexture takes a snapshot of any Display Object given to its render method. The position and rotation of the given Display Objects is ignored. For example:
//
//      var renderTexture = new PIXI.RenderTexture(800, 600);
//      var sprite = PIXI.Sprite.fromImage("spinObj_01.png");
//
//      sprite.position.x = 800/2;
//      sprite.position.y = 600/2;
//      sprite.anchor.x = 0.5;
//      sprite.anchor.y = 0.5;
//
//      renderTexture.render(sprite);
//      //The Sprite in this case will be rendered to a position of 0,0. To render this sprite at its actual position a Container should be used:
//
//      var doc = new Container();
//
//      doc.addChild(sprite);
//
// args
//
//      renderTexture.render(doc);  // Renders to center of renderTexture
//      Name    Type    Default Description
//      renderer    CanvasRenderer | WebGLRenderer
//      The renderer used for this RenderTexture
//      width   number  100 optional
//      The width of the render texture
//      height  number  100 optional
//      The height of the render texture
//      scaleMode   number      optional
//      See SCALE_MODES for possible values
//      resolution  number  1   optional
//      The resolution of the texture being generated
func NewRenderTexture(renderer *Renderer, width, height float64, scaleMode, resolution int) *RenderTexture {
	return &RenderTexture{
		Texture: wrapTexture(pkg.Get("RenderTexture").New(renderer, width, height, scaleMode, resolution)),
	}
}
