package imagerequest

import (
	"fmt"
	"math"

	iiifimageapi "github.com/dpb587/go-iiif-image-api-v3"
)

func featureNameMap(fnl []iiifimageapi.FeatureName) map[iiifimageapi.FeatureName]struct{} {
	out := map[iiifimageapi.FeatureName]struct{}{}

	for _, fn := range fnl {
		out[fn] = struct{}{}
	}

	return out
}

func stringMap(fnl []string) map[string]struct{} {
	out := map[string]struct{}{}

	for _, fn := range fnl {
		out[fn] = struct{}{}
	}

	return out
}

type maxConstraint struct {
	maxHeight *uint32
	maxWidth  *uint32
	maxArea   *uint32
}

func (im maxConstraint) Validate(in [2]uint32) error {
	if im.maxWidth != nil && in[0] > *im.maxWidth {
		return iiifimageapi.NewInvalidValueError(fmt.Sprintf("width (%d) exceeds max width (%d)", in[0], *im.maxWidth))
	} else if im.maxHeight != nil && in[1] > *im.maxHeight {
		return iiifimageapi.NewInvalidValueError(fmt.Sprintf("height (%d) exceeds max height (%d)", in[1], *im.maxHeight))
	} else if im.maxArea != nil && area(in) > *im.maxArea {
		return iiifimageapi.NewInvalidValueError(fmt.Sprintf("area (%d) exceeds max area (%d)", area(in), *im.maxArea))
	}

	return nil
}

func (im maxConstraint) Constrain(wh [2]uint32, upscale bool) [2]uint32 {
	var nextwh = wh

	if im.Validate(nextwh) == nil {
		if !upscale || (im.maxWidth == nil && im.maxHeight == nil && im.maxArea == nil) {
			return nextwh
		}
	}

	var constrained = false

	{
		var constrainWidth = im.maxWidth != nil && *im.maxWidth < nextwh[0]
		var constrainHeight = im.maxHeight != nil && *im.maxHeight < nextwh[1]

		if constrainWidth && constrainHeight {
			if (float64(nextwh[0]) / float64(*im.maxWidth)) > (float64(nextwh[1]) / float64(*im.maxHeight)) {
				constrainHeight = false
			} else {
				constrainWidth = false
			}
		}

		if constrainWidth {
			constrained = true
			nextwh = [2]uint32{
				*im.maxWidth,
				uint32(math.Floor(float64(nextwh[1]) * float64(*im.maxWidth) / float64(nextwh[0]))),
			}
		} else if constrainHeight {
			constrained = true
			nextwh = [2]uint32{
				uint32(math.Floor(float64(nextwh[0]) * float64(*im.maxHeight) / float64(nextwh[1]))),
				*im.maxHeight,
			}
		}
	}

	nextwhArea := area(nextwh)

	if im.maxArea != nil && nextwhArea > *im.maxArea {
		scalewh := math.Sqrt(float64(nextwhArea) / float64(*im.maxArea))

		constrained = true
		nextwh = [2]uint32{
			uint32(math.Floor(float64(nextwh[0]) / scalewh)),
			uint32(math.Floor(float64(nextwh[1]) / scalewh)),
		}
	}

	if constrained || !upscale {
		return nextwh
	}

	{
		var upscaleWidth = im.maxWidth != nil && nextwh[0] < *im.maxWidth
		var upscaleHeight = im.maxHeight != nil && nextwh[1] < *im.maxHeight

		if upscaleWidth && upscaleHeight {
			if (float64(nextwh[0]) / float64(*im.maxWidth)) > (float64(nextwh[1]) / float64(*im.maxHeight)) {
				upscaleHeight = false
			} else {
				upscaleWidth = false
			}
		}

		if upscaleWidth {
			nextwh = [2]uint32{
				*im.maxWidth,
				uint32(math.Floor(float64(nextwh[1]) * float64(*im.maxWidth) / float64(nextwh[0]))),
			}
		} else if upscaleHeight {
			nextwh = [2]uint32{
				uint32(math.Floor(float64(nextwh[0]) * float64(*im.maxHeight) / float64(nextwh[1]))),
				*im.maxHeight,
			}
		}
	}

	if im.maxArea != nil && nextwhArea < *im.maxArea {
		scale := math.Sqrt(float64(*im.maxArea) / float64(nextwh[0]) / float64(nextwh[1]))

		nextwh = [2]uint32{
			uint32(math.Floor(float64(nextwh[0]) * scale)),
			uint32(math.Floor(float64(nextwh[1]) * scale)),
		}

		if im.maxHeight != nil || im.maxWidth != nil {
			// lazy algo; make sure we didn't upscale past a constraint
			return maxConstraint{
				maxHeight: im.maxHeight,
				maxWidth:  im.maxWidth,
			}.Constrain(nextwh, false)
		}
	}

	return nextwh
}

func area(in [2]uint32) uint32 {
	return in[0] * in[1]
}
