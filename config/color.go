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
func (c Color) MarshalYAML() (interface{}, error) {
	switch c.Type {
	case ColorTypeNone:
		return "", nil
	case ColorTypeName:
		return c.Name.String(), nil
	case ColorType8Bit:
		return c.Value8, nil
	case ColorType24Bit:
		return fmt.Sprintf(`#%x%x%x`, c.ValueR, c.ValueG, c.ValueB), nil
	}
	return nil, fmt.Errorf("invalid color type %s", c.Type)
}

// MarshalJSON implements Marshaler
func (c Color) MarshalJSON() ([]byte, error) {
	switch c.Type {
	case ColorTypeNone:
		return []byte(`""`), nil
	case ColorTypeName:
		return []byte(fmt.Sprintf(`"%s"`, c.Name.String())), nil
	case ColorType8Bit:
		return []byte(fmt.Sprintf("%d", c.Value8)), nil
	case ColorType24Bit:
		return []byte(fmt.Sprintf(`"#%x%x%x"`, c.ValueR, c.ValueG, c.ValueB)), nil
	}
	return nil, fmt.Errorf("invalid color type %s", c.Type)
}

// UnmarshalYAML implements Unmarshaler
func (c *Color) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}
	return c.unmarshal(s, false)
}

var errInvalidFormat = errors.New("invalid format")

var reg8Bit = regexp.MustCompile(`(?mi)^(\d{1,3})$`)

func (c *Color) unmarshalAs8Bit(str string) error {
	match := reg8Bit.FindStringSubmatch(str)
	if len(match) != 2 {
		return errInvalidFormat
	}
	v, err := atoi(match[1])
	if err != nil {
		return errInvalidFormat
	}
	c.Type = ColorType8Bit
	c.Value8 = v
	return nil
}

var reg8BitHex = regexp.MustCompile(`(?mi)^(0[xX][[:xdigit:]]{1,2})$`)

func (c *Color) unmarshalAs8BitHex(str string) error {
	match := reg8BitHex.FindStringSubmatch(str)
	if len(match) != 2 {
		return errInvalidFormat
	}
	v, _ := atoi(match[1])
	c.Type = ColorType8Bit
	c.Value8 = v
	return nil
}

var regRGB = regexp.MustCompile(`(?mi)^#([[:xdigit:]]{2})([[:xdigit:]]{2})([[:xdigit:]]{2})$`)

func (c *Color) unmarshalAs24BitRGB(str string) error {
	match := regRGB.FindStringSubmatch(str)
	if len(match) != 4 {
		return errInvalidFormat
	}
	r, _ := strconv.ParseUint(match[1], 16, 8)
	g, _ := strconv.ParseUint(match[2], 16, 8)
	b, _ := strconv.ParseUint(match[3], 16, 8)
	c.Type = ColorType24Bit
	c.ValueR = uint8(r)
	c.ValueG = uint8(g)
	c.ValueB = uint8(b)
	return nil
}

var regRGBFunc = regexp.MustCompile(`(?mi)^rgb\((0x[[:xdigit:]]{2}|\d{1,3}), *(0x[[:xdigit:]]{2}|\d{1,3}), *(0x[[:xdigit:]]{2}|\d{1,3})\)$`)

func (c *Color) unmarshalAsRGBFunc(str string) error {
	match := regRGBFunc.FindStringSubmatch(str)
	if len(match) != 4 {
		return errInvalidFormat
	}
	r, err := atoi(match[1])
	if err != nil {
		return errInvalidFormat
	}
	g, err := atoi(match[2])
	if err != nil {
		return errInvalidFormat
	}
	b, err := atoi(match[3])
	if err != nil {
		return errInvalidFormat
	}

	c.Type = ColorType24Bit
	c.ValueR = r
	c.ValueG = g
	c.ValueB = b
	return nil
}

// UnmarshalJSON implements Unmarshaler
func (c *Color) UnmarshalJSON(raw []byte) error {
	return c.unmarshal(string(raw), true)
}

func (c *Color) unmarshal(str string, unquote bool) error {
	if unquote {
		unquoted, err := strconv.Unquote(str)
		if err == nil {
			str = unquoted
		}
	}
	if str == "" {
		c.Type = ColorTypeNone
		return nil
	}
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

func atoi(s string) (uint8, error) {
	i, err := atoiCore(s)
	if err != nil {
		return 0, err
	}
	return uint8(i), nil
}

func atoiCore(s string) (uint64, error) {
	if strings.HasPrefix(s, "0x") {
		return strconv.ParseUint(strings.TrimPrefix(s, "0x"), 16, 8)
	}
	if strings.HasPrefix(s, "0X") {
		return strconv.ParseUint(strings.TrimPrefix(s, "0X"), 16, 8)
	}
	return strconv.ParseUint(s, 10, 8)
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

var emptyColor aec.ANSI

func init() {
	emptyColor = aec.EmptyBuilder.ANSI
}

// B gets background ANSI color
func (c *Color) B() aec.ANSI {
	switch c.Type {
	case ColorTypeNone:
		return emptyColor
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
	return emptyColor
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
	case ColorTypeNone:
		return emptyColor
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
	return emptyColor
}
