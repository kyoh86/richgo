// Package ptr provides utility functions to get the pointer of the built-in type
package ptr

import "time"

// Uint8 will get a pointer of the uint8
func Uint8(u uint8) *uint8 {
	return &u
}

// Uint16 will get a pointer of the uint16
func Uint16(u uint16) *uint16 {
	return &u
}

// Uint32 will get a pointer of the uint32
func Uint32(u uint32) *uint32 {
	return &u
}

// Uint64 will get a pointer of the uint64
func Uint64(u uint64) *uint64 {
	return &u
}

// Int8 will get a pointer of the int8
func Int8(i int8) *int8 {
	return &i
}

// Int16 will get a pointer of the int16
func Int16(i int16) *int16 {
	return &i
}

// Int32 will get a pointer of the int32
func Int32(i int32) *int32 {
	return &i
}

// Int64 will get a pointer of the int64
func Int64(i int64) *int64 {
	return &i
}

// Float32 will get a pointer of the float32
func Float32(f float32) *float32 {
	return &f
}

// Float64 will get a pointer of the float64
func Float64(f float64) *float64 {
	return &f
}

// Complex64 will get a pointer of the complex64
func Complex64(c complex64) *complex64 {
	return &c
}

// Complex128 will get a pointer of the complex128
func Complex128(c complex128) *complex128 {
	return &c
}

// Byte will get a pointer of the byte
func Byte(b byte) *byte {
	return &b
}

// Rune will get a pointer of the rune
func Rune(r rune) *rune {
	return &r
}

// Uint will get a pointer of the uint
func Uint(u uint) *uint {
	return &u
}

// Int will get a pointer of the int
func Int(i int) *int {
	return &i
}

// String will get a pointer of the string
func String(s string) *string {
	return &s
}

// Bool will get a pointer of the bool
func Bool(b bool) *bool {
	return &b
}

// True will get a pointer of the true
func True() *bool {
	return Bool(true)
}

// False will get a pointer of the false
func False() *bool {
	return Bool(false)
}

// Time will get a pointer of the time.Time
func Time(t time.Time) *time.Time {
	return &t
}
