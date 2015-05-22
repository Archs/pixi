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
	TRIANGLE_MESH float64 `js:"TRIANGLE_MESH"`
	// TRIANGLES	number
	TRIANGLES float64 `js:"TRIANGLES"`
}

// Base mesh class
// PIXI.Container
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
}

// new PIXI.mesh.Mesh(texture, vertices, uvs, indices, drawMode)

// Name	Type	Description
// texture	Texture
// The texture to use
// vertices		optional
// {Float32Arrif you want to specify the vertices
// uvs	Float32Array	optional
// if you want to specify the uvs
// indices	Uint16Array	optional
// if you want to specify the indices
// drawMode	number	optional
// the drawMode, can be any of the Mesh.DRAW_MODES consts
// Extends

type Rope struct {
}
