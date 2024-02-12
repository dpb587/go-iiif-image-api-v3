package imagerequest

import (
	"strings"
	"testing"
)

/*
 * Error handling
 */

// Region Percent

func TestParseRawParams_Region_ErrPctInvalidFormat(t *testing.T) {
	_, err := ParseRawParams(RawParams{"pct:invalid", "max", "0", "default.jpg"})
	if err == nil {
		t.Fatal("expected value but got `nil`")
	} else if _e, _a := "(expecting `pct:*,*,*,*`)", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	}
}

func TestParseRawParams_Region_ErrPctBadFloat(t *testing.T) {
	_, err := ParseRawParams(RawParams{"pct:0,1,2,invalid", "max", "0", "default.jpg"})
	if err == nil {
		t.Fatal("expected value but got `nil`")
	} else if _e, _a := "region: percents[3]", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	} else if _e, _a := "(expecting float, range [0, 100])", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	}
}

func TestParseRawParams_Region_ErrPctBadValueGt(t *testing.T) {
	_, err := ParseRawParams(RawParams{"pct:0,1,2,111", "max", "0", "default.jpg"})
	if err == nil {
		t.Fatal("expected value but got `nil`")
	} else if _e, _a := "region: percents[3]", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	} else if _e, _a := "(expecting float, range [0, 100])", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	}
}

func TestParseRawParams_Region_ErrPctBadValueLt(t *testing.T) {
	_, err := ParseRawParams(RawParams{"pct:0,1,2,-111", "max", "0", "default.jpg"})
	if err == nil {
		t.Fatal("expected value but got `nil`")
	} else if _e, _a := "region: percents[3]", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	} else if _e, _a := "(expecting float, range [0, 100])", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	}
}

// Region Pixels

func TestParseRawParams_Region_ErrPixelsInvalidFormat(t *testing.T) {
	_, err := ParseRawParams(RawParams{"invalid", "max", "0", "default.jpg"})
	if err == nil {
		t.Fatal("expected value but got `nil`")
	} else if _e, _a := "(expecting `*,*,*,*`)", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	}
}

func TestParseRawParams_Region_ErrPixelsBadFloat(t *testing.T) {
	_, err := ParseRawParams(RawParams{"0,1,2,3.3", "max", "0", "default.jpg"})
	if err == nil {
		t.Fatal("expected value but got `nil`")
	} else if _e, _a := "region: pixels[3]", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	} else if _e, _a := "(expecting integer, range [0, ))", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	}
}

func TestParseRawParams_Region_ErrPixelsBadValueLt(t *testing.T) {
	_, err := ParseRawParams(RawParams{"0,1,2,-111", "max", "0", "default.jpg"})
	if err == nil {
		t.Fatal("expected value but got `nil`")
	} else if _e, _a := "region: pixels[3]", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	} else if _e, _a := "(expecting integer, range [0, ))", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	}
}

// Size Percent

func TestParseRawParams_Size_ErrPctBadFloat(t *testing.T) {
	_, err := ParseRawParams(RawParams{"full", "pct:invalid.invalid", "0", "default.jpg"})
	if err == nil {
		t.Fatal("expected value but got `nil`")
	} else if _e, _a := "size: percent", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	} else if _e, _a := "(expecting float)", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	}
}

func TestParseRawParams_Size_ErrPctBadValueGt(t *testing.T) {
	_, err := ParseRawParams(RawParams{"full", "pct:111", "0", "default.jpg"})
	if err == nil {
		t.Fatal("expected value but got `nil`")
	} else if _e, _a := "size: percent", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	} else if _e, _a := "(expecting range (0, 100] for non-upscaled)", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	}
}

func TestParseRawParams_Size_ErrPctBadValueLt(t *testing.T) {
	_, err := ParseRawParams(RawParams{"full", "pct:-111", "0", "default.jpg"})
	if err == nil {
		t.Fatal("expected value but got `nil`")
	} else if _e, _a := "size: percent", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	} else if _e, _a := "(expecting range (0, ))", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	}
}

