package pixi

import (
	"github.com/gopherjs/gopherjs/js"
)

// A Text Object will create a line or multiple lines of text. To split a line you can use '\n' in your text string, or add a wordWrap property set to true and and wordWrapWidth property with a value in the style object.
//
// A Text can be created directly from a string and a style object
//
type Text struct {
	*Sprite
}

func wrapText(o *js.Object) *Text {
	return &Text{
		Sprite: wrapSprite(o),
	}
}

// var text = new PIXI.Text('This is a pixi text',{font : '24px Arial', fill : 0xff1010, align : 'center'});
//
// text	string
//  he copy that you would like the text to display
//
// style	object	optional
//
//  The style parameters
//  Name	Type	Default	Description
//  font	string		optional
//  	default 'bold 20px Arial' The style and size of the font
//  fill	String | Number	'black'	optional
//  	A canvas fillstyle that will be used on the text e.g 'red', '#00FF00'
//  align	string	'left'	optional
//  	Alignment for multiline text ('left', 'center' or 'right'), does not affect single line text
//  stroke	String | Number		optional
//  	A canvas fillstyle that will be used on the text stroke e.g 'blue', '#FCFF00'
//  strokeThickness	number	0	optional
//  	A number that represents the thickness of the stroke. Default is 0 (no stroke)
//  wordWrap	boolean	false	optional
//  	Indicates if word wrap should be used
//  wordWrapWidth	number	100	optional
//  	The width at which text will wrap, it needs wordWrap to be set to true
//  lineHeight	number		optional
//  	The line height, a number that represents the vertical space that a letter uses
//  dropShadow	boolean	false	optional
//  	Set a drop shadow for the text
//  dropShadowColor	string	'#000000'	optional
//  	A fill style to be used on the dropshadow e.g 'red', '#00FF00'
//  dropShadowAngle	number	Math.PI/4	optional
//  	Set a angle of the drop shadow
//  dropShadowDistance	number	5	optional
//  	Set a distance of the drop shadow
//  padding	number	0	optional
//  	Occasionally some fonts are cropped on top or bottom. Adding some padding will prevent this from happening by adding padding to the top and bottom of text height.
//  textBaseline	string	'alphabetic'	optional
//  	The baseline of the text that is rendered.
//  lineJoin	string	'miter'	optional
//  	The lineJoin property sets the type of corner created, it can resolve spiked text issues. Default is 'miter' (creates a sharp corner).
//  miterLimit	number	10	optional
//  	The miter limit to use when using the 'miter' lineJoin mode. This can reduce or increase the spikiness of rendered text.
func NewText(text string, style ...interface{}) *Text {
	args := []interface{}{text}
	args = append(args, style...)
	o := pkg.Call("Text", args...)
	return wrapText(o)
}
