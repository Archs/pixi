package pixi

import (
	"github.com/gopherjs/gopherjs/js"
)

type Graphics struct {
	*Container

	// blendMode number
	//
	// The blend mode to be applied to the graphic shape. Apply a value of blendModes.NORMAL to reset the blend mode.
	// Default Value:
	// CONST.BLEND_MODES.NORMAL;
	BlendMode int `js:"blendMode"`

	// boundsPadding number
	// The bounds' padding used for bounds calculation.
	BoundsPadding float64 `js:"boundsPadding"`

	// fillAlpha number
	// The alpha value used when filling the Graphics object.
	// Default Value:
	// 1
	FillAlpha float64 `js:"fillAlpha"`

	// isMask boolean
	// Whether this shape is being used as a mask.
	IsMask bool `js:"isMask"`

	// lineColor string
	// The color of any lines drawn.
	// Default Value:
	// 0
	LineColor string `js:"lineColor"`

	// lineWidth number
	//
	// The width (thickness) of any lines drawn.
	// Default Value:
	// 0
	LineWidth float64 `js:"lineWidth"`

	// The tint applied to the sprite.
	// This is a hex value. A value of 0xFFFFFF will remove any tint effect.
	Tint uint32 `js:"tint"`
}

func NewGraphics() *Graphics {
	g := pkg.Get("Graphics").New()
	return &Graphics{
		Container: wrapContainer(g),
	}
}

// arc(cx, cy, radius, startAngle, endAngle, anticlockwise){Graphics}
// core/graphics/Graphics.js, line 457
// The arc method creates an arc/curve (used to create circles, or parts of circles).
//
//  Name    Type    Description
//  cx  number
//  The x-coordinate of the center of the circle
//  cy  number
//  The y-coordinate of the center of the circle
//  radius  number
//  The radius of the circle
//  startAngle  number
//  The starting angle, in radians (0 is at the 3 o'clock position of the arc's circle)
//  endAngle    number
//  The ending angle, in radians
//  anticlockwise   boolean
//  Optional. Specifies whether the drawing should be counterclockwise or clockwise. False is default, and indicates clockwise, while true indicates counter-clockwise.
// Returns:
//  Type    Description
//  Graphics
func (g *Graphics) Arc(cx, cy, radius, startAngle, endAngle float64, anticlockwise bool) *Graphics {
	g.Call("arc", cx, cy, radius, startAngle, endAngle, anticlockwise)
	return g
}

// arcTo(x1, y1, x2, y2, radius){Graphics}
//
// core/graphics/Graphics.js, line 390
// The arcTo() method creates an arc/curve between two tangents on the canvas.
//
//  Name    Type    Description
//  x1  number
//  The x-coordinate of the beginning of the arc
//  y1  number
//  The y-coordinate of the beginning of the arc
//  x2  number
//  The x-coordinate of the end of the arc
//  y2  number
//  The y-coordinate of the end of the arc
//  radius  number
//  The radius of the arc
//  Returns:
//  Type    Description
//  Graphics
func (g *Graphics) ArcTo(x1, y1, x2, y2, radius float64) *Graphics {
	g.Call("arcTo", x1, y1, x2, y2, radius)
	return g
}

// beginFill(color, alpha){Graphics}
// Specifies a simple one-color fill that subsequent calls to other Graphics methods (such as lineTo() or drawCircle()) use when drawing.
//
//  Name    Type    Description
//  color   number
//  the color of the fill
//  alpha   number
//  the alpha of the fill
//  Returns:
//  Type    Description
//  Graphics
func (g *Graphics) BeginFill(color, alpha float64) *Graphics {
	g.Call("beginFill", color, alpha)
	return g
}

// bezierCurveTo(cpX, cpY, cpX2, cpY2, toX, toY){Graphics}
//
// Calculate the points for a bezier curve and then draws it.
//
//  Name    Type    Description
//  cpX number
//  Control point x
//  cpY number
//  Control point y
//  cpX2    number
//  Second Control point x
//  cpY2    number
//  Second Control point y
//  toX number
//  Destination point x
//  toY number
//  Destination point y
//  Returns:
//  Type    Description
//  Graphics
func (g *Graphics) BezierCurveTo(cpX, cpY, cpX2, cpY2, toX, toY float64) *Graphics {
	g.Call("bezierCurveTo", cpX, cpY, cpX2, cpY2, toX, toY)
	return g
}