// Size Pixels

func TestParseRawParams_Size_ErrPixelsInvalidFormat(t *testing.T) {
	_, err := ParseRawParams(RawParams{"full", "invalid", "0", "default.jpg"})
	if err == nil {
		t.Fatal("expected value but got `nil`")
	} else if _e, _a := "(expecting `*,*`)", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	}
}

func TestParseRawParams_Size_ErrPixelsBadFloat(t *testing.T) {
	_, err := ParseRawParams(RawParams{"full", ",3.3", "0", "default.jpg"})
	if err == nil {
		t.Fatal("expected value but got `nil`")
	} else if _e, _a := "size: pixels[1]", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	} else if _e, _a := "(expecting int)", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	}
}

func TestParseRawParams_Size_ErrPixelsBadValueLt(t *testing.T) {
	_, err := ParseRawParams(RawParams{"full", ",0", "0", "default.jpg"})
	if err == nil {
		t.Fatal("expected value but got `nil`")
	} else if _e, _a := "size: pixels[1]", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	} else if _e, _a := "(expecting range (0, ))", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	}
}

// Rotation

func TestParseRawParams_Rotation_ErrBadFloat(t *testing.T) {
	_, err := ParseRawParams(RawParams{"full", "max", "invalid.invalid", "default.jpg"})
	if err == nil {
		t.Fatal("expected value but got `nil`")
	} else if _e, _a := "rotation", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	} else if _e, _a := "(expecting float or int)", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	}
}

func TestParseRawParams_Rotation_ErrBadFloatGt(t *testing.T) {
	_, err := ParseRawParams(RawParams{"full", "max", "630.9", "default.jpg"})
	if err == nil {
		t.Fatal("expected value but got `nil`")
	} else if _e, _a := "rotation", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	} else if _e, _a := "(expecting range [0, 360])", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	}
}

func TestParseRawParams_Rotation_ErrBadFloatLt(t *testing.T) {
	_, err := ParseRawParams(RawParams{"full", "max", "-630.9", "default.jpg"})
	if err == nil {
		t.Fatal("expected value but got `nil`")
	} else if _e, _a := "rotation", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	} else if _e, _a := "(expecting range [0, 360])", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	}
}

func TestParseRawParams_Rotation_ErrBadInt(t *testing.T) {
	_, err := ParseRawParams(RawParams{"full", "max", "invalid", "default.jpg"})
	if err == nil {
		t.Fatal("expected value but got `nil`")
	} else if _e, _a := "rotation", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	} else if _e, _a := "(expecting float or int)", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	}
}

func TestParseRawParams_Rotation_ErrBadIntGt(t *testing.T) {
	_, err := ParseRawParams(RawParams{"full", "max", "630", "default.jpg"})
	if err == nil {
		t.Fatal("expected value but got `nil`")
	} else if _e, _a := "rotation", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	} else if _e, _a := "(expecting range [0, 360])", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	}
}

func TestParseRawParams_Rotation_ErrBadIntLt(t *testing.T) {
	_, err := ParseRawParams(RawParams{"full", "max", "-630", "default.jpg"})
	if err == nil {
		t.Fatal("expected value but got `nil`")
	} else if _e, _a := "rotation", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	} else if _e, _a := "(expecting range [0, 360])", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	}
}

// File

func TestParseRawParams_File_ErrInvalidFormat(t *testing.T) {
	_, err := ParseRawParams(RawParams{"full", "max", "0", "invalid"})
	if err == nil {
		t.Fatal("expected value but got `nil`")
	} else if _e, _a := "file", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	} else if _e, _a := "(expecting `*.*`)", err.Error(); !strings.Contains(_a, _e) {
		t.Fatalf("expected `%v` to contain `%v`", _a, _e)
	}
}
