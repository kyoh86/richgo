// Package pequal provides utility functions to compare two pointers of the built-in type
package pequal

import "time"

// Uint8 compares two uint8 equalities
func Uint8(u1 *uint8, u2 *uint8) bool {
	if u1 == nil {
		return u2 == nil
	}
	return u2 != nil && *u1 == *u2
}

// Uint16 compares two uint16 equalities
func Uint16(u1 *uint16, u2 *uint16) bool {
	if u1 == nil {
		return u2 == nil
	}
	return u2 != nil && *u1 == *u2
}

// Uint32 compares two uint32 equalities
func Uint32(u1 *uint32, u2 *uint32) bool {
	if u1 == nil {
		return u2 == nil
	}
	return u2 != nil && *u1 == *u2
}

// Uint64 compares two uint64 equalities
func Uint64(u1 *uint64, u2 *uint64) bool {
	if u1 == nil {
		return u2 == nil
	}
	return u2 != nil && *u1 == *u2
}

// Int8 compares two int8 equalities
func Int8(i1 *int8, i2 *int8) bool {
	if i1 == nil {
		return i2 == nil
	}
	return i2 != nil && *i1 == *i2
}

// Int16 compares two int16 equalities
func Int16(i1 *int16, i2 *int16) bool {
	if i1 == nil {
		return i2 == nil
	}
	return i2 != nil && *i1 == *i2
}

// Int32 compares two int32 equalities
func Int32(i1 *int32, i2 *int32) bool {
	if i1 == nil {
		return i2 == nil
	}
	return i2 != nil && *i1 == *i2
}

// Int64 compares two int64 equalities
func Int64(i1 *int64, i2 *int64) bool {
	if i1 == nil {
		return i2 == nil
	}
	return i2 != nil && *i1 == *i2
}

// Float32 compares two float32 equalities
func Float32(f1 *float32, f2 *float32) bool {
	if f1 == nil {
		return f2 == nil
	}
	return f2 != nil && *f1 == *f2
}

// Float64 compares two float64 equalities
func Float64(f1 *float64, f2 *float64) bool {
	if f1 == nil {
		return f2 == nil
	}
	return f2 != nil && *f1 == *f2
}

// Complex64 compares two complex64 equalities
func Complex64(c1 *complex64, c2 *complex64) bool {
	if c1 == nil {
		return c2 == nil
	}
	return c2 != nil && *c1 == *c2
}

// Complex128 compares two complex128 equalities
func Complex128(c1 *complex128, c2 *complex128) bool {
	if c1 == nil {
		return c2 == nil
	}
	return c2 != nil && *c1 == *c2
}

// Byte compares two byte equalities
func Byte(b1 *byte, b2 *byte) bool {
	if b1 == nil {
		return b2 == nil
	}
	return b2 != nil && *b1 == *b2
}

// Rune compares two rune equalities
func Rune(r1 *rune, r2 *rune) bool {
	if r1 == nil {
		return r2 == nil
	}
	return r2 != nil && *r1 == *r2
}

// Uint compares two uint equalities
func Uint(u1 *uint, u2 *uint) bool {
	if u1 == nil {
		return u2 == nil
	}
	return u2 != nil && *u1 == *u2
}

// Int compares two int equalities
func Int(i1 *int, i2 *int) bool {
	if i1 == nil {
		return i2 == nil
	}
	return i2 != nil && *i1 == *i2
}

// String compares two string equalities
func String(s1 *string, s2 *string) bool {
	if s1 == nil {
		return s2 == nil
	}
	return s2 != nil && *s1 == *s2
}

// Bool compares two bool equalities
func Bool(b1 *bool, b2 *bool) bool {
	if b1 == nil {
		return b2 == nil
	}
	return b2 != nil && *b1 == *b2
}

// Time compares two time.Time equalities
func Time(b1 *time.Time, b2 *time.Time) bool {
	if b1 == nil {
		return b2 == nil
	}
	return b2 != nil && b1.Equal(*b2)
}
