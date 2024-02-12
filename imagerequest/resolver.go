package imagerequest

import (
	"errors"
	"fmt"
	"math"

	iiifimageapi "github.com/dpb587/go-iiif-image-api-v3"
	"github.com/dpb587/go-iiif-image-api-v3/pixelset"
)

// ResolverOptions contains the properties which affect how parameters are interpreted.
type ResolveOptions struct {
	// ImageInformation must be configured with, at a minimum, the profile, width, and height fields.
	ImageInformation iiifimageapi.ImageInformation

	// DefaultQuality must be a valid quality which will be used when the input is "default".
	DefaultQuality string

	// ComplianceLevels may be a set to supported levels to use. If nil, [iiifimageapi.DefaultComplianceLevels] will be used.
	ComplianceLevels iiifimageapi.ComplianceLevels

	// IgnoreFeatureErrors may be set to true to ignore would-be errors due to unsupported features (according to
	// ImageInformation). Typically this should only be used when a caller is acting in an administrative role.
	IgnoreFeatureErrors bool

	// IgnoreFeatureErrors may be set to true to ignore max width/height/area constraints (according to ImageInformation).
	// Typically this should only be used when a caller is acting in an administrative role.
	IgnoreMaxConstraints bool
}

func (r ResolveOptions) getComplianceLevels() iiifimageapi.ComplianceLevels {
	if r.ComplianceLevels != nil {
		return r.ComplianceLevels
	}

	return iiifimageapi.DefaultComplianceLevels
}

