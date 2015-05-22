package pixi

import "github.com/gopherjs/gopherjs/js"

type Renderer struct {
	*js.Object
	View *js.Object `js:"view"`
	// 	width	number	800	optional
	// the width of the canvas view
	Width float64 `js:"width"`
	// height	number	600	optional
	// the height of the canvas view
	Height float64 `js:"height"`

	// autoResize boolean
	// Whether the render view should be resized automatically
	AutoResize bool `js:"autoResize"`

	// blendModes object.<string, mixed>
	// Tracks the blend modes useful for this renderer.
	BlendModes *js.Object `js:"blendModes"`

	// clearBeforeRender boolean
	// This sets if the CanvasRenderer will clear the canvas or not before the new render pass. If the scene is NOT transparent Pixi will use a canvas sized fillRect operation every frame to set the canvas background color. If the scene is transparent Pixi will use clearRect to clear the canvas every frame. Disable this by setting this to false. For example if your game has a canvas filling background image you often don't need this set.
	ClearBeforeRender bool `js:"clearBeforeRender"`

	// preserveDrawingBuffer boolean
	// The value of the preserveDrawingBuffer flag affects whether or not the contents of the stencil buffer is retained after rendering.
	PreserveDrawingBuffer bool `js:"preserveDrawingBuffer"`

	// resolution number
	// The resolution of the renderer
	// Default Value:
	// 1
	Resolution int `js:"resolution"`

	// transparent boolean
	// Whether the render view is transparent, read only
	Transparent bool `js:"transparent"`
	// background color, settable
	BackgroundColor float64 `js:"backgroundColor"`

	// type RENDERER_TYPE
	// The returntype of the renderer.
	// Default Value:
	// CONT.RENDERER_TYPE.UNKNOWN
	Type bool `js:"type"`
}

func (r Renderer) Render(c *Container) {
	r.Call("render", c.Object)
}

// resize(width, height)
//
// Resizes the canvas view to the specified width and height
//	 Name	Type	Description
//	 width	number
//	 the new width of the canvas view
//	 height	number
//	 the new height of the canvas view
func (r *Renderer) Resize(width, height float64) {
	r.Call("resize", width, height)
}

// destroy(removeView)
//
// Removes everything from the renderer and optionally removes the Canvas DOM element.
//	 Name	Type	Default	Description
//	 removeView	boolean	false	optional
//	 Removes the Canvas element from the DOM.
func (r *Renderer) Destroy(removeView bool) {
	r.Call("destroy", removeView)
}

// export interface RendererOptions {
//     view?: HTMLCanvasElement;
//     transparent?: boolean
//     antialias?: boolean;
//     resolution?: number;
//     clearBeforeRendering?: boolean;
//     preserveDrawingBuffer?: boolean;
//     forceFXAA?: boolean;
// }
type RendererOptions struct {
	*js.Object
	// transparent boolean
	// Whether the render view is transparent
	Transparent bool `js:"transparent"`
	// antialias?: boolean;
	Antialias bool `js:"antialias"`
	// resolution?: number;
	Resolution int `js:"resolution"`
	// clearBeforeRendering?: boolean;
	ClearBeforeRendering bool `js:"clearBeforeRendering"`
	// preserveDrawingBuffer?: boolean;
	PreserveDrawingBuffer bool `js:"preserveDrawingBuffer"`
	// forceFXAA?: boolean;
	ForceFXAA bool `js:"forceFXAA"`
	// background color
	BackgroundColor float64 `js:"backgroundColor"`
}

func NewRendererOptions() *RendererOptions {
	return &RendererOptions{
		Object: js.Global.Get("Object").New(),
	}
}

func AutoDetectRenderer(width, height int, options ...*RendererOptions) *Renderer {
	if len(options) > 0 {
		return &Renderer{Object: pkg.Call("autoDetectRenderer", width, height, options[0])}
	}
	return &Renderer{Object: pkg.Call("autoDetectRenderer", width, height)}
}