// clear(){Graphics}
//
// core/graphics/Graphics.js, line 676
// Clears the graphics that were drawn to this Graphics object, and resets fill and line style settings.
//
// Returns:
//
// Type    Description
// Graphics
func (g *Graphics) Clear() *Graphics {
	g.Call("clear")
	return g
}

// clone(){Graphics}
//
// core/graphics/Graphics.js, line 174
// Creates a new Graphics object with the same values as this one. Note that the only the properties of the object are cloned, not its transform (position,scale,etc)
//
// Returns:
//
// Type    Description
// Graphics
func (g *Graphics) Clone() *Graphics {
	o := g.Call("clone")
	return &Graphics{
		Container: wrapContainer(o),
	}
}

// containsPoint(point){boolean}
//
// core/graphics/Graphics.js, line 921
// Tests if a point is inside this graphics object
//
// Name    Type    Description
// point   Point
// the point to test
// Returns:
//
// Type    Description
// boolean the result of the test
func (g *Graphics) ContainsPoint(p Point) bool {
	o := g.Call("containsPoint", p.Object)
	return o.Bool()
}

// drawCircle(x, y, radius){Graphics}
//
// core/graphics/Graphics.js, line 619
// Draws a circle.
//
// Name    Type    Description
// x   number
// The X coordinate of the center of the circle
// y   number
// The Y coordinate of the center of the circle
// radius  number
// The radius of the circle
// Returns:
//
// Type    Description
// Graphics
func (g *Graphics) DrawCircle(x, y, radius float64) *Graphics {
	g.Call("drawCircle", x, y, radius)
	return g
}

// drawEllipse(x, y, width, height){Graphics}

// core/graphics/Graphics.js, line 635
// Draws an ellipse.

// Name    Type    Description
// x   number
// The X coordinate of the center of the ellipse
// y   number
// The Y coordinate of the center of the ellipse
// width   number
// The half width of the ellipse
// height  number
// The half height of the ellipse
// Returns:

// Type    Description
// Graphics
func (g *Graphics) DrawEllipse(x, y, width, height float64) *Graphics {
	g.Call("drawEllipse", x, y, width, height)
	return g
}

// drawPolygon(path){Graphics}

// core/graphics/Graphics.js, line 648
// Draws a polygon using the given path.

// Name    Type    Description
// path    Array
// The path data used to construct the polygon.
// Returns:

// Type    Description
// Graphics
func (g *Graphics) DrawPolygon(path float64) *Graphics {
	g.Call("drawPolygon", path)
	return g
}

// drawRect(x, y, width, height){Graphics}

// core/graphics/Graphics.js, line 589
// Name    Type    Description
// x   number
// The X coord of the top-left of the rectangle
// y   number
// The Y coord of the top-left of the rectangle
// width   number
// The width of the rectangle
// height  number
// The height of the rectangle
// Returns:

// Type    Description
// Graphics
func (g *Graphics) DrawRect(x, y, width, height float64) *Graphics {
	g.Call("drawRect", x, y, width, height)
	return g
}

// drawRoundedRect(x, y, width, height, radius)

// core/graphics/Graphics.js, line 604
// Name    Type    Description
// x   number
// The X coord of the top-left of the rectangle
// y   number
// The Y coord of the top-left of the rectangle
// width   number
// The width of the rectangle
// height  number
// The height of the rectangle
// radius  number
// Radius of the rectangle corners
func (g *Graphics) DrawRoundedRect(x, y, width, height, radius float64) *Graphics {
	g.Call("drawRoundedRect", x, y, width, height, radius)
	return g
}

// A GraphicsData  object.
type GraphicsData struct {
	*js.Object
	// Name    Type    Description
	// lineWidth   number
	// the width of the line to draw
	LineWidth float64 `js:"lineWidth"`
	// lineColor   number
	// the color of the line to draw
	LineColor float64 `js:"lineColor"`
	// lineAlpha   number
	// the alpha of the line to draw
	LineAlpha float64 `js:"lineAlpha"`
	// fillColor   number
	// the color of the fill
	FillColor float64 `js:"fillColor"`
	// fillAlpha   number
	// the alpha of the fill
	FillAlpha float64 `js:"fillAlpha"`
	// fill    boolean
	// whether or not the shape is filled with a colour
	Fill bool `js:"fill"`
	// shape   Circle | Rectangle | Ellipse | Line | Polygon
	// The shape object to draw.
	Shape *js.Object `js:"shape"`
}

