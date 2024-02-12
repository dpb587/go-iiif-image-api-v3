package imagerequest

import (
	"strings"
	"testing"

	iiifimageapi "github.com/dpb587/go-iiif-image-api-v3"
)

func mustParseImageRequestParams(path [4]string) ParsedParams {
	v, err := ParseRawParams(path)
	if err != nil {
		panic(err)
	}

	return v
}

func ptrUint32(v uint32) *uint32 {
	return &v
}

/*
 * The following are based on the examples listed in the spec.
 */

func normativeImageInformation() iiifimageapi.ImageInformation {
	return iiifimageapi.ImageInformation{
		Width:          300,
		Height:         200,
		Profile:        iiifimageapi.ComplianceLevel2Name,
		ExtraQualities: []string{"default", "color"},
		ExtraFormats:   []string{"jpg"},
		ExtraFeatures: []iiifimageapi.FeatureName{
			iiifimageapi.FeatureNameBaseUriRedirect,
			iiifimageapi.FeatureNameCanonicalLinkHeader,
			iiifimageapi.FeatureNameCors,
			iiifimageapi.FeatureNameJsonldMediaType,
			iiifimageapi.FeatureNameMirroring,
			iiifimageapi.FeatureNameProfileLinkHeader,
			iiifimageapi.FeatureNameRegionByPct,
			iiifimageapi.FeatureNameRegionByPx,
			iiifimageapi.FeatureNameRegionSquare,
			iiifimageapi.FeatureNameRotationArbitrary,
			iiifimageapi.FeatureNameRotationBy90s,
			iiifimageapi.FeatureNameSizeByConfinedWh,
			iiifimageapi.FeatureNameSizeByH,
			iiifimageapi.FeatureNameSizeByPct,
			iiifimageapi.FeatureNameSizeByW,
			iiifimageapi.FeatureNameSizeByWh,
			iiifimageapi.FeatureNameSizeUpscaling,
		},
	}
}

func TestParsedParams_Resolve_NormativeRegionFull(t *testing.T) {
	resolved, err := mustParseImageRequestParams([4]string{"full", "max", "0", "default.jpg"}).Resolve(
		ResolveOptions{
			ImageInformation: normativeImageInformation(),
			DefaultQuality:   "color",
		},
	)
	if err != nil {
		t.Fatalf("expected `nil` but got: %v", err)
	} else if _e, _a := [4]uint32{0, 0, 300, 200}, resolved.RegionPixels(); _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}
}

func TestParsedParams_Resolve_NormativeRegionSquare(t *testing.T) {
	resolved, err := mustParseImageRequestParams([4]string{"square", "max", "0", "default.jpg"}).Resolve(
		ResolveOptions{
			ImageInformation: normativeImageInformation(),
			DefaultQuality:   "color",
		},
	)
	if err != nil {
		t.Fatalf("expected `nil` but got: %v", err)
	} else if _e, _a := [4]uint32{50, 0, 200, 200}, resolved.RegionPixels(); _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}
}

func TestParsedParams_Resolve_NormativeRegionPixels(t *testing.T) {
	resolved, err := mustParseImageRequestParams([4]string{"125,15,120,140", "max", "0", "default.jpg"}).Resolve(
		ResolveOptions{
			ImageInformation: normativeImageInformation(),
			DefaultQuality:   "color",
		},
	)
	if err != nil {
		t.Fatalf("expected `nil` but got: %v", err)
	} else if _e, _a := [4]uint32{125, 15, 120, 140}, resolved.RegionPixels(); _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}
}

func TestParsedParams_Resolve_NormativeRegionPercent(t *testing.T) {
	resolved, err := mustParseImageRequestParams([4]string{"pct:41.6,7.5,40,70", "max", "0", "default.jpg"}).Resolve(
		ResolveOptions{
			ImageInformation: normativeImageInformation(),
			DefaultQuality:   "color",
		},
	)
	if err != nil {
		t.Fatalf("expected `nil` but got: %v", err)
	} else if _e, _a := [4]uint32{125, 15, 120, 140}, resolved.RegionPixels(); _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}
}

