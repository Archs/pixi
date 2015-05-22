package pixi

import (
	"github.com/Archs/js/dom"

	"github.com/gopherjs/gopherjs/js"
)

type Texture struct {
	*js.Object
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