// drawShape(shape){GraphicsData}
//
//  core/graphics/Graphics.js, line 1132
//  Draws the given shape to this Graphics object. Can be any of Circle, Rectangle, Ellipse, Line or Polygon.
//
//  Name    Type    Description
//  shape   Circle | Rectangle | Ellipse | Line | Polygon
//  The shape object to draw.
//  Returns:
//
//  Type    Description
//  GraphicsData    The generated GraphicsData object.
func (g *Graphics) DrawShape(shape Shape) *GraphicsData {
	o := g.Call("drawShape", shape.shape())
	return &GraphicsData{Object: o}
}

// endFill(){Graphics}
//
// core/graphics/Graphics.js, line 572
// Applies a fill to the lines and shapes that were added since the last call to the beginFill() method.
//
//  Returns:
//
//  Type    Description
//  Graphics
func (g *Graphics) EndFill() *Graphics {
	g.Call("endFill")
	return g
}

// generateTexture(resolution, scaleMode){Texture}
//
// core/graphics/Graphics.js, line 696
// Useful function that returns a texture of the graphics object that can then be used to create sprites This can be quite useful if your geometry is complicated and needs to be reused multiple times.
//
//  Name    Type    Description
//  resolution  number
//  The resolution of the texture being generated
//  scaleMode   number
//  Should be one of the scaleMode consts
//  Returns:
//
//  Type    Description
//  Texture a texture of the graphics object
func (g *Graphics) GenerateTexture(resolution float64, scaleMode int) *Texture {
	o := g.Call("generateTexture")
	return &Texture{Object: o}
}

// lineStyle(lineWidth, color, alpha){Graphics}
//
//  core/graphics/Graphics.js, line 211
//  Specifies the line style used for subsequent calls to Graphics methods such as the lineTo() method or the drawCircle() method.
//
//  Name    Type    Description
//  lineWidth   number
//  width of the line to draw, will update the objects stored style
//  color   number
//  color of the line to draw, will update the objects stored style
//  alpha   number
//  alpha of the line to draw, will update the objects stored style
//  Returns:
//
//  Type    Description
//  Graphics
func (g *Graphics) LineStyle(lineWidth, color, alpha float64) *Graphics {
	g.Call("lineStyle", lineWidth, color, alpha)
	return g
}

// lineTo(x, y){Graphics}
//
//  core/graphics/Graphics.js, line 258
//  Draws a line using the current line style from the current drawing position to (x, y); The current drawing position is then set to (x, y).
//
//  Name    Type    Description
//  x   number
//  the X coordinate to draw to
//  y   number
//  the Y coordinate to draw to
//  Returns:
//
//  Type    Description
//  Graphics
func (g *Graphics) LineTo(x, y float64) *Graphics {
	g.Call("lineTo", x, y)
	return g
}

// moveTo(x, y){Graphics}
//
//  core/graphics/Graphics.js, line 243
//  Moves the current drawing position to x, y.
//
//  Name    Type    Description
//  x   number
//  the X coordinate to move to
//  y   number
//  the Y coordinate to move to
//  Returns:
//
//  Type    Description
//  Graphics
func (g *Graphics) MoveTo(x, y float64) *Graphics {
	g.Call("moveTo", x, y)
	return g
}

// quadraticCurveTo(cpX, cpY, toX, toY){Graphics}
//
// core/graphics/Graphics.js, line 276
// Calculate the points for a quadratic bezier curve and then draws it. Based on: https://stackoverflow.com/questions/785097/how-do-i-implement-a-bezier-curve-in-c
//
//  Name    Type    Description
//  cpX number
//  Control point x
//  cpY number
//  Control point y
//  toX number
//  Destination point x
//  toY number
//  Destination point y
//  Returns:
//
//  Type    Description
//  Graphics
func (g *Graphics) QuadraticCurveTo(cpX, cpY, toX, toY float64) *Graphics {
	g.Call("quadraticCurveTo", cpX, cpY, toX, toY)
	return g
}

// updateLocalBounds()
//
// core/graphics/Graphics.js, line 953
// Update the bounds of the object
func (g *Graphics) UpdateLocalBounds() {
	g.Call("updateLocalBounds")
}
