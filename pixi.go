package pixi

import (
	_ "github.com/ajhager/pixi/shim" // Shim library javascript code.
	"github.com/gopherjs/gopherjs/js"
)

var pkg = js.Global.Get("PIXI")

var (
	WEBGL_RENDERER  = pkg.Get("WEBGL_RENDERER").Int()
	CANVAS_RENDERER = pkg.Get("WEBGL_RENDERER").Int()

	VERSION = pkg.Get("VERSION").Str()
)

var BlendModes = blendModes{Object: pkg.Get("blendModes")}

type blendModes struct {
	js.Object
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

var ScaleModes = scaleModes{Object: pkg.Get("scaleModes")}

type scaleModes struct {
	js.Object
	Default int `js:"DEFAULT"`
	Linear  int `js:"LINEAR"`
	Nearest int `js:"NEAREST"`
}
