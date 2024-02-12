package iiifimageapi

import "fmt"

// FeatureNotSupportedError indicates an input constraint would have required a feature that is not supported.
type FeatureNotSupportedError FeatureName

func (e FeatureNotSupportedError) Error() string {
	return fmt.Sprintf("feature not supported: %s", string(e))
}

//

// InvalidValueError indicates a problem with an input constraint. For HTTP runtimes, this translates to an HTTP 400 Bad
// Request.
type InvalidValueError struct {
	s string
}

func NewInvalidValueError(s string) InvalidValueError {
	return InvalidValueError{
		s: s,
	}
}

func (e InvalidValueError) Error() string {
	return e.s
}
