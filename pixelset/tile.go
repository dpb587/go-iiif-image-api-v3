package pixelset

import "math"

func GetTileScaleFactorMax(imageSize [2]uint32, tileSize [2]uint32) int {
	wsf := math.Pow(2, math.Ceil(math.Log2(float64(imageSize[0])/float64(tileSize[0]))))
	hsf := math.Pow(2, math.Ceil(math.Log2(float64(imageSize[1])/float64(tileSize[1]))))

	return int(math.Max(wsf, hsf))
}

func GetTileScaleFactors(imageSize [2]uint32, tileSize [2]uint32) []uint32 {
	var out []uint32

	max := GetTileScaleFactorMax(imageSize, tileSize)

	for i := 1; i <= max; i *= 2 {
		out = append(out, uint32(i))
	}

	return out
}
