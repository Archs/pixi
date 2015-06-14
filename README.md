# GopherJS bindings to Pixi.js

## Install

We use gopkg.in for imports in order to provide different versions of the library that you can count on not to change. The .v### at the end of the import specifies the Pixi.js major version. For example, v2.0 wraps version 2.0.x and v2.1 would wrap version 2.1.x.

Currently v2.x.x and v3.x.x are supported.

```bash
go get gopkg.in/ajhager/pixi.v3
```

## Example

You can find the Pixi.js examples ported to Go in the `examples` directory. Below is the absolute bare minimum needed to get started.

```go
package main

import (
	"github.com/Archs/js/dom"
	"github.com/Archs/js/raf"
	"gopkg.in/ajhager/pixi.v3"
)

var (
	stage    = pixi.NewContainer()
	renderer = pixi.AutoDetectRenderer(800, 600, 0x00000)
)

func animate(t float64) {
	raf.RequestAnimationFrame(animate)
	renderer.Render(stage)
}

func main() {
	dom.OnDOMContentLoaded(func() {
		el := dom.Wrap(renderer.View)
		dom.Body().AppendChild(el)
		raf.RequestAnimationFrame(animate)
	})
}
```
