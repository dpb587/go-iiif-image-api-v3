A package focused on low-level data types and behaviors for the [IIIF Image API (v3)](https://iiif.io/api/image/3.0/) specification.

* parsing `{region}/{size}/{rotation}/{quality}.{format}` parameter values;
* validating parameters against advertised compliance levels, maximums, and extra features;
* resolving relative parameter values against an image profile for absolute pixels;
* converting parameter values to canonical form;
* enumerating region+size based on static size or tile configuration; and
* constructing basic `info.json` contents.

This package does not perform image processing. Resolved parameters should typically be used for performing upstream or RPC requests to dedicated image servers.

Learn more from [code documentation](https://pkg.go.dev/github.com/dpb587/go-iiif-image-api-v3), [`examples`](examples), or `*_test.go` files.

# Example

The following example is based on [`examples/parse/main.go`](examples/parse/main.go). Import the packages...

```go
import (
	iiifimageapi "github.com/dpb587/go-iiif-image-api-v3"
	"github.com/dpb587/go-iiif-image-api-v3/imagerequest"
)
```

And then use any of the functions necessary to process the request...

```go
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
```

# License

[MIT License](LICENSE)
