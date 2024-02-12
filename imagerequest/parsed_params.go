package imagerequest

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/dpb587/go-iiif-image-api-v3/pixelset"
)

// ParsedParams represents the destructured values of the image request params. This is typically constructed through
// the [ParseRawParams] function.
type ParsedParams struct {
	RegionIsEnum    bool
	RegionEnum      string
	RegionIsPercent bool
	RegionPercent   [4]float32
	RegionPixels    [4]uint32

	SizeIsConfined bool
	SizeIsUpscaled bool
	SizeIsEnum     bool
	SizeEnum       string
	SizeIsPercent  bool
	SizePercent    float32
	SizePixels     [2]*uint32

	RotationIsMirrored bool
	RotationAmount     float32

	Quality string

	Format string
}

func NewParsedParamsFromPixelset(regionSize pixelset.Value, quality, format string) ParsedParams {
	return ParsedParams{
		RegionPixels: regionSize.Region,
		SizePixels: [2]*uint32{
			&regionSize.Size[0],
			&regionSize.Size[1],
		},
		Quality: quality,
		Format:  format,
	}
}

func (p ParsedParams) Clone() ParsedParams {
	cloned := ParsedParams{
		RegionIsEnum:    p.RegionIsEnum,
		RegionEnum:      p.RegionEnum,
		RegionIsPercent: p.RegionIsPercent,
		RegionPercent:   p.RegionPercent,
		RegionPixels:    p.RegionPixels,

		SizeIsConfined: p.SizeIsConfined,
		SizeIsUpscaled: p.SizeIsUpscaled,
		SizeIsEnum:     p.SizeIsEnum,
		SizeEnum:       p.SizeEnum,
		SizeIsPercent:  p.SizeIsPercent,
		SizePercent:    p.SizePercent,

		RotationIsMirrored: p.RotationIsMirrored,
		RotationAmount:     p.RotationAmount,

		Quality: p.Quality,

		Format: p.Format,
	}

	if p.SizePixels[0] != nil {
		v := *p.SizePixels[0]

		cloned.SizePixels[0] = &v
	}

	if p.SizePixels[1] != nil {
		v := *p.SizePixels[1]

		cloned.SizePixels[1] = &v
	}

	return cloned
}

func (p ParsedParams) String() string {
	// this assumes region + size are valid and does not try to escape them
	return fmt.Sprintf("%s/%s/%s/%s", p.RegionString(), p.SizeString(), url.PathEscape(p.RotationString()), url.PathEscape(p.FileString()))
}

func (p ParsedParams) RegionString() string {
	if p.RegionIsEnum {
		return p.RegionEnum
	} else if p.RegionIsPercent {
		return fmt.Sprintf(
			"pct:%s,%s,%s,%s",
			strconv.FormatFloat(float64(p.RegionPercent[0]), 'f', -1, 32),
			strconv.FormatFloat(float64(p.RegionPercent[1]), 'f', -1, 32),
			strconv.FormatFloat(float64(p.RegionPercent[2]), 'f', -1, 32),
			strconv.FormatFloat(float64(p.RegionPercent[3]), 'f', -1, 32),
		)
	}

	return fmt.Sprintf("%d,%d,%d,%d", p.RegionPixels[0], p.RegionPixels[1], p.RegionPixels[2], p.RegionPixels[3])
}

func (p ParsedParams) SizeString() string {
	var prefix string

	if p.SizeIsConfined {
		prefix += "!"
	}

	if p.SizeIsUpscaled {
		prefix += "^"
	}

	if p.SizeIsEnum {
		return prefix + p.SizeEnum
	} else if p.SizeIsPercent {
		return prefix + strconv.FormatFloat(float64(p.SizePercent), 'f', -1, 64)
	}

	var pixelW, pixelH string

	if p.SizePixels[0] != nil {
		pixelW = strconv.FormatUint(uint64(*p.SizePixels[0]), 10)
	}

	if p.SizePixels[1] != nil {
		pixelH = strconv.FormatUint(uint64(*p.SizePixels[1]), 10)
	}

	return fmt.Sprintf("%s%s,%s", prefix, pixelW, pixelH)
}

func (p ParsedParams) RotationString() string {
	var prefix string

	if p.RotationIsMirrored {
		prefix = "!"
	}

	return prefix + strconv.FormatFloat(float64(p.RotationAmount), 'f', -1, 64)
}

func (p ParsedParams) FileString() string {
	return p.Quality + "." + p.Format
}
