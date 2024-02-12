package imagerequest

import "testing"

func TestMaxConstraintNone(t *testing.T) {
	res := maxConstraint{}.Constrain([2]uint32{4096, 8192}, false)

	if _e, _a := [2]uint32{4096, 8192}, res; _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}
}

func TestMaxConstraintWidth(t *testing.T) {
	res := maxConstraint{
		maxWidth: ptrUint32(2048),
	}.Constrain([2]uint32{4096, 8192}, false)

	if _e, _a := [2]uint32{2048, 4096}, res; _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}
}

func TestMaxConstraintWidthUpscale(t *testing.T) {
	res := maxConstraint{
		maxWidth: ptrUint32(16384),
	}.Constrain([2]uint32{4096, 8192}, true)

	if _e, _a := [2]uint32{16384, 32768}, res; _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}
}

func TestMaxConstraintHeight(t *testing.T) {
	res := maxConstraint{
		maxHeight: ptrUint32(2048),
	}.Constrain([2]uint32{4096, 8192}, false)

	if _e, _a := [2]uint32{1024, 2048}, res; _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}
}

func TestMaxConstraintHeightUpscale(t *testing.T) {
	res := maxConstraint{
		maxHeight: ptrUint32(16384),
	}.Constrain([2]uint32{4096, 8192}, true)

	if _e, _a := [2]uint32{8192, 16384}, res; _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}
}

func TestMaxConstraintWidthHeight(t *testing.T) {
	res := maxConstraint{
		maxWidth:  ptrUint32(1024),
		maxHeight: ptrUint32(1024),
	}.Constrain([2]uint32{4096, 8192}, true)

	if _e, _a := [2]uint32{512, 1024}, res; _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}
}

func TestMaxConstraintWidthHeightFlip(t *testing.T) {
	res := maxConstraint{
		maxWidth:  ptrUint32(1024),
		maxHeight: ptrUint32(1024),
	}.Constrain([2]uint32{8192, 4096}, true)

	if _e, _a := [2]uint32{1024, 512}, res; _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}
}

func TestMaxConstraintWidthHeight1(t *testing.T) {
	res := maxConstraint{
		maxWidth:  ptrUint32(2048),
		maxHeight: ptrUint32(1024),
	}.Constrain([2]uint32{4096, 8192}, true)

	if _e, _a := [2]uint32{512, 1024}, res; _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}
}

func TestMaxConstraintWidthHeight1Flip(t *testing.T) {
	res := maxConstraint{
		maxWidth:  ptrUint32(2048),
		maxHeight: ptrUint32(1024),
	}.Constrain([2]uint32{8192, 4096}, true)

	if _e, _a := [2]uint32{2048, 1024}, res; _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}
}

func TestMaxConstraintArea(t *testing.T) {
	res := maxConstraint{
		maxArea: ptrUint32(4096 * 4096),
	}.Constrain([2]uint32{4096, 8192}, false)

	if _e, _a := [2]uint32{2896, 5792}, res; _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}
}

func TestMaxConstraintAreaUpscale(t *testing.T) {
	/*
		67108864 = (d * 4096) * (d * 8192)
		67108864 = d * 4096 * d * 8192
		67108864 = d^2 * 4096 * 8192
		67108864/8192 = d^2 * 4096
		67108864/8192/4096 = d^2
		sqrt(67108864/8192/4096) = d
	*/

	res := maxConstraint{
		maxArea: ptrUint32(8192 * 8192),
	}.Constrain([2]uint32{4096, 8192}, true)

	if _e, _a := [2]uint32{5792, 11585}, res; _e != _a {
		t.Fatalf("expected `%v` but got: %v", _e, _a)
	}
}
