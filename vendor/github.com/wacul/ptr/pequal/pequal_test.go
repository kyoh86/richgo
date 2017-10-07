package pequal

import (
	"testing"
	"time"

	"github.com/smartystreets/assertions"
)

func TestUint8(t *testing.T) {
	var x uint8 = 9
	var y uint8 = 8
	if ok, message := assertions.So(Uint8(nil, nil), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Uint8(nil, &x), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Uint8(&x, nil), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Uint8(&x, &x), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Uint8(&x, &y), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
}

func TestUint16(t *testing.T) {
	var x uint16 = 9
	var y uint16 = 8
	if ok, message := assertions.So(Uint16(nil, nil), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Uint16(nil, &x), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Uint16(&x, nil), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Uint16(&x, &x), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Uint16(&x, &y), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
}

func TestUint32(t *testing.T) {
	var x uint32 = 9
	var y uint32 = 8
	if ok, message := assertions.So(Uint32(nil, nil), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Uint32(nil, &x), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Uint32(&x, nil), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Uint32(&x, &x), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Uint32(&x, &y), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
}

func TestUint64(t *testing.T) {
	var x uint64 = 9
	var y uint64 = 8
	if ok, message := assertions.So(Uint64(nil, nil), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Uint64(nil, &x), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Uint64(&x, nil), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Uint64(&x, &x), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Uint64(&x, &y), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
}

func TestInt8(t *testing.T) {
	var x int8 = 9
	var y int8 = 8
	if ok, message := assertions.So(Int8(nil, nil), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Int8(nil, &x), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Int8(&x, nil), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Int8(&x, &x), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Int8(&x, &y), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
}

func TestInt16(t *testing.T) {
	var x int16 = 9
	var y int16 = 8
	if ok, message := assertions.So(Int16(nil, nil), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Int16(nil, &x), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Int16(&x, nil), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Int16(&x, &x), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Int16(&x, &y), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
}

func TestInt32(t *testing.T) {
	var x int32 = 9
	var y int32 = 8
	if ok, message := assertions.So(Int32(nil, nil), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Int32(nil, &x), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Int32(&x, nil), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Int32(&x, &x), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Int32(&x, &y), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
}

func TestInt64(t *testing.T) {
	var x int64 = 9
	var y int64 = 8
	if ok, message := assertions.So(Int64(nil, nil), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Int64(nil, &x), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Int64(&x, nil), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Int64(&x, &x), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Int64(&x, &y), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
}

func TestFloat32(t *testing.T) {
	var x float32 = 9
	var y float32 = 8
	if ok, message := assertions.So(Float32(nil, nil), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Float32(nil, &x), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Float32(&x, nil), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Float32(&x, &x), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Float32(&x, &y), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
}

func TestFloat64(t *testing.T) {
	var x float64 = 9
	var y float64 = 8
	if ok, message := assertions.So(Float64(nil, nil), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Float64(nil, &x), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Float64(&x, nil), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Float64(&x, &x), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Float64(&x, &y), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
}

func TestComplex64(t *testing.T) {
	var x complex64 = 9
	var y complex64 = 8
	if ok, message := assertions.So(Complex64(nil, nil), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Complex64(nil, &x), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Complex64(&x, nil), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Complex64(&x, &x), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Complex64(&x, &y), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
}

func TestComplex128(t *testing.T) {
	var x complex128 = 9
	var y complex128 = 8
	if ok, message := assertions.So(Complex128(nil, nil), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Complex128(nil, &x), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Complex128(&x, nil), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Complex128(&x, &x), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Complex128(&x, &y), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
}

func TestByte(t *testing.T) {
	var x byte = 9
	var y byte = 8
	if ok, message := assertions.So(Byte(nil, nil), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Byte(nil, &x), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Byte(&x, nil), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Byte(&x, &x), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Byte(&x, &y), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
}

func TestRune(t *testing.T) {
	var x rune = 9
	var y rune = 8
	if ok, message := assertions.So(Rune(nil, nil), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Rune(nil, &x), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Rune(&x, nil), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Rune(&x, &x), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Rune(&x, &y), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
}

func TestUint(t *testing.T) {
	var x uint = 9
	var y uint = 8
	if ok, message := assertions.So(Uint(nil, nil), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Uint(nil, &x), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Uint(&x, nil), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Uint(&x, &x), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Uint(&x, &y), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
}

func TestInt(t *testing.T) {
	var x int = 9
	var y int = 8
	if ok, message := assertions.So(Int(nil, nil), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Int(nil, &x), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Int(&x, nil), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Int(&x, &x), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Int(&x, &y), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
}

func TestString(t *testing.T) {
	var x string = "test"
	var y string = "another"
	if ok, message := assertions.So(String(nil, nil), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(String(nil, &x), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(String(&x, nil), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(String(&x, &x), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(String(&x, &y), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
}

func TestBool(t *testing.T) {
	var x bool = true
	var y bool = false
	if ok, message := assertions.So(Bool(nil, nil), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Bool(nil, &x), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Bool(&x, nil), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Bool(&x, &x), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Bool(&x, &y), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
}

func TestTime(t *testing.T) {
	var x time.Time = time.Now()
	var y time.Time = x.Add(5 * time.Minute)
	if ok, message := assertions.So(Time(nil, nil), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Time(nil, &x), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Time(&x, nil), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Time(&x, &x), assertions.ShouldBeTrue); !ok {
		t.Fatal(message)
	}
	if ok, message := assertions.So(Time(&x, &y), assertions.ShouldBeFalse); !ok {
		t.Fatal(message)
	}
}
