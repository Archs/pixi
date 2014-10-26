package pixi

import "github.com/gopherjs/gopherjs/js"

type AssetLoader struct {
	js.Object
}

func NewAssetLoader(urls []string, crossOrigin bool) *AssetLoader {
	return &AssetLoader{Object: pkg.Get("AssetLoader").New(urls, crossOrigin)}
}

func (a *AssetLoader) OnComplete(cb func()) {
	a.Set("onComplete", cb)
}

func (a *AssetLoader) OnProgress(cb func()) {
	a.Set("onProgress", cb)
}

func (a *AssetLoader) Load() {
	a.Call("load")
}
