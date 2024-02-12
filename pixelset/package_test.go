package pixelset

import (
	"testing"

	iiifimageapi "github.com/dpb587/go-iiif-image-api-v3"
)

func TestImageDomain_Libvips_IIIF(t *testing.T) {
	// vips dzsave 3888x2592.jpg out/ --layout iiif
	// find out -type f -name '*.jpg' | cut -d/ -f2-3 | pbcopy

	exampleDomain := NewImageDomain(iiifimageapi.ImageInformation{
		Width:  3888,
		Height: 2592,
		Tiles: []iiifimageapi.ImageInformationTile{
			{
				Width:        512,
				ScaleFactors: []uint32{1, 2, 4},
			},
		},
	})

	expectedValuesMap := ValueList{
		{[4]uint32{3072, 0, 512, 512}, [2]uint32{512}},
		{[4]uint32{1536, 2048, 512, 512}, [2]uint32{512}},
		{[4]uint32{2560, 512, 512, 512}, [2]uint32{512}},
		{[4]uint32{2560, 1024, 512, 512}, [2]uint32{512}},
		{[4]uint32{512, 2560, 512, 32}, [2]uint32{512}},
		{[4]uint32{1024, 1024, 512, 512}, [2]uint32{512}},
		{[4]uint32{3072, 1024, 512, 512}, [2]uint32{512}},
		{[4]uint32{2048, 0, 1024, 1024}, [2]uint32{512}},
		{[4]uint32{0, 1024, 512, 512}, [2]uint32{512}},
		{[4]uint32{2048, 2048, 1024, 544}, [2]uint32{512}},
		{[4]uint32{2048, 1024, 512, 512}, [2]uint32{512}},
		{[4]uint32{2048, 0, 1840, 2048}, [2]uint32{460}},
		{[4]uint32{2560, 0, 512, 512}, [2]uint32{512}},
		{[4]uint32{3072, 512, 512, 512}, [2]uint32{512}},
		{[4]uint32{512, 1536, 512, 512}, [2]uint32{512}},
		{[4]uint32{3584, 512, 304, 512}, [2]uint32{304}},
		{[4]uint32{2048, 2048, 512, 512}, [2]uint32{512}},
		{[4]uint32{2560, 2560, 512, 32}, [2]uint32{512}},
		{[4]uint32{512, 512, 512, 512}, [2]uint32{512}},
		{[4]uint32{0, 2048, 2048, 544}, [2]uint32{512}},
		{[4]uint32{1024, 0, 1024, 1024}, [2]uint32{512}},
		{[4]uint32{0, 2048, 1024, 544}, [2]uint32{512}},
		{[4]uint32{0, 0, 1024, 1024}, [2]uint32{512}},
		{[4]uint32{2560, 2048, 512, 512}, [2]uint32{512}},
		{[4]uint32{1536, 1024, 512, 512}, [2]uint32{512}},
		{[4]uint32{3584, 1536, 304, 512}, [2]uint32{304}},
		{[4]uint32{1024, 2048, 512, 512}, [2]uint32{512}},
		{[4]uint32{3072, 2048, 816, 544}, [2]uint32{408}},
		{[4]uint32{3072, 2048, 512, 512}, [2]uint32{512}},
		{[4]uint32{3072, 2560, 512, 32}, [2]uint32{512}},
		// libvips seems to have a bug? it generates scaleFactors [1, 2, 4, 8]
		// but it only lists [1, 2, 4] in the info.json; TODO validate/report
		// {[4]uint32{0, 0, 3888, 2592}, [2]uint32{486}},
		{[4]uint32{1024, 2048, 1024, 544}, [2]uint32{512}},
		{[4]uint32{3584, 0, 304, 512}, [2]uint32{304}},
		{[4]uint32{0, 2048, 512, 512}, [2]uint32{512}},
		{[4]uint32{2560, 1536, 512, 512}, [2]uint32{512}},
		{[4]uint32{3584, 2048, 304, 512}, [2]uint32{304}},
		{[4]uint32{2048, 2048, 1840, 544}, [2]uint32{460}},
		{[4]uint32{1024, 1536, 512, 512}, [2]uint32{512}},
		{[4]uint32{3072, 1536, 512, 512}, [2]uint32{512}},
		{[4]uint32{512, 0, 512, 512}, [2]uint32{512}},
		{[4]uint32{1024, 1024, 1024, 1024}, [2]uint32{512}},
		{[4]uint32{0, 2560, 512, 32}, [2]uint32{512}},
		{[4]uint32{0, 1536, 512, 512}, [2]uint32{512}},
		{[4]uint32{2048, 1536, 512, 512}, [2]uint32{512}},
		{[4]uint32{1536, 2560, 512, 32}, [2]uint32{512}},
		{[4]uint32{1024, 2560, 512, 32}, [2]uint32{512}},
		{[4]uint32{2048, 512, 512, 512}, [2]uint32{512}},
		{[4]uint32{2048, 0, 512, 512}, [2]uint32{512}},
		{[4]uint32{512, 1024, 512, 512}, [2]uint32{512}},
		{[4]uint32{3072, 1024, 816, 1024}, [2]uint32{408}},
		{[4]uint32{3072, 0, 816, 1024}, [2]uint32{408}},
		{[4]uint32{0, 1024, 1024, 1024}, [2]uint32{512}},
		{[4]uint32{0, 512, 512, 512}, [2]uint32{512}},
		{[4]uint32{3584, 2560, 304, 32}, [2]uint32{304}},
		{[4]uint32{512, 2048, 512, 512}, [2]uint32{512}},
		{[4]uint32{1536, 0, 512, 512}, [2]uint32{512}},
		{[4]uint32{3584, 1024, 304, 512}, [2]uint32{304}},
		{[4]uint32{1536, 1536, 512, 512}, [2]uint32{512}},
		{[4]uint32{1536, 512, 512, 512}, [2]uint32{512}},
		{[4]uint32{1024, 512, 512, 512}, [2]uint32{512}},
		{[4]uint32{0, 0, 2048, 2048}, [2]uint32{512}},
		{[4]uint32{2048, 2560, 512, 32}, [2]uint32{512}},
		{[4]uint32{2048, 1024, 1024, 1024}, [2]uint32{512}},
		{[4]uint32{1024, 0, 512, 512}, [2]uint32{512}},
		{[4]uint32{0, 0, 512, 512}, [2]uint32{512}},
	}.toValueMap()

	actualValues := exampleDomain.Enumerate()

	// reverse, but with ignored size height
	actualValuesMap := valueMap{}

	for _, actualValue := range actualValues {
		if !exampleDomain.Contains(actualValue) {
			t.Fatalf("actual value `%v` is not contained in domain", actualValue)
		}

		actualValue.Size[1] = 0 // libvips does not include height

		_, found := expectedValuesMap[actualValue]
		if !found {
			t.Fatalf("actual value `%v` was not in expected list", actualValue)
		}

		actualValuesMap[actualValue] = struct{}{}
	}

	for expectedValue := range expectedValuesMap {
		_, found := actualValuesMap[expectedValue]
		if !found {
			t.Fatalf("expected value `%v` was not in actual list", expectedValue)
		}
	}
}

