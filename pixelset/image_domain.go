package pixelset

import iiifimageapi "github.com/dpb587/go-iiif-image-api-v3"

type ImageDomain struct {
	imageSize [2]uint32
	sizes     valueMap
	tiles     []ImageTileDomain
}

func NewImageDomain(info iiifimageapi.ImageInformation) ImageDomain {
	d := ImageDomain{
		imageSize: [2]uint32{info.Width, info.Height},
		sizes:     valueMap{},
	}

	for _, size := range info.Sizes {
		d.sizes[Value{
			Region: [4]uint32{0, 0, d.imageSize[0], d.imageSize[1]},
			Size:   [2]uint32{size.Width, size.Height},
		}] = struct{}{}
	}

	for _, tileSpec := range info.Tiles {
		d.tiles = append(d.tiles, NewImageTileDomain([2]uint32{info.Width, info.Height}, tileSpec))
	}

	return d
}

func (d ImageDomain) Contains(p Value) bool {
	if _, known := d.sizes[p]; known {
		return true
	}

	for _, tile := range d.tiles {
		if tile.Contains(p) {
			return true
		}
	}

	return false
}

func (d ImageDomain) Enumerate() ValueList {
	uniq := valueMap{}

	for size := range d.sizes {
		uniq[size] = struct{}{}
	}

	for _, tileDomain := range d.tiles {
		tileDomain.enumerateMap(&uniq)
	}

	return uniq.List()
}
