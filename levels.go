package iiifimageapi

type ComplianceLevelName string

type ComplianceLevelSpec interface {
	Name() ComplianceLevelName
	ProfileDocument() string

	BaseFeatures() FeatureNameList
	BaseQualities() []string
	BaseFormats() []string
}

type ComplianceLevels interface {
	GetByName(name ComplianceLevelName) (ComplianceLevelSpec, bool)
}

// DefaultComplianceLevels is a shared set of compliance levels that may be reconfigured. By default it is the set of
// levels defined in the specification.
var DefaultComplianceLevels ComplianceLevels = OfficialComplianceLevels
