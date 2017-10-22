package config

import (
	"errors"
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

// UnmarshalYAML implements Unmarshaler
func (c *Color) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}
	return c.UnmarshalJSON([]byte(s))
}

var errInvalidFormat = errors.New("invalid format")

func (c *Color) unmarshalAs8Bit(str string) error {
	var value8 uint8
	if n, err := fmt.Sscanf(str, "%d", &value8); err != nil || n != 1 {
		return errInvalidFormat
	}
	c.Type = ColorType8Bit
	c.Value8 = value8
	return nil
}

func (c *Color) unmarshalAs8BitHex(str string) error {
	var value8 uint8
	if n, err := fmt.Sscanf(str, "%#x", &value8); err != nil || n != 1 {
		return errInvalidFormat
	}
	c.Type = ColorType8Bit
	c.Value8 = value8
	return nil
}

func (c *Color) unmarshalAs24BitRGB(str string) error {
	var (
		valueR uint8
		valueG uint8
		valueB uint8
	)
	if n, err := fmt.Sscanf(str, "#%02x%02x%02x", &valueR, &valueG, &valueB); err != nil || n != 3 {
		return errInvalidFormat
	}
	c.Type = ColorType24Bit
	c.ValueR = valueR
	c.ValueG = valueG
	c.ValueB = valueB
	return nil
}

var (
	regexpRGB = regexp.MustCompile(`(?mi)^rgb\s*\((0x[[:xdigit:]]{2}|\d{0,3}),\s*(0x[[:xdigit:]]{2}|\d{0,3}),\s*(0x[[:xdigit:]]{2}|\d{0,3})\)$`)
)

func (c *Color) unmarshalAsRGBFunc(str string) error {
	match := regexpRGB.FindStringSubmatch(str)
	if len(match) != 3 {
		return errInvalidFormat
	}
	r, err := atoi(match[0])
	if err != nil {
		return errInvalidFormat
	}
	g, err := atoi(match[1])
	if err != nil {
		return errInvalidFormat
	}
	b, err := atoi(match[2])
	if err != nil {
		return errInvalidFormat
	}

	c.Type = ColorType24Bit
	c.ValueR = uint8(r)
	c.ValueG = uint8(g)
	c.ValueB = uint8(b)
	return nil
}

// UnmarshalJSON implements Unmarshaler
func (c *Color) UnmarshalJSON(raw []byte) error {
	str := string(raw)
	if err := c.unmarshalAs8Bit(str); err == nil {
		return nil
	}
	if err := c.unmarshalAs8BitHex(str); err == nil {
		return nil
	}
	if err := c.unmarshalAs24BitRGB(str); err == nil {
		return nil
	}
	if err := c.unmarshalAsRGBFunc(str); err == nil {
		return nil
	}
	for _, cn := range ColorNames() {
		if cn.String() == str {
			c.Type = ColorTypeName
			c.Name = cn
			return nil
		}
	}
	return errInvalidFormat
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

var backColors = map[ColorName]aec.ANSI{
	Black:        aec.BlackB,
	Red:          aec.RedB,
	Green:        aec.GreenB,
	Yellow:       aec.YellowB,
	Blue:         aec.BlueB,
	Magenta:      aec.MagentaB,
	Cyan:         aec.CyanB,
	White:        aec.WhiteB,
	LightBlack:   aec.LightBlackB,
	LightRed:     aec.LightRedB,
	LightGreen:   aec.LightGreenB,
	LightYellow:  aec.LightYellowB,
	LightBlue:    aec.LightBlueB,
	LightMagenta: aec.LightMagentaB,
	LightCyan:    aec.LightCyanB,
	LightWhite:   aec.LightWhiteB,
}

// B gets background ANSI color
func (c *Color) B() aec.ANSI {
	switch c.Type {
	case ColorTypeName:
		b, ok := backColors[c.Name]
		if ok {
			return b
		}
	case ColorType8Bit:
		return aec.Color8BitB(aec.RGB8Bit(c.Value8))
	case ColorType24Bit:
		return aec.FullColorB(c.ValueR, c.ValueG, c.ValueB)
	}
	return aec.DefaultB
}

var frontColors = map[ColorName]aec.ANSI{
	Black:        aec.BlackF,
	Red:          aec.RedF,
	Green:        aec.GreenF,
	Yellow:       aec.YellowF,
	Blue:         aec.BlueF,
	Magenta:      aec.MagentaF,
	Cyan:         aec.CyanF,
	White:        aec.WhiteF,
	LightBlack:   aec.LightBlackF,
	LightRed:     aec.LightRedF,
	LightGreen:   aec.LightGreenF,
	LightYellow:  aec.LightYellowF,
	LightBlue:    aec.LightBlueF,
	LightMagenta: aec.LightMagentaF,
	LightCyan:    aec.LightCyanF,
	LightWhite:   aec.LightWhiteF,
}

// F gets foreground ANSI color
func (c *Color) F() aec.ANSI {
	switch c.Type {
	case ColorTypeName:
		f, ok := frontColors[c.Name]
		if ok {
			return f
		}
	case ColorType8Bit:
		return aec.Color8BitF(aec.RGB8Bit(c.Value8))
	case ColorType24Bit:
		return aec.FullColorF(c.ValueR, c.ValueG, c.ValueB)
	}
	return aec.DefaultF
}
