package pixelset

import (
	"math"

	iiifimageapi "github.com/dpb587/go-iiif-image-api-v3"
)

type ImageTileDomain struct {
	imageSize [2]uint32

	tileSize     [2]uint32
	scaleFactors []uint32
}

func NewImageTileDomain(imageSize [2]uint32, tileSpec iiifimageapi.ImageInformationTile) ImageTileDomain {
	d := ImageTileDomain{
		imageSize:    imageSize,
		tileSize:     [2]uint32{tileSpec.Width, tileSpec.Height},
		scaleFactors: tileSpec.ScaleFactors,
	}

	if d.tileSize[1] == 0 {
		// height is considered optional
		d.tileSize[1] = d.tileSize[0]
	}

	return d
}

func (d ImageTileDomain) Contains(p Value) bool {
	for _, scaleFactor := range d.scaleFactors {
		scaledTileWidth := d.tileSize[0] * scaleFactor
		scaledTileHeight := d.tileSize[1] * scaleFactor

		if p.Region[2] == scaledTileWidth && p.Size[0] == d.tileSize[0] {
			// easy
		} else if d.imageSize[0]-p.Region[0] == p.Region[2] && p.Size[0] == uint32(math.Ceil(float64(p.Region[2])*(1/float64(scaleFactor)))) {
			// good
		} else {
			continue
		}

		if p.Region[3] == scaledTileHeight && p.Size[1] == d.tileSize[1] {
			// easy
		} else if d.imageSize[1]-p.Region[1] == p.Region[3] && p.Size[1] == uint32(math.Ceil(float64(p.Region[3])*(1/float64(scaleFactor)))) {
			// good
		} else {
			continue
		}

		return true
	}

	return false
}

func (d ImageTileDomain) Enumerate() ValueList {
	uniq := valueMap{}

	d.enumerateMap(&uniq)

	return uniq.List()
}

func (d ImageTileDomain) enumerateMap(uniqPtr *valueMap) {
	uniq := *uniqPtr

	for _, scaleFactor := range d.scaleFactors {
		scaledTileWidth := d.tileSize[0] * scaleFactor
		scaledTileHeight := d.tileSize[1] * scaleFactor

		maxTileX := uint32(math.Floor(float64(d.imageSize[0]) / float64(scaledTileWidth)))
		maxTileY := uint32(math.Floor(float64(d.imageSize[1]) / float64(scaledTileHeight)))

		for tileX := uint32(0); tileX <= maxTileX; tileX++ {
			for tileY := uint32(0); tileY <= maxTileY; tileY++ {
				regionX := tileX * scaledTileWidth
				regionY := tileY * scaledTileHeight
				regionWidth := scaledTileWidth
				regionHeight := scaledTileHeight

				sizeW := d.tileSize[0]
				sizeH := d.tileSize[1]

				if regionX+regionWidth > d.imageSize[0] {
					regionWidth = d.imageSize[0] - regionX
					sizeW = uint32(math.Ceil(float64(regionWidth) / float64(scaleFactor)))
				}

				if regionY+regionHeight > d.imageSize[1] {
					regionHeight = d.imageSize[1] - regionY
					sizeH = uint32(math.Ceil(float64(regionHeight) / float64(scaleFactor)))
				}

				uniq[Value{
					Region: [4]uint32{regionX, regionY, regionWidth, regionHeight},
					Size:   [2]uint32{sizeW, sizeH},
				}] = struct{}{}
			}
		}
	}
}
