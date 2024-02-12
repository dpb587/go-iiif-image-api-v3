package iiifimageapi

const (
	ComplianceLevel0Name            ComplianceLevelName = "level0"
	ComplianceLevel0ProfileDocument string              = "http://iiif.io/api/image/3/level0.json"

	ComplianceLevel1Name            ComplianceLevelName = "level1"
	ComplianceLevel1ProfileDocument string              = "http://iiif.io/api/image/3/level1.json"

	ComplianceLevel2Name            ComplianceLevelName = "level2"
	ComplianceLevel2ProfileDocument string              = "http://iiif.io/api/image/3/level2.json"
)

type officialComplianceLevelSpec struct {
	name            ComplianceLevelName
	profileDocument string
	baseFeatures    FeatureNameList
	baseQualities   []string
	baseFormats     []string
}

func (cl officialComplianceLevelSpec) Name() ComplianceLevelName {
	return cl.name
}

func (cl officialComplianceLevelSpec) ProfileDocument() string {
	return cl.profileDocument
}

func (cl officialComplianceLevelSpec) BaseFeatures() FeatureNameList {
	return cl.baseFeatures
}

func (cl officialComplianceLevelSpec) BaseQualities() []string {
	return cl.baseQualities
}

func (cl officialComplianceLevelSpec) BaseFormats() []string {
	return cl.baseFormats
}

var (
	complianceLevel0Spec = &officialComplianceLevelSpec{
		name:            ComplianceLevel0Name,
		profileDocument: ComplianceLevel0ProfileDocument,
		baseQualities:   []string{"default"},
		baseFormats:     []string{"jpg"},
	}

	complianceLevel1Spec = &officialComplianceLevelSpec{
		name:            ComplianceLevel1Name,
		profileDocument: ComplianceLevel1ProfileDocument,
		baseFeatures: FeatureNameList{
			FeatureNameRegionByPx,
			FeatureNameRegionSquare,
			FeatureNameSizeByW,
			FeatureNameSizeByH,
			FeatureNameSizeByWh,
			FeatureNameBaseUriRedirect,
			FeatureNameCors,
			FeatureNameJsonldMediaType,
		},
		baseQualities: []string{"default"},
		baseFormats:   []string{"jpg"},
	}

	complianceLevel2Spec = &officialComplianceLevelSpec{
		name:            ComplianceLevel2Name,
		profileDocument: ComplianceLevel2ProfileDocument,
		baseFeatures: FeatureNameList{
			FeatureNameRegionByPx,
			FeatureNameRegionByPct,
			FeatureNameRegionSquare,
			FeatureNameSizeByW,
			FeatureNameSizeByH,
			FeatureNameSizeByPct,
			FeatureNameSizeByConfinedWh,
			FeatureNameSizeByWh,
			FeatureNameRotationBy90s,
			FeatureNameBaseUriRedirect,
			FeatureNameCors,
			FeatureNameJsonldMediaType,
		},
		baseQualities: []string{"default", "color", "gray"},
		baseFormats:   []string{"jpg", "png"},
	}
)

type officialComplianceLevels map[ComplianceLevelName]ComplianceLevelSpec

func (cl officialComplianceLevels) GetByName(name ComplianceLevelName) (ComplianceLevelSpec, bool) {
	cls, ok := cl[name]

	return cls, ok
}

// OfficialComplianceLevels is the official levels defined by the spec. [DefaultComplianceLevels] should typically
// be used instead.
var OfficialComplianceLevels = officialComplianceLevels{
	ComplianceLevel0Name: complianceLevel0Spec,
	ComplianceLevel1Name: complianceLevel1Spec,
	ComplianceLevel2Name: complianceLevel2Spec,
}
