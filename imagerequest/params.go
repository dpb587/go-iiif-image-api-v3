package imagerequest

import (
	"fmt"
	"net/url"
	"strings"

	iiifimageapi "github.com/dpb587/go-iiif-image-api-v3"
)

// RawParams represents the raw [{region}, {size}, {rotation}, {quality.format}] values described in the specification.
type RawParams [4]string

func (p RawParams) String() string {
	return url.PathEscape(p[0]) + "/" + url.PathEscape(p[1]) + "/" + url.PathEscape(p[2]) + "/" + url.PathEscape(p[3])
}

// RawParamsFromString decodes path segments from a `{region}/{size}/{rotation}/{quality}.{format}` string.
func RawParamsFromString(raw string) (RawParams, error) {
	pathSplit := strings.SplitN(raw, "/", 5)
	if len(pathSplit) != 4 {
		return RawParams{}, iiifimageapi.NewInvalidValueError("invalid path format (expecting `*/*/*/*`)")
	}

	var res RawParams

	for k, v := range pathSplit {
		vu, err := url.PathUnescape(v)
		if err != nil {
			return RawParams{}, iiifimageapi.NewInvalidValueError(fmt.Sprintf("parsing (path[%d]): invalid encoding", k))
		}

		res[k] = vu
	}

	return res, nil
}