func TestImageDomain_LibvipsExample_IIIF3(t *testing.T) {
	// vips dzsave 3888x2592.jpg out/ --layout iiif3
	// find out -type f -name '*.jpg' | cut -d/ -f2-3 | pbcopy

	exampleDomain := NewImageDomain(iiifimageapi.ImageInformation{
		Width:  3888,
		Height: 2592,
		Tiles: []iiifimageapi.ImageInformationTile{
			{
				Width:        512,
				ScaleFactors: []uint32{1, 2, 4},
			},
		},
	})

	expectedValuesMap := ValueList{
		{[4]uint32{3072, 0, 512, 512}, [2]uint32{512, 512}},
		{[4]uint32{1536, 2048, 512, 512}, [2]uint32{512, 512}},
		{[4]uint32{2560, 512, 512, 512}, [2]uint32{512, 512}},
		{[4]uint32{2560, 1024, 512, 512}, [2]uint32{512, 512}},
		{[4]uint32{512, 2560, 512, 32}, [2]uint32{512, 32}},
		{[4]uint32{1024, 1024, 512, 512}, [2]uint32{512, 512}},
		{[4]uint32{3072, 1024, 512, 512}, [2]uint32{512, 512}},
		{[4]uint32{2048, 0, 1024, 1024}, [2]uint32{512, 512}},
		{[4]uint32{0, 1024, 512, 512}, [2]uint32{512, 512}},
		{[4]uint32{2048, 2048, 1024, 544}, [2]uint32{512, 272}},
		{[4]uint32{2048, 1024, 512, 512}, [2]uint32{512, 512}},
		{[4]uint32{2048, 0, 1840, 2048}, [2]uint32{460, 512}},
		{[4]uint32{2560, 0, 512, 512}, [2]uint32{512, 512}},
		{[4]uint32{3072, 512, 512, 512}, [2]uint32{512, 512}},
		{[4]uint32{512, 1536, 512, 512}, [2]uint32{512, 512}},
		{[4]uint32{3584, 512, 304, 512}, [2]uint32{304, 512}},
		{[4]uint32{2048, 2048, 512, 512}, [2]uint32{512, 512}},
		{[4]uint32{2560, 2560, 512, 32}, [2]uint32{512, 32}},
		{[4]uint32{512, 512, 512, 512}, [2]uint32{512, 512}},
		{[4]uint32{0, 2048, 2048, 544}, [2]uint32{512, 136}},
		{[4]uint32{1024, 0, 1024, 1024}, [2]uint32{512, 512}},
		{[4]uint32{0, 2048, 1024, 544}, [2]uint32{512, 272}},
		{[4]uint32{0, 0, 1024, 1024}, [2]uint32{512, 512}},
		{[4]uint32{2560, 2048, 512, 512}, [2]uint32{512, 512}},
		{[4]uint32{1536, 1024, 512, 512}, [2]uint32{512, 512}},
		{[4]uint32{3584, 1536, 304, 512}, [2]uint32{304, 512}},
		{[4]uint32{1024, 2048, 512, 512}, [2]uint32{512, 512}},
		{[4]uint32{3072, 2048, 816, 544}, [2]uint32{408, 272}},
		{[4]uint32{3072, 2048, 512, 512}, [2]uint32{512, 512}},
		{[4]uint32{3072, 2560, 512, 32}, [2]uint32{512, 32}},
		// libvips seems to have a bug? it generates scaleFactors [1, 2, 4, 8]
		// but it only lists [1, 2, 4] in the info.json; TODO validate/report
		// {[4]uint32{0, 0, 3888, 2592}, [2]uint32{486, 324}},
		{[4]uint32{1024, 2048, 1024, 544}, [2]uint32{512, 272}},
		{[4]uint32{3584, 0, 304, 512}, [2]uint32{304, 512}},
		{[4]uint32{0, 2048, 512, 512}, [2]uint32{512, 512}},
		{[4]uint32{2560, 1536, 512, 512}, [2]uint32{512, 512}},
		{[4]uint32{3584, 2048, 304, 512}, [2]uint32{304, 512}},
		{[4]uint32{2048, 2048, 1840, 544}, [2]uint32{460, 136}},
		{[4]uint32{1024, 1536, 512, 512}, [2]uint32{512, 512}},
		{[4]uint32{3072, 1536, 512, 512}, [2]uint32{512, 512}},
		{[4]uint32{512, 0, 512, 512}, [2]uint32{512, 512}},
		{[4]uint32{1024, 1024, 1024, 1024}, [2]uint32{512, 512}},
		{[4]uint32{0, 2560, 512, 32}, [2]uint32{512, 32}},
		{[4]uint32{0, 1536, 512, 512}, [2]uint32{512, 512}},
		{[4]uint32{2048, 1536, 512, 512}, [2]uint32{512, 512}},
		{[4]uint32{1536, 2560, 512, 32}, [2]uint32{512, 32}},
		{[4]uint32{1024, 2560, 512, 32}, [2]uint32{512, 32}},
		{[4]uint32{2048, 512, 512, 512}, [2]uint32{512, 512}},
		{[4]uint32{2048, 0, 512, 512}, [2]uint32{512, 512}},
		{[4]uint32{512, 1024, 512, 512}, [2]uint32{512, 512}},
		{[4]uint32{3072, 1024, 816, 1024}, [2]uint32{408, 512}},
		{[4]uint32{3072, 0, 816, 1024}, [2]uint32{408, 512}},
		{[4]uint32{0, 1024, 1024, 1024}, [2]uint32{512, 512}},
		{[4]uint32{0, 512, 512, 512}, [2]uint32{512, 512}},
		{[4]uint32{3584, 2560, 304, 32}, [2]uint32{304, 32}},
		{[4]uint32{512, 2048, 512, 512}, [2]uint32{512, 512}},
		{[4]uint32{1536, 0, 512, 512}, [2]uint32{512, 512}},
		{[4]uint32{3584, 1024, 304, 512}, [2]uint32{304, 512}},
		{[4]uint32{1536, 1536, 512, 512}, [2]uint32{512, 512}},
		{[4]uint32{1536, 512, 512, 512}, [2]uint32{512, 512}},
		{[4]uint32{1024, 512, 512, 512}, [2]uint32{512, 512}},
		{[4]uint32{0, 0, 2048, 2048}, [2]uint32{512, 512}},
		{[4]uint32{2048, 2560, 512, 32}, [2]uint32{512, 32}},
		{[4]uint32{2048, 1024, 1024, 1024}, [2]uint32{512, 512}},
		{[4]uint32{1024, 0, 512, 512}, [2]uint32{512, 512}},
		{[4]uint32{0, 0, 512, 512}, [2]uint32{512, 512}},
	}.toValueMap()

	actualValues := exampleDomain.Enumerate()

	for _, actualValue := range actualValues {
		if !exampleDomain.Contains(actualValue) {
			t.Fatalf("actual value `%v` is not contained in domain", actualValue)
		}

		_, found := expectedValuesMap[actualValue]
		if !found {
			t.Fatalf("actual value `%v` was not in expected list", actualValue)
		}
	}

	actualValuesMap := actualValues.toValueMap()

	for expectedValue := range expectedValuesMap {
		_, found := actualValuesMap[expectedValue]
		if !found {
			t.Fatalf("expected value `%v` was not in actual list", expectedValue)
		}
	}
}

