package pixi

// PIXI.extras.TilingSprite Extends PIXI.Sprite
//
// A tiling sprite is a fast way of rendering a tiling image
type TilingSprite struct {
	*Sprite
	// Name	Type	Description
	// texture	Texture
	// the texture of the tiling sprite
	Texture *Texture `js:"texture"`
	// width	number
	// the width of the tiling sprite
	Width float64 `js:"width"`
	// height	number
	// the height of the tiling sprite
	Height float64 `js:"height"`
	// tiling control variables
	TilePosition Point `js:"tilePosition"`
	TileScale    Point `js:"tileScale"`
}

// new PIXI.extras.TilingSprite(texture, width, height)
func NewTilingSprite(texture *Texture, width, height float64) *TilingSprite {
	return &TilingSprite{
		Sprite: wrapSprite(pkg.Get("TilingSprite").New(texture, width, height)),
	}
}

// PIXI.extras.TilingSprite.fromImage(imageId, width, height, crossorigin, scaleMode){TilingSprite}
//
// Helper function that creates a sprite that will contain a texture based on an image url If the image is not in the texture cache it will be loaded
//
//  Name	Type	Default	Description
//  imageId	String
//  The image url of the texture
//  width	number
//  the width of the tiling sprite
//  height	number
//  the height of the tiling sprite
//  crossorigin	boolean	(auto)	optional
//  if you want to specify the cross-origin parameter
//  scaleMode	number	scaleModes.DEFAULT	optional
//  if you want to specify the scale mode, see SCALE_MODES for possible values
func TilingTextureFromImage(url string, width, height float64, crossOrigin bool, scaleMode int) *Texture {
	return &Texture{Object: pkg.Get("TilingSprite").Call("fromImage", url, width, height, crossOrigin, scaleMode)}
}

// PIXI.extras.TilingSprite.fromFrame(frameId, width, height){TilingSprite}
//
// Helper function that creates a tiling sprite that will use a texture from the TextureCache based on the frameId The frame ids are created when a Texture packer file has been loaded
//
//  Name	Type	Description
//  frameId	String
//  The frame Id of the texture in the cache
//  width	number
//  the width of the tiling sprite
//  height	number
//  the height of the tiling sprite
//  Returns:
//
//  Type	Description
//  TilingSprite	A new TilingSprite using a texture from the texture cache matching the frameId
func TilingTextureFromFrame(frameId string, width, height float64) *Texture {
	return &Texture{Object: pkg.Get("TilingSprite").Call("fromFrame", frameId, width, height)}
}
