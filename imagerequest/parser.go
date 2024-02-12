package imagerequest

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	iiifimageapi "github.com/dpb587/go-iiif-image-api-v3"
)

// ParseRawParams parses [RawParams] values and performs syntactic validation of the values. The result represents
// specification-valid parameters, however it may not be valid for an individual image. The [ParsedParams.Resolve]
// function should be used to ensure it is fully valid.
func ParseRawParams(segments RawParams) (ParsedParams, error) {
	var p ParsedParams

	{ // region
		value := segments[0]

		if value == "full" {
			p.RegionIsEnum = true
			p.RegionEnum = "full"
		} else if value == "square" {
			p.RegionIsEnum = true
			p.RegionEnum = "square"
		} else if strings.HasPrefix(value, "pct:") {
			p.RegionIsPercent = true

			regionSplit := strings.SplitN(strings.TrimPrefix(value, "pct:"), ",", 5)
			if len(regionSplit) != 4 {
				return ParsedParams{}, iiifimageapi.NewInvalidValueError("parsing region: percents: invalid value format (expecting `pct:*,*,*,*`)")
			}

			for fieldIdx, field := range regionSplit {
				v, err := strconv.ParseFloat(field, 32)
				if err != nil || v < 0 || v > 100 {
					return ParsedParams{}, iiifimageapi.NewInvalidValueError(fmt.Sprintf("parsing region: percents[%d]: invalid value (expecting float, range [0, 100])", fieldIdx))
				}

				p.RegionPercent[fieldIdx] = float32(v)
			}
		} else {
			regionSplit := strings.SplitN(value, ",", 5)
			if len(regionSplit) != 4 {
				return ParsedParams{}, iiifimageapi.NewInvalidValueError("parsing region: pixels: invalid value format (expecting `*,*,*,*`)")
			}

			for fieldIdx, field := range regionSplit {
				v, err := strconv.ParseUint(field, 10, 32)
				if err != nil {
					return ParsedParams{}, iiifimageapi.NewInvalidValueError(fmt.Sprintf("parsing region: pixels[%d]: invalid value (expecting integer, range [0, ))", fieldIdx))
				}

				p.RegionPixels[fieldIdx] = uint32(v)
			}
		}
	}

	{ // size
		value := segments[1]

		if strings.HasPrefix(value, "^") {
			p.SizeIsUpscaled = true
			value = strings.TrimPrefix(value, "^")
		}

		if strings.HasPrefix(value, "!") {
			p.SizeIsConfined = true
			value = strings.TrimPrefix(value, "!")
		}

		if value == "max" {
			p.SizeIsEnum = true
			p.SizeEnum = "max"
		} else if strings.HasPrefix(value, "pct:") {
			p.SizeIsPercent = true

			v, err := strconv.ParseFloat(strings.TrimPrefix(value, "pct:"), 32)
			if err != nil {
				return ParsedParams{}, iiifimageapi.NewInvalidValueError("parsing size: percent: invalid value (expecting float)")
			}

			if v < 0 {
				return ParsedParams{}, iiifimageapi.NewInvalidValueError("parsing size: percent: invalid value (expecting range (0, ))")
			} else if v > 100 && !p.SizeIsUpscaled {
				return ParsedParams{}, iiifimageapi.NewInvalidValueError("parsing size: percent: invalid value (expecting range (0, 100] for non-upscaled)")
			}

			p.SizePercent = float32(v)
		} else {
			sizeSplit := strings.SplitN(value, ",", 3)
			if len(sizeSplit) != 2 {
				return ParsedParams{}, iiifimageapi.NewInvalidValueError("parsing size: pixels: invalid value format (expecting `*,*`)")
			}

			for fieldIdx, field := range sizeSplit {
				if len(field) == 0 {
					continue
				}

				v, err := strconv.ParseUint(field, 10, 32)
				if err != nil {
					return ParsedParams{}, iiifimageapi.NewInvalidValueError(fmt.Sprintf("parsing size: pixels[%d]: invalid value (expecting int)", fieldIdx))
				} else if v <= 0 {
					return ParsedParams{}, iiifimageapi.NewInvalidValueError(fmt.Sprintf("parsing size: pixels[%d]: invalid value (expecting range (0, ))", fieldIdx))
				}

				vT := uint32(v)

				p.SizePixels[fieldIdx] = &vT
			}
		}
	}

	{ // rotation
		value := segments[2]

		if strings.HasPrefix(value, "!") {
			p.RotationIsMirrored = true
			value = strings.TrimPrefix(value, "!")
		}

		if strings.HasPrefix(value, "-") {
			return ParsedParams{}, iiifimageapi.NewInvalidValueError("parsing rotation: invalid value (expecting range [0, 360])")
		}

		if strings.Contains(value, ".") {
			v, err := strconv.ParseFloat(value, 32)
			if err != nil {
				return ParsedParams{}, iiifimageapi.NewInvalidValueError("parsing rotation: invalid value (expecting float or int)")
			} else if v < 0 || v > 360 {
				return ParsedParams{}, iiifimageapi.NewInvalidValueError("parsing rotation: invalid value (expecting range [0, 360])")
			}

			p.RotationAmount = float32(v)
		} else {
			v, err := strconv.ParseInt(value, 10, 0)
			if err != nil {
				return ParsedParams{}, iiifimageapi.NewInvalidValueError("parsing rotation: invalid value (expecting float or int)")
			} else if v < 0 || v > 360 {
				return ParsedParams{}, iiifimageapi.NewInvalidValueError("parsing rotation: invalid value (expecting range [0, 360])")
			}

			p.RotationAmount = float32(v)
		}
	}

	{ // file
		ext := filepath.Ext(segments[3])
		if len(ext) < 2 {
			return ParsedParams{}, iiifimageapi.NewInvalidValueError("parsing file: invalid format (expecting `*.*`)")
		}

		p.Quality = strings.TrimSuffix(segments[3], ext)
		p.Format = strings.TrimPrefix(ext, ".")
	}

	return p, nil
}
