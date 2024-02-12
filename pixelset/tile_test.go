package pixelset

import "testing"

func TestGetTileScaleFactorMax_Common(t *testing.T) {
	m := GetTileScaleFactorMax([2]uint32{3888, 2592}, [2]uint32{512, 512})
	if _e, _a := 8, m; _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}
}

func TestGetTileScaleFactorMax_Mismatch(t *testing.T) {
	m := GetTileScaleFactorMax([2]uint32{3888, 2592}, [2]uint32{256, 512})
	if _e, _a := 16, m; _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}
}

func TestGetTileScaleFactorMax_Lesser(t *testing.T) {
	m := GetTileScaleFactorMax([2]uint32{388, 259}, [2]uint32{512, 512})
	if _e, _a := 1, m; _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}
}
