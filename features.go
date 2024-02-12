package iiifimageapi

import (
	"sort"
	"strings"
)

type FeatureName string

type FeatureNameList []FeatureName

func (fnl FeatureNameList) Sort() {
	sort.Slice(fnl, func(i, j int) bool {
		return strings.Compare(string(fnl[i]), string(fnl[j])) < 0
	})
}
