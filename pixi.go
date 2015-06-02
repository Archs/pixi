package pixi

import (
	_ "github.com/Archs/pixi/shim" // Shim library javascript code.
	"github.com/gopherjs/gopherjs/js"
)

var pkg = js.Global.Get("PIXI")

var (
	WEBGL_RENDERER  = pkg.Get("WEBGL_RENDERER").Int()
	CANVAS_RENDERER = pkg.Get("WEBGL_RENDERER").Int()

	VERSION = pkg.Get("VERSION").String()

	ScaleModes   = scaleModes{Object: pkg.Get("SCALE_MODES")}
	BlendModes   = blendModes{Object: pkg.Get("BLEND_MODES")}
	RENDERERTYPE = rendererType{Object: pkg.Get("RENDERER_TYPE")}
)

// PIXI.RENDERER_TYPE
//
// Constant to identify the Renderer Type.
// RENDERER_TYPE	object
// Properties
//
// Name	Type	Description
// UNKNOWN	number
// WEBGL	number
// CANVAS	number

type rendererType struct {
	*js.Object
	// UNKNOWN	number
	UNKNOWN int `js:"UNKNOWN"`
	// WEBGL	number
	WEBGL int `js:"WEBGL"`
	// CANVAS	number
	CANVAS int `js:"CANVAS"`
}

type blendModes struct {
	*js.Object
	Normal     int `js:"NORMAL"`
	Add        int `js:"ADD"`
	Multiply   int `js:"MULTIPLY"`
	Screen     int `js:"SCREEN"`
	Overlay    int `js:"OVERLAY"`
	Darken     int `js:"DARKEN"`
	Lighten    int `js:"LIGHTEN"`
	ColorDodge int `js:"COLOR_DODGE"`
	ColorBurn  int `js:"COLOR_BURN"`
	HardLight  int `js:"HARD_LIGHT"`
	SoftLight  int `js:"SOFT_LIGHT"`
	Difference int `js:"DIFFERENCE"`
	Exclusion  int `js:"EXCLUSION"`
	Hue        int `js:"HUE"`
	Saturation int `js:"SATURATION"`
	Color      int `js:"COLOR"`
	Luminosity int `js:"LUMINOSITY"`
}

type scaleModes struct {
	*js.Object
	Default int `js:"DEFAULT"`
	Linear  int `js:"LINEAR"`
	Nearest int `js:"NEAREST"`
}

type InteractionData struct {
	*js.Object

	// This point stores the global coords of where the touch/mouse event happened
	Global *Point `js:"global"`
	// The target Sprite that was interacted with
	Target *Sprite `js:"target"`
	// When passed to an event handler, this will be the original DOM Event that was captured
	OriginalEvent *js.Object `js:"originalEvent"`
}

func wrapInteractionData(object *js.Object) *InteractionData {
	return &InteractionData{
		Object: object,
	}
}

func (id *InteractionData) LocalPosition(do displayObject) *Point {
	return &Point{Object: id.Object.Call("getLocalPosition", do.displayer())}
}

type InteractionEvent struct {
	*js.Object
	// stopped: false,
	Stopped bool `js:"stopped"`
	// target: null,
	Target *DisplayObject `js:"target"`
	// type: null,
	Type string `js:"type"`
	// data: this.mouse,
	Data *InteractionData `js:"data"`
	// stopPropagation:function(){
	//     this.stopped = true;
	// }
	StopPropagation func() `js:"stopPropagation"`
}
