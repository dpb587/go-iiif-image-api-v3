package iiifimageapi

type ImageInformation struct {
	Context  string              `json:"@context"`
	ID       string              `json:"id"`
	Type     string              `json:"type"`
	Protocol string              `json:"protocol"`
	Profile  ComplianceLevelName `json:"profile"`

	Width  uint32 `json:"width"`
	Height uint32 `json:"height"`

	Sizes []ImageInformationSize `json:"sizes,omitempty"`
	Tiles []ImageInformationTile `json:"tiles,omitempty"`

	PreferredFormats []string `json:"preferredFormats,omitempty"`

	ExtraQualities []string        `json:"extraQualities,omitempty"`
	ExtraFormats   []string        `json:"extraFormats,omitempty"`
	ExtraFeatures  FeatureNameList `json:"extraFeatures,omitempty"`

	MaxHeight *uint32 `json:"maxHeight,omitempty"`
	MaxWidth  *uint32 `json:"maxWidth,omitempty"`
	MaxArea   *uint32 `json:"maxArea,omitempty"`

	Rights  string        `json:"rights,omitempty"`
	PartOf  []interface{} `json:"partOf,omitempty"`
	SeeAlso []interface{} `json:"seeAlso,omitempty"`
	Service []interface{} `json:"service,omitempty"`
}

type ImageInformationSize struct {
	Width  uint32 `json:"width"`
	Height uint32 `json:"height"`
}

type ImageInformationTile struct {
	Width        uint32   `json:"width"`
	Height       uint32   `json:"height,omitempty"`
	ScaleFactors []uint32 `json:"scaleFactors"`
}

// NewImageInformation is a light wrapper to add the default, spec-defined properties:
// context, type, and protocol.
func NewImageInformation(info ImageInformation) ImageInformation {
	info.Context = Context
	info.Type = Type
	info.Protocol = Protocol

	return info
}
