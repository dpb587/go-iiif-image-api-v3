package main

import (
	"log"

	iiifimageapi "github.com/dpb587/go-iiif-image-api-v3"
	"github.com/dpb587/go-iiif-image-api-v3/imagerequest"
)

func main() {
	// Params are the raw values, typically received as paths of an HTTP request.
	var params = imagerequest.RawParams{"pct:25,25,50,50", "max", "90", "color.png"}

	// ParseParams extracts typed values from the raw parameters of an image
	// request and ensure they are valid values per the specification.
	parsedParams, err := imagerequest.ParseRawParams(params)
	if err != nil {
		log.Fatalf("parsing params: %v", err)
	}

	// Resolve validates parameters according to a specific image (e.g. profile/
	// feature support, region/size bounds), and converts any relative values into
	// an absolute form (e.g. pct:* to pixels, size upscaling, and quality).
	resolvedParams, err := parsedParams.Resolve(imagerequest.ResolveOptions{
		ImageInformation: iiifimageapi.ImageInformation{
			Profile: iiifimageapi.ComplianceLevel2Name,
			Width:   3024,
			Height:  4032,
		},
		DefaultQuality: "color",
	})
	if err != nil {
		log.Fatalf("resolving params: %v", err)
	}

	// Optionally refer to the canonical form of parameters (per specification).
	canonicalParams := resolvedParams.Canonical()

	log.Printf("params.String() = %s\n", params)
	log.Printf("parsedParams = %#+v\n", parsedParams)
	log.Printf("resolvedParams = %#+v\n", resolvedParams)
	log.Printf("canonicalParams.String() = %s\n", canonicalParams)

	// Call an (external) image processor with the final, resolved configuration.
	_, err = imageProcessor.Render(
		resolvedParams.RegionPixels(),       // [4]int32{X, Y, Width, Height}
		resolvedParams.SizePixels(),         // [2]int32{Width, Height}
		resolvedParams.RotationIsMirrored(), // bool
		resolvedParams.RotationAmount(),     // float32
		resolvedParams.Quality(),            // string
		resolvedParams.Format(),             // string
	)
	if err != nil {
		log.Fatalf("rendering: %v", err)
	}
}

// Image processing is not within the scope of this package.
var imageProcessor = ImageProcessorTODO{}

type ImageProcessorTODO struct{}

func (ImageProcessorTODO) Render(...interface{}) (interface{}, error) {
	return nil, nil
}
