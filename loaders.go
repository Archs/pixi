package pixi

import "github.com/gopherjs/gopherjs/js"

var (
	Loader = NewResourceLoader()
)

type ResourceLoader struct {
	*js.Object
}

func NewResourceLoader() *ResourceLoader {
	return &ResourceLoader{Object: pkg.Get("loaders").Get("Loader").New()}
}

// Adds a resource (or multiple resources) to the loader queue.
func (r *ResourceLoader) Add(name, url string, onComplete ...func()) *ResourceLoader {
	if len(onComplete) > 0 {
		r.Call("add", name, url, onComplete[0])
	} else {
		r.Call("add", name, url)
	}
	return r
}

// Starts loading the queued resources.
// Optinal callback that will be bound to the `complete` event
func (r *ResourceLoader) Load(callback ...func(resource *js.Object)) *ResourceLoader {
	if len(callback) > 0 {
		fn := func(loader, res *js.Object) {
			callback[0](res)
		}
		r.Call("load", fn)
	} else {
		r.Call("load")
	}
	return r
}
