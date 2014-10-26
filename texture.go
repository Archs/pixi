package pixi

import "github.com/gopherjs/gopherjs/js"

type Texture struct {
	js.Object
}

func TextureFromImage(url string, crossOrigin bool, scaleMode int) *Texture {
	return &Texture{Object: pkg.Get("Texture").Call("fromImage", url, crossOrigin, scaleMode)}
}

func TextureFromFrame(frameId string) *Texture {
	return &Texture{Object: pkg.Get("Texture").Call("fromFrame", frameId)}
}
