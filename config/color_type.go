package config

// ColorType is the type of color
type ColorType string

const (
	// ColorTypeNone defines empty
	ColorTypeNone = ColorType("none")
	// ColorType8Bit defines 8-bit (256) colors
	ColorType8Bit = ColorType("8bit")
	// ColorType24Bit defines 24-bit (R: 8bit + G: 8bit + B: 8bit ; full) colors
	ColorType24Bit = ColorType("24bit")
	// ColorTypeName defines named colors
	ColorTypeName = ColorType("named")
)

func (l ColorType) String() string { return string(l) }
