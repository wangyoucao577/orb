package tilecover

import (
	"github.com/paulmach/orb/geo"
	"github.com/paulmach/orb/tile"
)

// ForBound creates a tile cover for the bound. i.e. all the tiles
// that intersect the bound.
func ForBound(b geo.Bound, z tile.Zoom) tile.Tiles {
	lo := tile.At(b[0], z)
	hi := tile.At(b[1], z)

	result := make(tile.Tiles, 0, (hi.X-lo.X+1)*(lo.Y-hi.Y+1))

	for x := lo.X; x <= hi.X; x++ {
		for y := hi.Y; y <= lo.Y; y++ {
			result = append(result, tile.Tile{X: x, Y: y, Z: z})
		}
	}

	return result
}

// ForPoint creates a tile cover for the point, i.e. just the tile
// containing the point.
func ForPoint(ll geo.Point, z tile.Zoom) tile.Tiles {
	return tile.Tiles{tile.At(ll, z)}
}