func TestParsedParams_Resolve_NormativeRegionPixelsExtended(t *testing.T) {
	resolved, err := mustParseImageRequestParams([4]string{"125,15,200,200", "max", "0", "default.jpg"}).Resolve(
		ResolveOptions{
			ImageInformation: normativeImageInformation(),
			DefaultQuality:   "color",
		},
	)
	if err != nil {
		t.Fatalf("expected `nil` but got: %v", err)
	} else if _e, _a := [4]uint32{125, 15, 175, 185}, resolved.RegionPixels(); _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}
}

func TestParsedParams_Resolve_NormativeRegionPercentExtended(t *testing.T) {
	resolved, err := mustParseImageRequestParams([4]string{"pct:41.6,7.5,66.6,100", "max", "0", "default.jpg"}).Resolve(
		ResolveOptions{
			ImageInformation: normativeImageInformation(),
			DefaultQuality:   "color",
		},
	)
	if err != nil {
		t.Fatalf("expected `nil` but got: %v", err)
	} else if _e, _a := [4]uint32{125, 15, 175, 185}, resolved.RegionPixels(); _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}
}

func TestParsedParams_Resolve_NormativeSizeMaxWidth(t *testing.T) {
	info := normativeImageInformation()
	info.MaxWidth = ptrUint32(200)

	resolved, err := mustParseImageRequestParams([4]string{"full", "max", "0", "default.jpg"}).Resolve(
		ResolveOptions{
			ImageInformation: info,
			DefaultQuality:   "color",
		},
	)
	if err != nil {
		t.Fatalf("expected `nil` but got: %v", err)
	} else if _e, _a := [2]uint32{200, 133}, resolved.SizePixels(); _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}
}

func TestParsedParams_Resolve_NormativeSizeMaxWidthUpscale(t *testing.T) {
	t.Skip("example conflicts with validator https://github.com/IIIF/image-validator/blob/890b120d19a7db1cde8a64da0ae6a986ae13e342/iiif_validator/tests/size_up.py#L33")

	info := normativeImageInformation()
	info.MaxWidth = ptrUint32(360)

	resolved, err := mustParseImageRequestParams([4]string{"full", "^max", "0", "default.jpg"}).Resolve(
		ResolveOptions{
			ImageInformation: info,
			DefaultQuality:   "color",
		},
	)
	if err != nil {
		t.Fatalf("expected `nil` but got: %v", err)
	} else if _e, _a := [2]uint32{360, 240}, resolved.SizePixels(); _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}
}

func TestParsedParams_Resolve_NormativeSizeWidth(t *testing.T) {
	resolved, err := mustParseImageRequestParams([4]string{"full", "150,", "0", "default.jpg"}).Resolve(
		ResolveOptions{
			ImageInformation: normativeImageInformation(),
			DefaultQuality:   "color",
		},
	)
	if err != nil {
		t.Fatalf("expected `nil` but got: %v", err)
	} else if _e, _a := [2]uint32{150, 100}, resolved.SizePixels(); _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}
}

func TestParsedParams_Resolve_NormativeSizeWidthUpscale(t *testing.T) {
	resolved, err := mustParseImageRequestParams([4]string{"full", "^360,", "0", "default.jpg"}).Resolve(
		ResolveOptions{
			ImageInformation: normativeImageInformation(),
			DefaultQuality:   "color",
		},
	)
	if err != nil {
		t.Fatalf("expected `nil` but got: %v", err)
	} else if _e, _a := [2]uint32{360, 240}, resolved.SizePixels(); _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}
}

func TestParsedParams_Resolve_NormativeSizeHeight(t *testing.T) {
	resolved, err := mustParseImageRequestParams([4]string{"full", ",150", "0", "default.jpg"}).Resolve(
		ResolveOptions{
			ImageInformation: normativeImageInformation(),
			DefaultQuality:   "color",
		},
	)
	if err != nil {
		t.Fatalf("expected `nil` but got: %v", err)
	} else if _e, _a := [2]uint32{225, 150}, resolved.SizePixels(); _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}
}

func TestParsedParams_Resolve_NormativeSizeHeightUpscale(t *testing.T) {
	resolved, err := mustParseImageRequestParams([4]string{"full", "^,240", "0", "default.jpg"}).Resolve(
		ResolveOptions{
			ImageInformation: normativeImageInformation(),
			DefaultQuality:   "color",
		},
	)
	if err != nil {
		t.Fatalf("expected `nil` but got: %v", err)
	} else if _e, _a := [2]uint32{360, 240}, resolved.SizePixels(); _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}
}

