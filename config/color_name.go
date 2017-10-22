package config

// ColorName is the name of ANSI-colors
type ColorName string

func (n ColorName) String() string { return string(n) }

// Colors
const (
	DefaultColor = ColorName("default")

	Black   = ColorName("black")
	Red     = ColorName("red")
	Green   = ColorName("green")
	Yellow  = ColorName("yellow")
	Blue    = ColorName("blue")
	Magenta = ColorName("magenta")
	Cyan    = ColorName("cyan")
	White   = ColorName("white")

	LightBlack   = ColorName("lightBlack")
	LightRed     = ColorName("lightRed")
	LightGreen   = ColorName("lightGreen")
	LightYellow  = ColorName("lightYellow")
	LightBlue    = ColorName("lightBlue")
	LightMagenta = ColorName("lightMagenta")
	LightCyan    = ColorName("lightCyan")
	LightWhite   = ColorName("lightWhite")
)

// ColorNames will get variations of the colors.
func ColorNames() []ColorName {
	return []ColorName{
		DefaultColor,
		Black,
		Red,
		Green,
		Yellow,
		Blue,
		Magenta,
		Cyan,
		White,
		LightBlack,
		LightRed,
		LightGreen,
		LightYellow,
		LightBlue,
		LightMagenta,
		LightCyan,
		LightWhite,
	}
}
