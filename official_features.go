package iiifimageapi

const (
	// FeatureNameBaseUriRedirect means the base URI of the service will redirect to the image information document.
	FeatureNameBaseUriRedirect FeatureName = "baseUriRedirect"

	// FeatureNameCanonicalLinkHeader means the canonical image URI HTTP link header is provided on image responses.
	FeatureNameCanonicalLinkHeader FeatureName = "canonicalLinkHeader"

	// FeatureNameCors means the CORS HTTP headers are provided on all responses.
	FeatureNameCors FeatureName = "cors"

	// FeatureNameJsonldMediaType means the JSON-LD media type is provided when requested.
	FeatureNameJsonldMediaType FeatureName = "jsonldMediaType"

	// FeatureNameMirroring means the image may be rotated around the vertical axis, resulting in a left-to-right mirroring of the content.
	FeatureNameMirroring FeatureName = "mirroring"

	// FeatureNameProfileLinkHeader means the profile HTTP link header is provided on image responses.
	FeatureNameProfileLinkHeader FeatureName = "profileLinkHeader"

	// FeatureNameRegionByPct means regions of the full image may be requested by percentage.
	FeatureNameRegionByPct FeatureName = "regionByPct"

	// FeatureNameRegionByPx means regions of the full image may be requested by pixel dimensions.
	FeatureNameRegionByPx FeatureName = "regionByPx"

	// FeatureNameRegionSquare means a square region may be requested, where the width and height are equal to the shorter dimension of the full image.
	FeatureNameRegionSquare FeatureName = "regionSquare"

	// FeatureNameRotationArbitrary means image rotation may be requested using values other than multiples of 90 degrees.
	FeatureNameRotationArbitrary FeatureName = "rotationArbitrary"

	// FeatureNameRotationBy90s means image rotation may be requested in multiples of 90 degrees.
	FeatureNameRotationBy90s FeatureName = "rotationBy90s"

	// FeatureNameSizeByConfinedWh means image size may be requested in the form !w,h.
	FeatureNameSizeByConfinedWh FeatureName = "sizeByConfinedWh"

	// FeatureNameSizeByH means image size may be requested in the form ,h.
	FeatureNameSizeByH FeatureName = "sizeByH"

	// FeatureNameSizeByPct means images size may be requested in the form pct:n.
	FeatureNameSizeByPct FeatureName = "sizeByPct"

	// FeatureNameSizeByW means image size may be requested in the form w,.
	FeatureNameSizeByW FeatureName = "sizeByW"

	// FeatureNameSizeByWh means image size may be requested in the form w,h.
	FeatureNameSizeByWh FeatureName = "sizeByWh"

	// FeatureNameSizeUpscaling means image sizes prefixed with ^ may be requested.
	FeatureNameSizeUpscaling FeatureName = "sizeUpscaling"
)
