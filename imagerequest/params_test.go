package imagerequest

import (
	"strings"
	"testing"
)

func TestRawParamsFromString_ErrInvalidFormatShort(t *testing.T) {
	_, err := RawParamsFromString("full/max/0")
	if err == nil {
		t.Fatal("expected value but got `nil`")
	} else if _e, _a := "(expecting `*/*/*/*`)", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	}
}

func TestRawParamsFromString_ErrInvalidFormatLong(t *testing.T) {
	_, err := RawParamsFromString("identifier/full/max/0/default.png")
	if err == nil {
		t.Fatal("expected value but got `nil`")
	} else if _e, _a := "(expecting `*/*/*/*`)", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	}
}

func TestRawParamsFromString_ErrInvalidEncoding(t *testing.T) {
	_, err := RawParamsFromString("full/max/%%%2/default.png")
	if err == nil {
		t.Fatal("expected value but got `nil`")
	} else if _e, _a := "(path[2])", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	} else if _e, _a := "invalid encoding", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	}
}
