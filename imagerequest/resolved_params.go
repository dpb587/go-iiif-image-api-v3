package imagerequest

// ResolvedParams is the result of [ParsedParams.Resolve] for a particular [spec.ImageInformation]. It contains absolute
// pixel dimensions and is the minimal set of information required by an image processor.
type ResolvedParams struct {
	regionPixels       [4]uint32
	sizePixels         [2]uint32
	rotationIsMirrored bool
	rotationAmount     float32
	quality            string
	format             string

	canonical ParsedParams
}

// RegionPixels returns the pixel set of [X, Y, Width, Height].
func (p ResolvedParams) RegionPixels() [4]uint32 {
	return p.regionPixels
}

// SizePixels returns the pixel set of [Width, Height].
func (p ResolvedParams) SizePixels() [2]uint32 {
	return p.sizePixels
}

// RotationIsMirrored returns true if it should be mirrored acorss the vertical axis before rotation.
func (p ResolvedParams) RotationIsMirrored() bool {
	return p.rotationIsMirrored
}

// RotationAmount returns the amount an image should be rotated from 0 (inclusive) to 360 (exclusive).
func (p ResolvedParams) RotationAmount() float32 {
	return p.rotationAmount
}

// Quality is the literal quality to render (not "default").
func (p ResolvedParams) Quality() string {
	return p.quality
}

// Format is the requested format (e.g. "jpg").
func (p ResolvedParams) Format() string {
	return p.format
}

// ToParsedParams returns the parameters in expanded form.
func (p ResolvedParams) ToParsedParams() ParsedParams {
	return ParsedParams{
		RegionPixels:       p.regionPixels,
		SizePixels:         [2]*uint32{&p.sizePixels[0], &p.sizePixels[1]},
		RotationIsMirrored: p.rotationIsMirrored,
		RotationAmount:     p.rotationAmount,
		Quality:            p.quality,
		Format:             p.format,
	}
}

// Canonical returns the expanded form according to the original resolver options.
func (p ResolvedParams) Canonical() ParsedParams {
	return p.canonical.Clone()
}

// ResolveCanonical returns the expanded form according to the given opts.
func (p ResolvedParams) ResolveCanonical(opts ResolveOptions) ParsedParams {
	info := opts.ImageInformation
	canonical := p.ToParsedParams()

	{ // region
		if p.regionPixels[0] == 0 && p.regionPixels[1] == 0 && p.regionPixels[2] == info.Width && p.regionPixels[3] == info.Height {
			canonical.RegionIsEnum = true
			canonical.RegionEnum = "full"
			canonical.RegionPixels = [4]uint32{}
		}
	}

	{ // size
		if p.sizePixels[0] > p.regionPixels[2] || p.sizePixels[1] > p.regionPixels[3] {
			canonical.SizeIsUpscaled = true
		}

		if info.MaxWidth != nil && info.MaxHeight != nil && *info.MaxWidth == p.sizePixels[0] && *info.MaxHeight == p.sizePixels[1] {
			canonical.SizeIsEnum = true
			canonical.SizeEnum = "max"
			canonical.SizePixels = [2]*uint32{}
		}

		// TODO area
	}

	{ // quality
		if p.quality == opts.DefaultQuality {
			canonical.Quality = "default"
		}
	}

	return canonical
}
