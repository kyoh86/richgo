package config

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/morikuni/aec"
)

// Color is the color in the ANSI for configuration
type Color struct {
	Type ColorType

	Value8 uint8

	ValueR uint8
	ValueG uint8
	ValueB uint8

	Name ColorName
}

// MarshalYAML implements Marshaler
func (c Color) MarshalYAML() ([]byte, error) {
	return c.MarshalJSON()
}

// MarshalJSON implements Marshaler
func (c Color) MarshalJSON() ([]byte, error) {
	switch c.Type {
	case ColorTypeName:
		return []byte(c.Name.String()), nil
	case ColorType8Bit:
		return []byte(fmt.Sprintf("%d", c.Value8)), nil
	case ColorType24Bit:
		return []byte(fmt.Sprintf("#%x%x%x", c.ValueR, c.ValueG, c.ValueB)), nil
	}
	return nil, fmt.Errorf("invalid color type %s", c.Type)
}

var (
	regexpRGB = regexp.MustCompile(`(?mi)^rgb\s*\((0x[[:xdigit:]]{2}|%d{0,3}),\s*(0x[[:xdigit:]]{2}|%d{0,3}),\s*(0x[[:xdigit:]]{2}|%d{0,3})\)$`)
)

// UnmarshalYAML implements Unmarshaler
func (c *Color) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}
	return c.UnmarshalJSON([]byte(s))
}

// UnmarshalJSON implements Unmarshaler
func (c *Color) UnmarshalJSON(raw []byte) error {
	var (
		value8 uint8
		valueR uint8
		valueG uint8
		valueB uint8
	)

	str := string(raw)
	if n, err := fmt.Sscanf(str, "%d", &value8); err != nil && n == 1 {
		c.Type = ColorType8Bit
		c.Value8 = value8
	} else if n, err := fmt.Sscanf(str, "%#x", &value8); err != nil && n == 1 {
		c.Type = ColorType8Bit
		c.Value8 = value8
	} else if n, err := fmt.Sscanf(str, "#%x%x%x", &valueR, &valueG, &valueB); err != nil && n == 3 {
		c.Type = ColorType24Bit
		c.ValueR = valueR
		c.ValueG = valueG
		c.ValueB = valueB
	} else if match := regexpRGB.FindStringSubmatch(str); len(match) > 0 {
		c.Type = ColorType24Bit
		if len(match) != 3 {
			return fmt.Errorf("invalid RGB format : %s", str)
		}
		r, err := atoi(match[0])
		if err != nil {
			return fmt.Errorf("invalid RGB format in Red : %s", match[0])
		}
		c.ValueR = uint8(r)
		g, err := atoi(match[1])
		if err != nil {
			return fmt.Errorf("invalid RGB format in Green : %s", match[1])
		}
		c.ValueG = uint8(g)
		b, err := atoi(match[2])
		if err != nil {
			return fmt.Errorf("invalid RGB format in Blue : %s", match[2])
		}
		c.ValueB = uint8(b)
	} else {
		switch str {
		case "default":
			c.Name = DefaultColor

		case "black":
			c.Name = Black
		case "red":
			c.Name = Red
		case "green":
			c.Name = Green
		case "yellow":
			c.Name = Yellow
		case "blue":
			c.Name = Blue
		case "magenta":
			c.Name = Magenta
		case "cyan":
			c.Name = Cyan
		case "white":
			c.Name = White

		case "lightBlack":
			c.Name = LightBlack
		case "lightRed":
			c.Name = LightRed
		case "lightGreen":
			c.Name = LightGreen
		case "lightYellow":
			c.Name = LightYellow
		case "lightBlue":
			c.Name = LightBlue
		case "lightMagenta":
			c.Name = LightMagenta
		case "lightCyan":
			c.Name = LightCyan
		case "lightWhite":
			c.Name = LightWhite
		default:
			return fmt.Errorf("invalid color name : %s", str)
		}
	}
	return nil
}

func atoi(s string) (uint64, error) {
	if strings.HasPrefix(s, "0x") {
		return strconv.ParseUint(strings.TrimPrefix(s, "0x"), 16, 8)
	}
	return strconv.ParseUint(s, 10, 8)
}

func concatColor(a, b *Color) *Color {
	if a == nil {
		return b
	}
	return a
}

func actualColor(c *Color) *Color {
	if c == nil {
		return &Color{
			Type: ColorTypeName,
			Name: DefaultColor,
		}
	}
	return c
}

// B gets background ANSI color
func (c *Color) B() aec.ANSI {
	switch c.Type {
	case ColorTypeName:
		switch c.Name {
		case Black:
			return aec.BlackB
		case Red:
			return aec.RedB
		case Green:
			return aec.GreenB
		case Yellow:
			return aec.YellowB
		case Blue:
			return aec.BlueB
		case Magenta:
			return aec.MagentaB
		case Cyan:
			return aec.CyanB
		case White:
			return aec.WhiteB

		case LightBlack:
			return aec.LightBlackB
		case LightRed:
			return aec.LightRedB
		case LightGreen:
			return aec.LightGreenB
		case LightYellow:
			return aec.LightYellowB
		case LightBlue:
			return aec.LightBlueB
		case LightMagenta:
			return aec.LightMagentaB
		case LightCyan:
			return aec.LightCyanB
		case LightWhite:
			return aec.LightWhiteB
		}
	case ColorType8Bit:
		return aec.Color8BitB(aec.RGB8Bit(c.Value8))
	case ColorType24Bit:
		return aec.FullColorB(c.ValueR, c.ValueG, c.ValueB)
	}
	return aec.DefaultB
}

// F gets foreground ANSI color
func (c *Color) F() aec.ANSI {
	switch c.Type {
	case ColorTypeName:
		switch c.Name {
		case Black:
			return aec.BlackF
		case Red:
			return aec.RedF
		case Green:
			return aec.GreenF
		case Yellow:
			return aec.YellowF
		case Blue:
			return aec.BlueF
		case Magenta:
			return aec.MagentaF
		case Cyan:
			return aec.CyanF
		case White:
			return aec.WhiteF

		case LightBlack:
			return aec.LightBlackF
		case LightRed:
			return aec.LightRedF
		case LightGreen:
			return aec.LightGreenF
		case LightYellow:
			return aec.LightYellowF
		case LightBlue:
			return aec.LightBlueF
		case LightMagenta:
			return aec.LightMagentaF
		case LightCyan:
			return aec.LightCyanF
		case LightWhite:
			return aec.LightWhiteF
		}
	case ColorType8Bit:
		return aec.Color8BitF(aec.RGB8Bit(c.Value8))
	case ColorType24Bit:
		return aec.FullColorF(c.ValueR, c.ValueG, c.ValueB)
	}
	return aec.DefaultF
}
