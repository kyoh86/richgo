package config

import (
	"github.com/morikuni/aec"
	"github.com/wacul/ptr"
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

func concatStyle(base, other *Style) *Style {
	if base == nil {
		if other == nil {
			return nil
		}
		base = &Style{}
	}
	if other == nil {
		other = &Style{}
	}
	return &Style{
		Hide:       concatBool(base.Hide, other.Hide),
		Bold:       concatBool(base.Bold, other.Bold),
		Faint:      concatBool(base.Faint, other.Faint),
		Italic:     concatBool(base.Italic, other.Italic),
		Underline:  concatBool(base.Underline, other.Underline),
		BlinkSlow:  concatBool(base.BlinkSlow, other.BlinkSlow),
		BlinkRapid: concatBool(base.BlinkRapid, other.BlinkRapid),
		Inverse:    concatBool(base.Inverse, other.Inverse),
		Conceal:    concatBool(base.Conceal, other.Conceal),
		CrossOut:   concatBool(base.CrossOut, other.CrossOut),
		Frame:      concatBool(base.Frame, other.Frame),
		Encircle:   concatBool(base.Encircle, other.Encircle),
		Overline:   concatBool(base.Overline, other.Overline),
		Foreground: concatColor(base.Foreground, other.Foreground),
		Background: concatColor(base.Background, other.Background),
	}
}

func actualStyle(s *Style) *Style {
	if s == nil {
		s = &Style{}
	}
	return &Style{
		Hide:       actualBool(s.Hide),
		Bold:       actualBool(s.Bold),
		Faint:      actualBool(s.Faint),
		Italic:     actualBool(s.Italic),
		Underline:  actualBool(s.Underline),
		BlinkSlow:  actualBool(s.BlinkSlow),
		BlinkRapid: actualBool(s.BlinkRapid),
		Inverse:    actualBool(s.Inverse),
		Conceal:    actualBool(s.Conceal),
		CrossOut:   actualBool(s.CrossOut),
		Frame:      actualBool(s.Frame),
		Encircle:   actualBool(s.Encircle),
		Overline:   actualBool(s.Overline),
		Foreground: actualColor(s.Foreground),
		Background: actualColor(s.Background),
	}
}

func concatBool(a, b *bool) *bool {
	if a == nil {
		if b == nil {
			return nil
		}
		f := *b
		return &f
	}
	return a
}

func actualBool(b *bool) *bool {
	if b == nil {
		return ptr.Bool(false)
	}
	return b
}

// Apply style To string
func (s *Style) Apply(str string) string {
	if s == nil {
		return str
	}
	if is(s.Hide) {
		return ""
	}

	ansi := []aec.ANSI{
		s.Background.B(),
		s.Foreground.F(),
	}

	if is(s.Bold) {
		ansi = append(ansi, aec.Bold)
	}
	if is(s.Faint) {
		ansi = append(ansi, aec.Faint)
	}
	if is(s.Italic) {
		ansi = append(ansi, aec.Italic)
	}
	if is(s.Underline) {
		ansi = append(ansi, aec.Underline)
	}
	if is(s.BlinkSlow) {
		ansi = append(ansi, aec.BlinkSlow)
	}
	if is(s.BlinkRapid) {
		ansi = append(ansi, aec.BlinkRapid)
	}
	if is(s.Inverse) {
		ansi = append(ansi, aec.Inverse)
	}
	if is(s.Conceal) {
		ansi = append(ansi, aec.Conceal)
	}
	if is(s.CrossOut) {
		ansi = append(ansi, aec.CrossOut)
	}
	if is(s.Frame) {
		ansi = append(ansi, aec.Frame)
	}
	if is(s.Encircle) {
		ansi = append(ansi, aec.Encircle)
	}
	if is(s.Overline) {
		ansi = append(ansi, aec.Overline)
	}
	return aec.Apply(str, ansi...)
}

func is(f *bool) bool {
	if f == nil {
		return false
	}
	return *f
}
