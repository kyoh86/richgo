package config

import "github.com/wacul/ptr"

func concatInt(a, b *int) *int {
	if a == nil {
		if b == nil {
			return nil
		}
		f := *b
		return &f
	}
	return a
}

func actualInt(b *int) *int {
	if b == nil {
		return ptr.Int(0)
	}
	return b
}

func concatConfig(base, other *Config) *Config {
	if base == nil {
		if other == nil {
			return nil
		}
		base = &Config{}
	}
	if other == nil {
		other = &Config{}
	}
	return &Config{
		LabelType:        concatLabelType(base.LabelType, other.LabelType),
		BuildStyle:       concatStyle(base.BuildStyle, other.BuildStyle),
		StartStyle:       concatStyle(base.StartStyle, other.StartStyle),
		PassStyle:        concatStyle(base.PassStyle, other.PassStyle),
		FailStyle:        concatStyle(base.FailStyle, other.FailStyle),
		PassPackageStyle: concatStyle(base.PassPackageStyle, other.PassPackageStyle),
		FailPackageStyle: concatStyle(base.FailPackageStyle, other.FailPackageStyle),
		SkipStyle:        concatStyle(base.SkipStyle, other.SkipStyle),
		FileStyle:        concatStyle(base.FileStyle, other.FileStyle),
		LineStyle:        concatStyle(base.LineStyle, other.LineStyle),
		CoverThreshold:   concatInt(base.CoverThreshold, other.CoverThreshold),
		CoveredStyle:     concatStyle(base.CoveredStyle, other.CoveredStyle),
		UncoveredStyle:   concatStyle(base.UncoveredStyle, other.UncoveredStyle),
		Removals:         append(base.Removals, other.Removals...),
		LeaveTestPrefix:  concatBool(base.LeaveTestPrefix, other.LeaveTestPrefix),
	}
}

func actualConfig(config *Config) *Config {
	if config == nil {
		config = &Config{}
	}
	return &Config{
		LabelType:        actualLabelType(config.LabelType),
		BuildStyle:       actualStyle(config.BuildStyle),
		StartStyle:       actualStyle(config.StartStyle),
		PassStyle:        actualStyle(config.PassStyle),
		FailStyle:        actualStyle(config.FailStyle),
		PassPackageStyle: actualStyle(config.PassPackageStyle),
		FailPackageStyle: actualStyle(config.FailPackageStyle),
		SkipStyle:        actualStyle(config.SkipStyle),
		FileStyle:        actualStyle(config.FileStyle),
		LineStyle:        actualStyle(config.LineStyle),
		CoverThreshold:   actualInt(config.CoverThreshold),
		CoveredStyle:     actualStyle(config.CoveredStyle),
		UncoveredStyle:   actualStyle(config.UncoveredStyle),
		Removals:         config.Removals,
		LeaveTestPrefix:  actualBool(config.LeaveTestPrefix),
	}
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
			Type: ColorTypeNone,
		}
	}
	return c
}

func concatLabelType(base, other *LabelType) *LabelType {
	if base == nil {
		if other == nil {
			return nil
		}
		return other
	}
	return base
}

func actualLabelType(t *LabelType) *LabelType {
	if t == nil {
		l := LabelTypeLong
		return &l
	}
	return t
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
