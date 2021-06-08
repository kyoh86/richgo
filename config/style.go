package config

import (
	"github.com/morikuni/aec"
)

// Style format the text with ANSI
type Style struct {
	// Hide text
	Hide *bool `json:"hide,omitempty" yaml:"hide,omitempty"`

	// Bold set the text style to bold or increased intensity.
	Bold *bool `json:"bold,omitempty" yaml:"bold,omitempty"`
	// Faint set the text style to faint.
	Faint *bool `json:"faint,omitempty" yaml:"faint,omitempty"`
	// Italic set the text style to italic.
	Italic *bool `json:"italic,omitempty" yaml:"italic,omitempty"`
	// Underline set the text style to underline.
	Underline *bool `json:"underline,omitempty" yaml:"underline,omitempty"`
	// BlinkSlow set the text style to slow blink.
	BlinkSlow *bool `json:"blinkSlow,omitempty" yaml:"blinkSlow,omitempty"`
	// BlinkRapid set the text style to rapid blink.
	BlinkRapid *bool `json:"blinkRapid,omitempty" yaml:"blinkRapid,omitempty"`
	// Inverse swap the foreground color and background color.
	Inverse *bool `json:"inverse,omitempty" yaml:"inverse,omitempty"`
	// Conceal set the text style to conceal.
	Conceal *bool `json:"conceal,omitempty" yaml:"conceal,omitempty"`
	// CrossOut set the text style to crossed out.
	CrossOut *bool `json:"crossOut,omitempty" yaml:"crossOut,omitempty"`
	// Frame set the text style to framed.
	Frame *bool `json:"frame,omitempty" yaml:"frame,omitempty"`
	// Encircle set the text style to encircled.
	Encircle *bool `json:"encircle,omitempty" yaml:"encircle,omitempty"`
	// Overline set the text style to overlined.
	Overline *bool `json:"overline,omitempty" yaml:"overline,omitempty"`

	// Foreground set the fore-color of text
	Foreground *Color `json:"foreground,omitempty" yaml:"foreground,omitempty"`
	// Foreground set the back-color of text
	Background *Color `json:"background,omitempty" yaml:"background,omitempty"`
}

// ANSI get the ANSI string
func (s *Style) ANSI() aec.ANSI {
	if s == nil {
		return emptyColor // when a prevLineStyle is not set, editor/test/test.go calls it in nil
	}

	ansi := s.Background.B()
	ansi = ansi.With(s.Foreground.F())
	for _, style := range []struct {
		flag *bool
		ansi aec.ANSI
	}{
		{s.Bold, aec.Bold},
		{s.Faint, aec.Faint},
		{s.Italic, aec.Italic},
		{s.Underline, aec.Underline},
		{s.BlinkSlow, aec.BlinkSlow},
		{s.BlinkRapid, aec.BlinkRapid},
		{s.Inverse, aec.Inverse},
		{s.Conceal, aec.Conceal},
		{s.CrossOut, aec.CrossOut},
		{s.Frame, aec.Frame},
		{s.Encircle, aec.Encircle},
		{s.Overline, aec.Overline},
	} {
		if *style.flag {
			ansi = ansi.With(style.ansi)
		}
	}
	return ansi
}

// Apply style To string
func (s *Style) Apply(str string) string {
	if s == nil {
		return str
	}

	if s.Hide != nil && *s.Hide {
		return ""
	}

	ansi := s.ANSI()
	if ansi == emptyColor {
		return str
	}

	if len(ansi.String()) == 0 {
		return str
	}

	return aec.Apply(str, ansi)
}
