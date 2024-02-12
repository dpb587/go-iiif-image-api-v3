package main

import (
	"encoding/json"
	"log"

	iiifimageapi "github.com/dpb587/go-iiif-image-api-v3"
	"github.com/dpb587/go-iiif-image-api-v3/imagerequest"
	"github.com/dpb587/go-iiif-image-api-v3/pixelset"
)

func main() {
	// Prepare the image info that we're going to base everything on.
	var imageSize = [2]uint32{3024, 4032}
	var tileWidth uint32 = 512

	imageInfo := iiifimageapi.NewImageInformation(iiifimageapi.ImageInformation{
		ID:      "urn:uuid:f3e2c0b2-dc6a-4bb5-a711-2d8af146b72d",
		Profile: iiifimageapi.ComplianceLevel0Name,
		Width:   imageSize[0],
		Height:  imageSize[1],
		Tiles: []iiifimageapi.ImageInformationTile{
			{
				Width:        tileWidth,
				ScaleFactors: pixelset.GetTileScaleFactors(imageSize, [2]uint32{tileWidth, tileWidth}),
			},
		},
	})

	// Configure how we want to resolve all our pixel regions. Since our imageInfo
	// is listed as a level0 profile, we should ignore all the feature errors
	// because we are intentionally images it would not normally support.
	resolveOptions := imagerequest.ResolveOptions{
		ImageInformation:    imageInfo,
		DefaultQuality:      "color",
		IgnoreFeatureErrors: true,
	}

	// Pretend we have some conventional preferred sizes that we want to calculate
	// and make sure they are added to our image info.
	for _, standard := range []imagerequest.RawParams{
		{"full", "256,256", "0", "default.jpg"},
		{"full", "1024,", "0", "default.jpg"},
		{"full", ",2048", "0", "default.jpg"},
	} {
		parsed, _ := imagerequest.ParseRawParams(standard)
		resolved, _ := parsed.Resolve(resolveOptions)

		imageInfo.Sizes = append(imageInfo.Sizes, iiifimageapi.ImageInformationSize{
			Width:  resolved.SizePixels()[0],
			Height: resolved.SizePixels()[1],
		})
	}

	// Once we have some preferred sizes and tiles known for our image, generate
	// the full list of possible images and render each of them.
	for _, regionSize := range pixelset.NewImageDomain(imageInfo).Enumerate() {
		parsed := imagerequest.NewParsedParamsFromPixelset(regionSize, "default", "jpg")
		resolved, _ := parsed.Resolve(resolveOptions)

		_, err := imageProcessor.Render(resolved)
		if err != nil {
			log.Fatalf("rendering: %v", err)
		}
	}

	// We can convert our image info to JSON and write that out, too.
	err := json.NewEncoder(log.Writer()).Encode(imageInfo)
	if err != nil {
		log.Fatalf("marshaling: %v", err)
	}
}

// Image processing is not within the scope of this package.
var imageProcessor = ImageProcessorTODO{}

type ImageProcessorTODO struct{}

func (ImageProcessorTODO) Render(resolved imagerequest.ResolvedParams) (interface{}, error) {
	log.Printf("fake export: %s\n", resolved.Canonical())

	return nil, nil
}
