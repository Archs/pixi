package pixi

import (
	"github.com/gopherjs/gopherjs/js"
)

// type filter interface {
// 	filter() *js.Object
// }

// func (d *DisplayObject) AddFilter(f filter) {
// 	if d.filters == nil {
// 		d.filters = []filter{}
// 	}
// 	d.filters = append(d.filters, f.filter())
// }

// func (d *DisplayObject) RemoveFilter(f filter) {
// 	target := f.filter()
// 	for i, v := range d.filters {
// 		if v == target {
// 			d.filters = append(d.filters[:i], d.filters[i+1:]...)
// 			return
// 		}
// 	}
// }

// func (d *DisplayObject) RemoveAllFilters() {
// 	d.filters = nil
// }

// PIXI. AbstractFilter
// core/renderers/webgl/filters/AbstractFilter.js, line 13
type Filter struct {
	*js.Object
	// Members

	// padding number
	// The extra padding that the filter might need
	Padding float64 `js:"padding"`

	// uniforms object
	// The uniforms as an object
	Uniforms *js.Object `js:"uniforms"`
}

func wrapFilter(o *js.Object) *Filter {
	return &Filter{
		Object: o,
	}
}

// new PIXI.AbstractFilter(vertexSrc, fragmentSrc, uniforms)
// This is the base class for creating a PIXI filter.
// Currently only WebGL supports filters.
// If you want to make a custom filter this should be your base class.
//	Name	Type	Description
//	vertexSrc	string | Array.<string>
//	The vertex shader source as an array of strings.
//	fragmentSrc	string | Array.<string>
//	The fragment shader source as an array of strings.
//	uniforms	object
//	An object containing the uniforms for this filter.
func NewFilter(vertexSrc, fragmentSrc, uniforms interface{}) *Filter {
	return wrapFilter(pkg.Get("AbstractFilter").New(vertexSrc, fragmentSrc, uniforms))
}

// Methods
// syncUniform()
// Syncs a uniform between the class object and the shaders.
func (f *Filter) SyncUniform() {
	f.Call("syncUniform")
}

// new PIXI.filters.InvertFilter()
// filters/invert/InvertFilter.js, line 12
// This inverts your Display Objects colors.
func NewInvertFilter() *Filter {
	return wrapFilter(pkg.Get("filters").Call("InvertFilter"))
}
