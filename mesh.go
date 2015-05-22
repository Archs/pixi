package pixi

import (
	"github.com/gopherjs/gopherjs/js"
)

var (
	DRAWMODES = drawMODES{Object: pkg.Get("mesh").Get("Mesh").Get("DRAW_MODES")}
)

// PIXI.mesh.Mesh.DRAW_MODES
type drawMODES struct {
	*js.Object
	// Members
	// static,constantPIXI.mesh.Mesh.DRAW_MODES
	//
	// Different drawing buffer modes supported
	//
	// Properties:
	//
	// Name	Type	Description
	// DRAW_MODES	object
	// Properties
	//
	// Name	Type	Description
	// TRIANGLE_MESH	number
	TRIANGLE_MESH int `js:"TRIANGLE_MESH"`
	// TRIANGLES	number
	TRIANGLES int `js:"TRIANGLES"`
}

// Base mesh class Extends PIXI.Container
type Mesh struct {
	*Container
	// 	uvs Float32Array
	// The Uvs of the Mesh
	Uvs []float64 `js:"uvs"`

	// vertices Float32Array
	// An array of vertices
	Vertices []float64 `js:"vertices"`

	// canvasPaddingnumber
	// Triangles in canvas mode are automatically antialiased, use this value to force triangles to overlap a bit with each other.
	CanvasPaddingnumber float64 `js:"canvasPaddingnumber"`

	// dirty boolean
	// Whether the Mesh is dirty or not
	Dirty bool `js:"dirty"`

	// 	drawMode number
	// The way the Mesh should be drawn, can be any of the Mesh.DRAW_MODES consts
	DrawMode float64 `js:"drawMode"`
	// 	texturePIXI.Texture
	// The texture that the sprite is using
	Texture *Texture `js:"texture"`
}

func wrapMesh(o *js.Object) *Mesh {
	return &Mesh{
		Container: wrapContainer(o),
	}
}

// new PIXI.mesh.Mesh(texture, vertices, uvs, indices, drawMode)
//
//  Name	Type	Description
//  texture	Texture
//  The texture to use
//  vertices		optional
//  Float32Arr if you want to specify the vertices
//  uvs	Float32Array	optional
//  if you want to specify the uvs
//  indices	Uint16Array	optional
//  if you want to specify the indices
//  drawMode	number	optional
//  the drawMode, can be any of the Mesh.DRAW_MODES consts
func NewMesh(texture *Texture, vertices, uvs, indices []float64, drawMode int) *Mesh {
	return wrapMesh(pkg.Get("mesh").Get("Mesh").New(texture, vertices, uvs, indices, drawMode))
}

type Rope struct {
	*Mesh
	Colors []float64 `js:"colors"`
}

// new PIXI.mesh.Rope(texture, points)
//
// mesh/Rope.js, line 21
// The rope allows you to draw a texture across several points
// and then manipulate these points
//
// for (var i = 0; i < 20; i++) {
//     points.push(new PIXI.Point(i * 50, 0));
// };
// var rope = new PIXI.Rope(PIXI.Texture.fromImage("snake.png"), points);
// Name	Type	Description
// texture	Texture
// The texture to use on the rope.
// points	Array
// An array of {Point} objects to construct this rope.
func NewRope(texture *Texture, points []*Point) *Rope {
	o := pkg.Get("mesh").Get("Rope").New(texture, points)
	return &Rope{Mesh: wrapMesh(o)}
}

// refresh()
//
// Refreshes
func (r *Rope) Refreshe() {
	r.Call("refresh")
}