func TestImageDomain_IiiftilerExample(t *testing.T) {
	// docker run -v $PWD:$PWD --workdir=$PWD amazoncorretto java -jar iiif-tiler.jar 3888x2592.jpg
	// find out/3888x2592 -type f -name '*.jpg' | cut -d/ -f3-4 | sed -E 's#^full/#0,0,3888,2592/#' | pbcopy

	exampleDomain := NewImageDomain(iiifimageapi.ImageInformation{
		Width:  3888,
		Height: 2592,
		Sizes: []iiifimageapi.ImageInformationSize{
			{
				Width:  122,
				Height: 81,
			},
			{
				Width:  243,
				Height: 162,
			},
			{
				Width:  486,
				Height: 324,
			},
			{
				Width:  972,
				Height: 648,
			},
			{
				Width:  1944,
				Height: 1296,
			},
			{
				Width:  3888,
				Height: 2592,
			},
		},
		Tiles: []iiifimageapi.ImageInformationTile{
			{
				Width:        1024,
				ScaleFactors: []uint32{1, 2, 4, 8, 16, 32},
			},
		},
	})

	expectedValuesMap := ValueList{
		// iiif-tiler generates duplicate files (for compatibility?): "full" + "3888,"
		// ignore this "full" one
		// {[4]uint32{0, 0, 3888, 2592}, [2]uint32{3888, 2592}},
		{[4]uint32{0, 0, 3888, 2592}, [2]uint32{243}},
		{[4]uint32{0, 0, 3888, 2592}, [2]uint32{486}},
		{[4]uint32{0, 0, 3888, 2592}, [2]uint32{972}},
		{[4]uint32{0, 0, 3888, 2592}, [2]uint32{1944}},
		{[4]uint32{0, 0, 3888, 2592}, [2]uint32{3888}},
		{[4]uint32{0, 0, 3888, 2592}, [2]uint32{122}},
		{[4]uint32{2048, 0, 1024, 1024}, [2]uint32{1024}},
		{[4]uint32{2048, 2048, 1024, 544}, [2]uint32{1024}},
		{[4]uint32{2048, 0, 1840, 2048}, [2]uint32{920}},
		{[4]uint32{0, 2048, 2048, 544}, [2]uint32{1024}},
		{[4]uint32{1024, 0, 1024, 1024}, [2]uint32{1024}},
		{[4]uint32{0, 2048, 1024, 544}, [2]uint32{1024}},
		{[4]uint32{0, 0, 1024, 1024}, [2]uint32{1024}},
		{[4]uint32{3072, 2048, 816, 544}, [2]uint32{816}},
		{[4]uint32{0, 0, 3888, 2592}, [2]uint32{243}},
		{[4]uint32{0, 0, 3888, 2592}, [2]uint32{486}},
		{[4]uint32{0, 0, 3888, 2592}, [2]uint32{972}},
		{[4]uint32{0, 0, 3888, 2592}, [2]uint32{122}},
		{[4]uint32{1024, 2048, 1024, 544}, [2]uint32{1024}},
		{[4]uint32{2048, 2048, 1840, 544}, [2]uint32{920}},
		{[4]uint32{1024, 1024, 1024, 1024}, [2]uint32{1024}},
		{[4]uint32{3072, 1024, 816, 1024}, [2]uint32{816}},
		{[4]uint32{3072, 0, 816, 1024}, [2]uint32{816}},
		{[4]uint32{0, 1024, 1024, 1024}, [2]uint32{1024}},
		{[4]uint32{0, 0, 2048, 2048}, [2]uint32{1024}},
		{[4]uint32{2048, 1024, 1024, 1024}, [2]uint32{1024}},
	}.toValueMap()

	actualValues := exampleDomain.Enumerate()

	// reverse, but with ignored size height
	actualValuesMap := valueMap{}

	for _, actualValue := range actualValues {
		if !exampleDomain.Contains(actualValue) {
			t.Fatalf("actual value `%v` is not contained in domain", actualValue)
		}

		actualValue.Size[1] = 0 // libvips does not include height

		_, found := expectedValuesMap[actualValue]
		if !found {
			t.Fatalf("actual value `%v` was not in expected list", actualValue)
		}

		actualValuesMap[actualValue] = struct{}{}
	}

	for expectedValue := range expectedValuesMap {
		_, found := actualValuesMap[expectedValue]
		if !found {
			t.Fatalf("expected value `%v` was not in actual list", expectedValue)
		}
	}
}
