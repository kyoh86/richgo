package ptr

import (
	"testing"
	"time"
)

func TestUint8(t *testing.T) {
	var x uint8
	x = 9
	p := Uint8(x)
	if p == nil {
		t.Fatalf("failed to get pointer of Uint8, got nil")
	}
	if *p != x {
		t.Fatalf("failed to get pointer of Uint8, got invalid value: %d", *p)
	}
}

func TestUint16(t *testing.T) {
	var x uint16
	x = 9
	p := Uint16(x)
	if p == nil {
		t.Fatalf("failed to get pointer of Uint16, got nil")
	}
	if *p != x {
		t.Fatalf("failed to get pointer of Uint16, got invalid value: %d", *p)
	}
}

func TestUint32(t *testing.T) {
	var x uint32
	x = 9
	p := Uint32(x)
	if p == nil {
		t.Fatalf("failed to get pointer of Uint32, got nil")
	}
	if *p != x {
		t.Fatalf("failed to get pointer of Uint32, got invalid value: %d", *p)
	}
}

func TestUint64(t *testing.T) {
	var x uint64
	x = 9
	p := Uint64(x)
	if p == nil {
		t.Fatalf("failed to get pointer of Uint64, got nil")
	}
	if *p != x {
		t.Fatalf("failed to get pointer of Uint64, got invalid value: %d", *p)
	}
}

func TestInt8(t *testing.T) {
	var x int8
	x = 9
	p := Int8(x)
	if p == nil {
		t.Fatalf("failed to get pointer of Int8, got nil")
	}
	if *p != x {
		t.Fatalf("failed to get pointer of Int8, got invalid value: %d", *p)
	}
}

func TestInt16(t *testing.T) {
	var x int16
	x = 9
	p := Int16(x)
	if p == nil {
		t.Fatalf("failed to get pointer of Int16, got nil")
	}
	if *p != x {
		t.Fatalf("failed to get pointer of Int16, got invalid value: %d", *p)
	}
}

func TestInt32(t *testing.T) {
	var x int32
	x = 9
	p := Int32(x)
	if p == nil {
		t.Fatalf("failed to get pointer of Int32, got nil")
	}
	if *p != x {
		t.Fatalf("failed to get pointer of Int32, got invalid value: %d", *p)
	}
}

func TestInt64(t *testing.T) {
	var x int64
	x = 9
	p := Int64(x)
	if p == nil {
		t.Fatalf("failed to get pointer of Int64, got nil")
	}
	if *p != x {
		t.Fatalf("failed to get pointer of Int64, got invalid value: %d", *p)
	}
}

func TestFloat32(t *testing.T) {
	var x float32
	x = 9
	p := Float32(x)
	if p == nil {
		t.Fatalf("failed to get pointer of Float32, got nil")
	}
	if *p != x {
		t.Fatalf("failed to get pointer of Float32, got invalid value: %f", *p)
	}
}

func TestFloat64(t *testing.T) {
	var x float64
	x = 9
	p := Float64(x)
	if p == nil {
		t.Fatalf("failed to get pointer of Float64, got nil")
	}
	if *p != x {
		t.Fatalf("failed to get pointer of Float64, got invalid value: %f", *p)
	}
}

func TestComplex64(t *testing.T) {
	var x complex64
	x = 9
	p := Complex64(x)
	if p == nil {
		t.Fatalf("failed to get pointer of Complex64, got nil")
	}
	if *p != x {
		t.Fatalf("failed to get pointer of Complex64, got invalid value: %v", *p)
	}
}

func TestComplex128(t *testing.T) {
	var x complex128
	x = 9
	p := Complex128(x)
	if p == nil {
		t.Fatalf("failed to get pointer of Complex128, got nil")
	}
	if *p != x {
		t.Fatalf("failed to get pointer of Complex128, got invalid value: %v", *p)
	}
}

func TestByte(t *testing.T) {
	var x byte
	x = 9
	p := Byte(x)
	if p == nil {
		t.Fatalf("failed to get pointer of Byte, got nil")
	}
	if *p != x {
		t.Fatalf("failed to get pointer of Byte, got invalid value: %d", *p)
	}
}

func TestRune(t *testing.T) {
	var x rune
	x = 9
	p := Rune(x)
	if p == nil {
		t.Fatalf("failed to get pointer of Rune, got nil")
	}
	if *p != x {
		t.Fatalf("failed to get pointer of Rune, got invalid value: %d", *p)
	}
}

func TestUint(t *testing.T) {
	var x uint
	x = 9
	p := Uint(x)
	if p == nil {
		t.Fatalf("failed to get pointer of Uint, got nil")
	}
	if *p != x {
		t.Fatalf("failed to get pointer of Uint, got invalid value: %d", *p)
	}
}

func TestInt(t *testing.T) {
	var x int
	x = 9
	p := Int(x)
	if p == nil {
		t.Fatalf("failed to get pointer of Int, got nil")
	}
	if *p != x {
		t.Fatalf("failed to get pointer of Int, got invalid value: %d", *p)
	}
}

func TestString(t *testing.T) {
	var x string
	x = "test"
	p := String(x)
	if p == nil {
		t.Fatalf("failed to get pointer of String, got nil")
	}
	if *p != x {
		t.Fatalf("failed to get pointer of String, got invalid value: %s", *p)
	}
}

func TestBool(t *testing.T) {
	var x bool
	x = true
	p := Bool(x)
	if p == nil {
		t.Fatalf("failed to get pointer of Bool, got nil")
	}
	if *p != x {
		t.Fatalf("failed to get pointer of Bool, got invalid value: %v", *p)
	}
}

func TestTrue(t *testing.T) {
	p := True()
	if p == nil {
		t.Fatalf("failed to get pointer of True, got nil")
	}
	if *p != true {
		t.Fatalf("failed to get pointer of True, got invalid value: %v", *p)
	}
}

func TestFalse(t *testing.T) {
	p := False()
	if p == nil {
		t.Fatalf("failed to get pointer of False, got nil")
	}
	if *p != false {
		t.Fatalf("failed to get pointer of False, got invalid value: %v", *p)
	}
}

func TestTime(t *testing.T) {
	var x time.Time
	x = time.Now()
	p := Time(x)
	if p == nil {
		t.Fatalf("failed to get pointer of time.Time, got nil")
	}
	if *p != x {
		t.Fatalf("failed to get pointer of time.Time, got invalid value: %v", *p)
	}
}
