package validation

import "testing"

type minValueTestType struct {
	Int8   int8   `validation:"min=-20"`
	Uint8  uint8  `validation:min=20`
	Int16  int16  `validation:"min=-20"`
	Uint16 uint16 `validation:min=20`
	Int32  int32  `validation:"min=-20"`
	Uint32 uint32 `validation:min=20`
	Int64  int64  `validation:min=-20`
	Uint64 uint64 `validation:min=20`
	Int    int    `validation:min=-20`
	Uint   uint   `validation:min=20`

	// Floats
	Float32 float32 `validation:"min=-20"`
	Float64 float64 `validation:"min=-20"`
}

func TestMinValueInt8Negative(t *testing.T) {
	type minValueTestType struct {
		Value int8 `validation:"min=-20"`
	}
	obj := minValueTestType{}

	ok, _ := IsValid(obj)

	if !ok {
		t.Fatal("Empty Int8(0) should be valid (>= -20)")
	}

	obj.Value = -40

	ok, _ = IsValid(obj)

	if ok {
		t.Fatal("Expected failure as -40 is less than min -20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)

	if !ok {
		t.Fatal("Valid: 40 is greater than -20", errs)
	}
}

func TestMinValueInt16Negative(t *testing.T) {
	type minValueTestType struct {
		Value int16 `validation:"min=-20"`
	}
	obj := minValueTestType{}

	ok, _ := IsValid(obj)

	if !ok {
		t.Fatal("Empty Int8(0) should be valid (>= -20)")
	}

	obj.Value = -40

	ok, _ = IsValid(obj)

	if ok {
		t.Fatal("Expected failure as -40 is less than min -20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)

	if !ok {
		t.Fatal("Valid: 40 is greater than -20", errs)
	}
}

func TestMinValueInt32Negative(t *testing.T) {
	type minValueTestType struct {
		Value int32 `validation:"min=-20"`
	}
	obj := minValueTestType{}

	ok, _ := IsValid(obj)

	if !ok {
		t.Fatal("Empty Int8(0) should be valid (>= -20)")
	}

	obj.Value = -40

	ok, _ = IsValid(obj)

	if ok {
		t.Fatal("Expected failure as -40 is less than min -20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)

	if !ok {
		t.Fatal("Valid: 40 is greater than -20", errs)
	}
}

func TestMinValueInt64Negative(t *testing.T) {
	type minValueTestType struct {
		Value int64 `validation:"min=-20"`
	}
	obj := minValueTestType{}

	ok, _ := IsValid(obj)

	if !ok {
		t.Fatal("Empty Int8(0) should be valid (>= -20)")
	}

	obj.Value = -40

	ok, _ = IsValid(obj)

	if ok {
		t.Fatal("Expected failure as -40 is less than min -20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)

	if !ok {
		t.Fatal("Valid: 40 is greater than -20", errs)
	}
}

func TestMinValueIntNegative(t *testing.T) {
	type minValueTestType struct {
		Value int `validation:"min=-20"`
	}
	obj := minValueTestType{}

	ok, _ := IsValid(obj)

	if !ok {
		t.Fatal("Empty Int8(0) should be valid (>= -20)")
	}

	obj.Value = -40

	ok, _ = IsValid(obj)

	if ok {
		t.Fatal("Expected failure as -40 is less than min -20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)

	if !ok {
		t.Fatal("Valid: 40 is greater than -20", errs)
	}
}

// Uint test
func TestMinValueUint8Negative(t *testing.T) {
	type minValueTestType struct {
		Value uint8 `validation:"min=20"`
	}
	obj := minValueTestType{}

	ok, _ := IsValid(obj)

	if ok {
		t.Fatal("Expected failure as 0 is less than min 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)

	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

func TestMinValueUint16Negative(t *testing.T) {
	type minValueTestType struct {
		Value uint16 `validation:"min=20"`
	}
	obj := minValueTestType{}

	ok, _ := IsValid(obj)

	if ok {
		t.Fatal("Expected failure as 0 is less than min 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)

	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

func TestMinValueUint32Negative(t *testing.T) {
	type minValueTestType struct {
		Value uint32 `validation:"min=20"`
	}
	obj := minValueTestType{}

	ok, _ := IsValid(obj)

	if ok {
		t.Fatal("Expected failure as 0 is less than min 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)

	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

func TestMinValueUint64Negative(t *testing.T) {
	type minValueTestType struct {
		Value uint64 `validation:"min=20"`
	}
	obj := minValueTestType{}

	ok, _ := IsValid(obj)

	if ok {
		t.Fatal("Expected failure as 0 is less than min 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)

	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

func TestMinValueUintNegative(t *testing.T) {
	type minValueTestType struct {
		Value uint `validation:"min=20"`
	}
	obj := minValueTestType{}

	ok, _ := IsValid(obj)

	if ok {
		t.Fatal("Expected failure as 0 is less than min 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)

	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

// Floats

func TestMinValueFloat32(t *testing.T) {
	type minValueTestType struct {
		Value float32 `validation:"min=20"`
	}
	obj := minValueTestType{}

	ok, _ := IsValid(obj)

	if ok {
		t.Fatal("Expected failure as 0 is less than min 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)

	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

func TestMinValueFloat64(t *testing.T) {
	type minValueTestType struct {
		Value float64 `validation:"min=20"`
	}
	obj := minValueTestType{}

	ok, _ := IsValid(obj)

	if ok {
		t.Fatal("Expected failure as 0 is less than min 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)

	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

// Max Value
func TestMaxValueInt8Negative(t *testing.T) {
	type maxValueTestType struct {
		Value int8 `validation:"max=-20"`
	}
	obj := maxValueTestType{}

	ok, _ := IsValid(obj)

	if ok {
		t.Fatal("Empty Int8(0) should be invalid (>= -20)")
	}

	obj.Value = -40

	ok, errs := IsValid(obj)

	if !ok {
		t.Fatal("Expected valid as -40 is less than max -20", errs)
	}

}

func TestMaxValueInt16Negative(t *testing.T) {
	type maxValueTestType struct {
		Value int16 `validation:"max=-20"`
	}
	obj := maxValueTestType{}

	ok, _ := IsValid(obj)

	if ok {
		t.Fatal("Empty Int16(0) should be invalid (>= -20)")
	}

	obj.Value = -40

	ok, errs := IsValid(obj)

	if !ok {
		t.Fatal("Expected valid as -40 is less than max -20", errs)
	}

}

func TestMaxValueInt32Negative(t *testing.T) {
	type maxValueTestType struct {
		Value int32 `validation:"max=-20"`
	}
	obj := maxValueTestType{}

	ok, _ := IsValid(obj)

	if ok {
		t.Fatal("Empty Int32(0) should be invalid (>= -20)")
	}

	obj.Value = -40

	ok, errs := IsValid(obj)

	if !ok {
		t.Fatal("Expected valid as -40 is less than max -20", errs)
	}
}

func TestMaxValueInt64Negative(t *testing.T) {
	type maxValueTestType struct {
		Value int64 `validation:"max=-20"`
	}
	obj := maxValueTestType{}

	ok, _ := IsValid(obj)

	if ok {
		t.Fatal("Empty Int64(0) should be invalid (>= -20)")
	}

	obj.Value = -40

	ok, errs := IsValid(obj)

	if !ok {
		t.Fatal("Expected valid as -40 is less than max -20", errs)
	}
}

func TestMaxValueIntNegative(t *testing.T) {
	type maxValueTestType struct {
		Value int `validation:"max=-20"`
	}
	obj := maxValueTestType{}

	ok, _ := IsValid(obj)

	if ok {
		t.Fatal("Empty Int(0) should be invalid (>= -20)")
	}

	obj.Value = -40

	ok, errs := IsValid(obj)

	if !ok {
		t.Fatal("Expected valid as -40 is less than max -20", errs)
	}
}

// Uint test
func TestMaxValueUint8Negative(t *testing.T) {
	type maxValueTestType struct {
		Value uint8 `validation:"max=20"`
	}
	obj := maxValueTestType{}

	ok, _ := IsValid(obj)

	if !ok {
		t.Fatal("Expected success as 0 is less than max 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)

	if ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

func TestMaxValueUint16Negative(t *testing.T) {
	type maxValueTestType struct {
		Value uint16 `validation:"max=20"`
	}
	obj := maxValueTestType{}

	ok, _ := IsValid(obj)

	if !ok {
		t.Fatal("Expected success as 0 is less than max 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)

	if ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

func TestMaxValueUint32Negative(t *testing.T) {
	type maxValueTestType struct {
		Value uint32 `validation:"max=20"`
	}
	obj := maxValueTestType{}

	ok, _ := IsValid(obj)

	if !ok {
		t.Fatal("Expected success as 0 is less than max 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)

	if ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

func TestMaxValueUint64Negative(t *testing.T) {
	type maxValueTestType struct {
		Value uint64 `validation:"max=20"`
	}
	obj := maxValueTestType{}

	ok, _ := IsValid(obj)

	if !ok {
		t.Fatal("Expected success as 0 is less than max 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)

	if ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

func TestMaxValueUintNegative(t *testing.T) {
	type maxValueTestType struct {
		Value uint `validation:"max=20"`
	}
	obj := maxValueTestType{}

	ok, _ := IsValid(obj)

	if !ok {
		t.Fatal("Expected success as 0 is less than max 20")
	}

	obj.Value = 40

	ok, errs := IsValid(obj)

	if ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

// Floats

func TestMaxValueFloat32(t *testing.T) {
	type maxValueTestType struct {
		Value float32 `validation:"max=-20"`
	}
	obj := maxValueTestType{}

	ok, _ := IsValid(obj)

	if ok {
		t.Fatal("Expected failure as 0 is less than max -20")
	}

	obj.Value = -40

	ok, errs := IsValid(obj)

	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}

func TestMaxValueFloat64(t *testing.T) {
	type maxValueTestType struct {
		Value float64 `validation:"max=-20"`
	}
	obj := maxValueTestType{}

	ok, _ := IsValid(obj)

	if ok {
		t.Fatal("Expected failure as 0 is less than max -20")
	}

	obj.Value = -40

	ok, errs := IsValid(obj)

	if !ok {
		t.Fatal("Valid: 40 is greater than 20", errs)
	}
}