func TestParsedParams_Resolve_NormativeSizePercent(t *testing.T) {
	resolved, err := mustParseImageRequestParams([4]string{"full", "pct:50", "0", "default.jpg"}).Resolve(
		ResolveOptions{
			ImageInformation: normativeImageInformation(),
			DefaultQuality:   "color",
		},
	)
	if err != nil {
		t.Fatalf("expected `nil` but got: %v", err)
	} else if _e, _a := [2]uint32{150, 100}, resolved.SizePixels(); _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}
}

// TODO report spec inconsistency (missing caret); see also https://github.com/IIIF/api/issues/2070
func TestParsedParams_Resolve_NormativeSizePercentUpscale(t *testing.T) {
	resolved, err := mustParseImageRequestParams([4]string{"full", "^pct:120", "0", "default.jpg"}).Resolve(
		ResolveOptions{
			ImageInformation: normativeImageInformation(),
			DefaultQuality:   "color",
		},
	)
	if err != nil {
		t.Fatalf("expected `nil` but got: %v", err)
	} else if _e, _a := [2]uint32{360, 240}, resolved.SizePixels(); _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}
}

func TestParsedParams_Resolve_NormativeSizePixels(t *testing.T) {
	resolved, err := mustParseImageRequestParams([4]string{"full", "225,100", "0", "default.jpg"}).Resolve(
		ResolveOptions{
			ImageInformation: normativeImageInformation(),
			DefaultQuality:   "color",
		},
	)
	if err != nil {
		t.Fatalf("expected `nil` but got: %v", err)
	} else if _e, _a := [2]uint32{225, 100}, resolved.SizePixels(); _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}
}

func TestParsedParams_Resolve_NormativeSizePixelsUpscale(t *testing.T) {
	resolved, err := mustParseImageRequestParams([4]string{"full", "^360,360", "0", "default.jpg"}).Resolve(
		ResolveOptions{
			ImageInformation: normativeImageInformation(),
			DefaultQuality:   "color",
		},
	)
	if err != nil {
		t.Fatalf("expected `nil` but got: %v", err)
	} else if _e, _a := [2]uint32{360, 360}, resolved.SizePixels(); _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}
}

func TestParsedParams_Resolve_NormativeSizePixelsConfine(t *testing.T) {
	resolved, err := mustParseImageRequestParams([4]string{"full", "!225,100", "0", "default.jpg"}).Resolve(
		ResolveOptions{
			ImageInformation: normativeImageInformation(),
			DefaultQuality:   "color",
		},
	)
	if err != nil {
		t.Fatalf("expected `nil` but got: %v", err)
	} else if _e, _a := [2]uint32{150, 100}, resolved.SizePixels(); _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}
}

func TestParsedParams_Resolve_NormativeSizePixelsConfineUpscale(t *testing.T) {
	resolved, err := mustParseImageRequestParams([4]string{"full", "^!360,360", "0", "default.jpg"}).Resolve(
		ResolveOptions{
			ImageInformation: normativeImageInformation(),
			DefaultQuality:   "color",
		},
	)
	if err != nil {
		t.Fatalf("expected `nil` but got: %v", err)
	} else if _e, _a := [2]uint32{360, 240}, resolved.SizePixels(); _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}
}

// iiif-validator

func TestParsedParams_Resolve_ValidatorSizeRequireUpscale(t *testing.T) {
	_, err := mustParseImageRequestParams([4]string{"full", "1000,1000", "0", "default.jpg"}).Resolve(
		ResolveOptions{
			ImageInformation: normativeImageInformation(),
			DefaultQuality:   "color",
		},
	)
	if err == nil {
		t.Fatal("expected error but got none")
	} else if _e, _a := "upscale is required:", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	}
}

func TestParsedParams_Resolve_ValidatorConfinedSizeRequireUpscale(t *testing.T) {
	_, err := mustParseImageRequestParams([4]string{"full", "!1000,1000", "0", "default.jpg"}).Resolve(
		ResolveOptions{
			ImageInformation: normativeImageInformation(),
			DefaultQuality:   "color",
		},
	)
	if err == nil {
		t.Fatal("expected error but got none")
	} else if _e, _a := "upscale is required:", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	}
}