// Resolve validates all properties to ensure they are supported by the given [iiifimageapi.ImageInformation] and converts any
// relative values to absolute pixel values. The result should be usable by an image processor with no additional
// calculations.
func (p ParsedParams) Resolve(opts ResolveOptions) (ResolvedParams, error) {
	var resolved ResolvedParams

	if opts.ImageInformation.Width == 0 || opts.ImageInformation.Height == 0 {
		return ResolvedParams{}, errors.New("invalid options: original image height and width must not be 0")
	} else if opts.DefaultQuality == "" {
		return ResolvedParams{}, errors.New("invalid options: default quality must not be empty")
	}

	cl, ok := opts.getComplianceLevels().GetByName(opts.ImageInformation.Profile)
	if !ok {
		return ResolvedParams{}, fmt.Errorf("invalid options: image profile (%s) is not supported", opts.ImageInformation.Profile)
	}

	{ // format
		supportedFormats := stringMap(append(cl.BaseFormats(), opts.ImageInformation.ExtraFormats...))

		resolved.format = p.Format

		if _, ok := supportedFormats[resolved.format]; !ok {
			return ResolvedParams{}, iiifimageapi.NewInvalidValueError(fmt.Sprintf("format: value (%s) is not supported", resolved.format))
		}
	}

	{ // quality
		supportedQualities := stringMap(append(cl.BaseQualities(), opts.ImageInformation.ExtraQualities...))

		resolved.quality = p.Quality

		if resolved.quality == "default" {
			resolved.quality = opts.DefaultQuality
		}

		if resolved.quality == "color" || resolved.quality == "gray" {
			// always supported even if unlisted
		} else if _, ok := supportedQualities[resolved.quality]; !ok {
			return ResolvedParams{}, iiifimageapi.NewInvalidValueError(fmt.Sprintf("quality: value (%s) is not supported", resolved.quality))
		}
	}

	var supportedFeatures map[iiifimageapi.FeatureName]struct{}

	if opts.IgnoreFeatureErrors {
		// simpler to pretend like we support everything
		supportedFeatures = featureNameMap(iiifimageapi.FeatureNameList{
			iiifimageapi.FeatureNameRegionSquare,
			iiifimageapi.FeatureNameRegionByPct,
			iiifimageapi.FeatureNameRegionByPx,
			iiifimageapi.FeatureNameSizeUpscaling,
			iiifimageapi.FeatureNameSizeByPct,
			iiifimageapi.FeatureNameSizeByWh,
			iiifimageapi.FeatureNameSizeByConfinedWh,
			iiifimageapi.FeatureNameSizeByW,
			iiifimageapi.FeatureNameSizeByH,
			iiifimageapi.FeatureNameMirroring,
			iiifimageapi.FeatureNameRotationBy90s,
			iiifimageapi.FeatureNameRotationArbitrary,
		})
	} else {
		supportedFeatures = featureNameMap(append(cl.BaseFeatures(), opts.ImageInformation.ExtraFeatures...))
	}

	var supportedMax maxConstraint

	if !opts.IgnoreMaxConstraints {
		supportedMax = maxConstraint{
			maxArea:   opts.ImageInformation.MaxArea,
			maxHeight: opts.ImageInformation.MaxHeight,
			maxWidth:  opts.ImageInformation.MaxWidth,
		}
	}

	var deferredPixelsFeatureErr error

	{ // region
		if p.RegionIsEnum {
			switch p.RegionEnum {
			case "full":
				resolved.regionPixels = [4]uint32{
					0,
					0,
					opts.ImageInformation.Width,
					opts.ImageInformation.Height,
				}
			case "square":
				if _, ok := supportedFeatures[iiifimageapi.FeatureNameRegionSquare]; !ok {
					return ResolvedParams{}, iiifimageapi.FeatureNotSupportedError(iiifimageapi.FeatureNameRegionSquare)
				}

				shorter := uint32(math.Min(float64(opts.ImageInformation.Width), float64(opts.ImageInformation.Height)))

				resolved.regionPixels = [4]uint32{
					uint32(math.Round(float64(opts.ImageInformation.Width-shorter) / 2)),
					uint32(math.Round(float64(opts.ImageInformation.Height-shorter) / 2)),
					shorter,
					shorter,
				}
			default:
				return ResolvedParams{}, iiifimageapi.NewInvalidValueError(fmt.Sprintf("region: value (%s) is not valid", p.RegionEnum))
			}
		} else {
			if p.RegionIsPercent {
				if _, ok := supportedFeatures[iiifimageapi.FeatureNameRegionByPct]; !ok {
					return ResolvedParams{}, iiifimageapi.FeatureNotSupportedError(iiifimageapi.FeatureNameRegionByPct)
				}

				resolved.regionPixels = [4]uint32{
					uint32(math.Round(float64(opts.ImageInformation.Width) * float64(p.RegionPercent[0]) / 100)),
					uint32(math.Round(float64(opts.ImageInformation.Height) * float64(p.RegionPercent[1]) / 100)),
					uint32(math.Round(float64(opts.ImageInformation.Width) * float64(p.RegionPercent[2]) / 100)),
					uint32(math.Round(float64(opts.ImageInformation.Height) * float64(p.RegionPercent[3]) / 100)),
				}
			} else {
				if _, ok := supportedFeatures[iiifimageapi.FeatureNameRegionByPx]; !ok {
					if deferredPixelsFeatureErr == nil {
						deferredPixelsFeatureErr = iiifimageapi.FeatureNotSupportedError(iiifimageapi.FeatureNameRegionByPx)
					}
				}

				resolved.regionPixels = [4]uint32{
					p.RegionPixels[0],
					p.RegionPixels[1],
					p.RegionPixels[2],
					p.RegionPixels[3],
				}
			}

			if resolved.regionPixels[0] > opts.ImageInformation.Width {
				return ResolvedParams{}, iiifimageapi.NewInvalidValueError(fmt.Sprintf("region: x (%d) exceeds image width (%d)", resolved.regionPixels[0], opts.ImageInformation.Width))
			} else if resolved.regionPixels[1] > opts.ImageInformation.Height {
				return ResolvedParams{}, iiifimageapi.NewInvalidValueError(fmt.Sprintf("region: y (%d) exceeds image height (%d)", resolved.regionPixels[1], opts.ImageInformation.Height))
			}

			resolved.regionPixels[2] = uint32(math.Min(float64(opts.ImageInformation.Width-resolved.regionPixels[0]), float64(resolved.regionPixels[2])))
			resolved.regionPixels[3] = uint32(math.Min(float64(opts.ImageInformation.Height-resolved.regionPixels[1]), float64(resolved.regionPixels[3])))

			if resolved.regionPixels[2] == 0 {
				return ResolvedParams{}, iiifimageapi.NewInvalidValueError("region: width must be greater than 0")
			} else if resolved.regionPixels[3] == 0 {
				return ResolvedParams{}, iiifimageapi.NewInvalidValueError("region: height must be greater than 0")
			}
		}
	}

	{ // size
		if p.SizeIsUpscaled {
			if _, ok := supportedFeatures[iiifimageapi.FeatureNameSizeUpscaling]; !ok {
				return ResolvedParams{}, iiifimageapi.FeatureNotSupportedError(iiifimageapi.FeatureNameSizeUpscaling)
			}
		}

		if p.SizeIsEnum {
			switch p.SizeEnum {
			case "max":
				wh := supportedMax.Constrain(
					[2]uint32{
						resolved.regionPixels[2],
						resolved.regionPixels[3],
					},
					// originally used p.SizeIsUpscaled, but image-validator expects ^max = original size
					// weird to ignore parameter (vs error), but better to not allow unbounded-max sizes
					false,
				)

				resolved.sizePixels = [2]uint32{wh[0], wh[1]}
			default:
				return ResolvedParams{}, iiifimageapi.NewInvalidValueError(fmt.Sprintf("size: value (%s) is not valid", p.SizeEnum))
			}
		} else if p.SizeIsPercent {
			if _, ok := supportedFeatures[iiifimageapi.FeatureNameSizeByPct]; !ok {
				return ResolvedParams{}, iiifimageapi.FeatureNotSupportedError(iiifimageapi.FeatureNameSizeByPct)
			}

			w := uint32(math.Round(float64(resolved.regionPixels[2]) * float64(p.SizePercent) / 100))
			h := uint32(int(math.Round(float64(resolved.regionPixels[3]) * float64(p.SizePercent) / 100)))

			resolved.sizePixels = [2]uint32{w, h}
		} else if p.SizePixels[0] != nil && p.SizePixels[1] != nil {
			// w,h

			if _, ok := supportedFeatures[iiifimageapi.FeatureNameSizeByWh]; !ok {
				if deferredPixelsFeatureErr == nil {
					deferredPixelsFeatureErr = iiifimageapi.FeatureNotSupportedError(iiifimageapi.FeatureNameSizeByWh)
				}
			}

			if p.SizeIsConfined {
				if _, ok := supportedFeatures[iiifimageapi.FeatureNameSizeByConfinedWh]; !ok {
					return ResolvedParams{}, iiifimageapi.FeatureNotSupportedError(iiifimageapi.FeatureNameSizeByConfinedWh)
				} else if !p.SizeIsUpscaled {
					// our constraint below would otherwise ignore this expected error, so we check early
					if *p.SizePixels[0] > resolved.regionPixels[2] {
						return ResolvedParams{}, iiifimageapi.NewInvalidValueError(fmt.Sprintf("size: upscale is required: w (%d) exceeds region width (%d)", *p.SizePixels[0], resolved.regionPixels[2]))
					} else if *p.SizePixels[1] > resolved.regionPixels[3] {
						return ResolvedParams{}, iiifimageapi.NewInvalidValueError(fmt.Sprintf("size: upscale is required: h (%d) exceeds region height (%d)", *p.SizePixels[1], resolved.regionPixels[3]))
					}
				}

				wh := maxConstraint{
					maxWidth:  p.SizePixels[0],
					maxHeight: p.SizePixels[1],
				}.Constrain(
					[2]uint32{
						resolved.regionPixels[2],
						resolved.regionPixels[3],
					},
					p.SizeIsUpscaled,
				)

				resolved.sizePixels = [2]uint32{wh[0], wh[1]}
			} else {
				w := *p.SizePixels[0]
				h := *p.SizePixels[1]

				resolved.sizePixels = [2]uint32{w, h}
			}
		} else if p.SizePixels[0] != nil && p.SizePixels[1] == nil {
			// w,

			if _, ok := supportedFeatures[iiifimageapi.FeatureNameSizeByW]; !ok {
				return ResolvedParams{}, iiifimageapi.FeatureNotSupportedError(iiifimageapi.FeatureNameSizeByW)
			}

			w := *p.SizePixels[0]
			h := uint32(math.Round(float64(resolved.regionPixels[3]*w) / float64(resolved.regionPixels[2])))

			resolved.sizePixels = [2]uint32{w, h}
		} else if p.SizePixels[0] == nil && p.SizePixels[1] != nil {
			// ,h

			if _, ok := supportedFeatures[iiifimageapi.FeatureNameSizeByH]; !ok {
				return ResolvedParams{}, iiifimageapi.FeatureNotSupportedError(iiifimageapi.FeatureNameSizeByH)
			}

			h := *p.SizePixels[1]
			w := uint32(math.Round(float64(resolved.regionPixels[2]*h) / float64(resolved.regionPixels[3])))

			resolved.sizePixels = [2]uint32{w, h}
		} else {
			return ResolvedParams{}, iiifimageapi.NewInvalidValueError("size: value is not valid")
		}

		if !p.SizeIsUpscaled {
			if resolved.sizePixels[0] > resolved.regionPixels[2] {
				return ResolvedParams{}, iiifimageapi.NewInvalidValueError(fmt.Sprintf("size: upscale is required: w (%d) exceeds region width (%d)", resolved.sizePixels[0], resolved.regionPixels[2]))
			} else if resolved.sizePixels[1] > resolved.regionPixels[3] {
				return ResolvedParams{}, iiifimageapi.NewInvalidValueError(fmt.Sprintf("size: upscale is required: h (%d) exceeds region height (%d)", resolved.sizePixels[1], resolved.regionPixels[3]))
			}
		}

		err := supportedMax.Validate(resolved.sizePixels)
		if err != nil {
			return ResolvedParams{}, iiifimageapi.NewInvalidValueError(fmt.Sprintf("size: %v", err.Error()))
		}
	}

	if deferredPixelsFeatureErr != nil {
		// now that we resolved region+size, if we would have errored about a feature, make sure it isn't an advertised pixelset
		if !pixelset.NewImageDomain(opts.ImageInformation).Contains(pixelset.Value{
			Region: [4]uint32{resolved.regionPixels[0], resolved.regionPixels[1], resolved.regionPixels[2], resolved.regionPixels[3]},
			Size:   [2]uint32{resolved.sizePixels[0], resolved.sizePixels[1]},
		}) {
			return ResolvedParams{}, deferredPixelsFeatureErr
		}
	}

	{ // rotation
		resolved.rotationIsMirrored = p.RotationIsMirrored

		if resolved.rotationIsMirrored {
			if _, ok := supportedFeatures[iiifimageapi.FeatureNameMirroring]; !ok {
				return ResolvedParams{}, iiifimageapi.FeatureNotSupportedError(iiifimageapi.FeatureNameMirroring)
			}
		}

		resolved.rotationAmount = p.RotationAmount

		if resolved.rotationAmount == 360 {
			resolved.rotationAmount = 0
		} else if resolved.rotationAmount != 0 {
			if resolved.rotationAmount == 90 || resolved.rotationAmount == 180 || resolved.rotationAmount == 270 {
				if _, ok := supportedFeatures[iiifimageapi.FeatureNameRotationBy90s]; !ok {
					return ResolvedParams{}, iiifimageapi.FeatureNotSupportedError(iiifimageapi.FeatureNameRotationBy90s)
				}
			} else {
				if _, ok := supportedFeatures[iiifimageapi.FeatureNameRotationArbitrary]; !ok {
					return ResolvedParams{}, iiifimageapi.FeatureNotSupportedError(iiifimageapi.FeatureNameRotationArbitrary)
				}
			}
		}
	}

	resolved.canonical = resolved.ResolveCanonical(opts)

	return resolved, nil
}
