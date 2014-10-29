# GopherJS bindings to Pixi.js

## Install

We use gopkg.in for imports in order to provide different versions of the library that you can count on not to change. The .v## at the end of the import specifies the Pixi.js major version. For example, v20 will wraps version 2.0.x, v21 would wrap version 2.1.x.

```bash
go get gopkg.in/ajhager/pixi.v20
```

## Example

You can find the Pixi.js examples ported to Go in the `examples` directory. Below is the absolute bare minimum needed to get started.

```go
import (
	"gopkg.in/ajhager/pixi.v20"
	"github.com/ajhager/raf"

	"github.com/gopherjs/gopherjs/js"
)

var (
	stage = pixi.NewStage(0x000000)
	renderer = pixi.AutoDetectRenderer(800, 600)
)

func animate(t float32) {
	raf.RequestAnimationFrame(animate)
	renderer.Render(stage)
}

func main() {
	js.Global.Get("document").Get("body").Call("appendChild", renderer.View)
	raf.RequestAnimationFrame(animate)
}
```
