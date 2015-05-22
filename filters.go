package pixi

import (
	"github.com/gopherjs/gopherjs/js"
)

type filter interface {
	filter() *js.Object
}

// all filters implements the filter interface
//
// IMPORTANT: This is a webGL only feature and will be ignored by the canvas renderer.
// To remove filters simply set this property to 'null'
func (d *DisplayObject) AddFilter(f filter) {
	if d.filters == nil {
		d.filters = []*js.Object{}
	}
	d.filters = append(d.filters, f.filter())
}

func (d *DisplayObject) RemoveFilter(f filter) {
	target := f.filter()
	for i, v := range d.filters {
		if v == target {
			d.filters = append(d.filters[:i], d.filters[i+1:]...)
			return
		}
	}
}

func (d *DisplayObject) RemoveAllFilters() {
	d.filters = nil
}

// PIXI. AbstractFilter
// core/renderers/webgl/filters/AbstractFilter.js, line 13
type AbstractFilter struct {
	*js.Object
	// Members

	// padding number
	// The extra padding that the filter might need
	Padding float64 `js:"padding"`

	// uniforms object
	// The uniforms as an object
	Uniforms *js.Object `js:"uniforms"`
}

func (a *AbstractFilter) filter() *js.Object {
	return a.Object
}

func wrapAbstractFilter(o *js.Object) *AbstractFilter {
	return &AbstractFilter{
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
func NewAbstractFilter(vertexSrc, fragmentSrc, uniforms interface{}) *AbstractFilter {
	return wrapAbstractFilter(pkg.Get("AbstractFilter").New(vertexSrc, fragmentSrc, uniforms))
}

// Methods
// syncUniform()
// Syncs a uniform between the class object and the shaders.
func (f *AbstractFilter) SyncUniform() {
	f.Call("syncUniform")
}

type InvertFilter struct {
	*AbstractFilter
	// invert number [0, 1]
	Invert float64 `js:"invert"`
}

// new PIXI.filters.InvertFilter()
// filters/invert/InvertFilter.js, line 12
// This inverts your Display Objects colors.
func NewInvertFilter() *InvertFilter {
	return &InvertFilter{
		AbstractFilter: wrapAbstractFilter(pkg.Get("filters").Call("InvertFilter")),
	}
}

// new PIXI.filters.BloomFilter()
//
// filters/bloom/BloomFilter.js, line 13
// The BloomFilter applies a Gaussian blur to an object. The strength of the blur can be set for x- and y-axis separately.
func NewBloomFilter() *AbstractFilter {
	return wrapAbstractFilter(pkg.Get("filters").Call("BloomFilter"))
}

// new PIXI.filters.BlurDirFilter(dirX, dirY)
// // filters/blur/BlurDirFilter.js, line 13
// The BlurDirFilter applies a Gaussian blur toward a direction to an object.
//
// Name	Type	Description
// dirX	number
// dirY	number
func NewBlurDirFilter(dirX, dirY float64) *AbstractFilter {
	return wrapAbstractFilter(pkg.Get("filters").Call("BlurDirFilter", dirX, dirY))
}
