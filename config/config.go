package config

import "github.com/wacul/ptr"

// Config holds settings for richgo
type Config struct {
	LabelType *LabelType `json:"labelType,omitempty" yaml:"labelType,omitempty"`

	BuildStyle *Style `json:"buildStyle,omitempty" yaml:"buildStyle,omitempty"`
	StartStyle *Style `json:"startStyle,omitempty" yaml:"startStyle,omitempty"`
	PassStyle  *Style `json:"passStyle,omitempty" yaml:"passStyle,omitempty"`
	FailStyle  *Style `json:"failStyle,omitempty" yaml:"failStyle,omitempty"`
	SkipStyle  *Style `json:"skipStyle,omitempty" yaml:"skipStyle,omitempty"`
	FileStyle  *Style `json:"fileStyle,omitempty" yaml:"fileStyle,omitempty"`
	LineStyle  *Style `json:"lineStyle,omitempty" yaml:"lineStyle,omitempty"`

	CoverThreshold *int   `json:"coverThreshold,omitempty" yaml:"coverThreshold,omitempty"`
	CoveredStyle   *Style `json:"coveredStyle,omitempty" yaml:"coveredStyle,omitempty"`
	UncoveredStyle *Style `json:"uncoveredStyle,omitempty" yaml:"uncoveredStyle,omitempty"`

	Removals []string `json:"removals,omitempty" yaml:"removals,omitempty"`
}

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
		LabelType:      concatLabelType(base.LabelType, other.LabelType),
		BuildStyle:     concatStyle(base.BuildStyle, other.BuildStyle),
		StartStyle:     concatStyle(base.StartStyle, other.StartStyle),
		PassStyle:      concatStyle(base.PassStyle, other.PassStyle),
		FailStyle:      concatStyle(base.FailStyle, other.FailStyle),
		SkipStyle:      concatStyle(base.SkipStyle, other.SkipStyle),
		FileStyle:      concatStyle(base.FileStyle, other.FileStyle),
		LineStyle:      concatStyle(base.LineStyle, other.LineStyle),
		CoverThreshold: concatInt(base.CoverThreshold, other.CoverThreshold),
		CoveredStyle:   concatStyle(base.CoveredStyle, other.CoveredStyle),
		UncoveredStyle: concatStyle(base.UncoveredStyle, other.UncoveredStyle),
		Removals:       append(base.Removals, other.Removals...),
	}
}

func actualConfig(config *Config) *Config {
	if config == nil {
		config = &Config{}
	}
	return &Config{
		LabelType:      actualLabelType(config.LabelType),
		BuildStyle:     actualStyle(config.BuildStyle),
		StartStyle:     actualStyle(config.StartStyle),
		PassStyle:      actualStyle(config.PassStyle),
		FailStyle:      actualStyle(config.FailStyle),
		SkipStyle:      actualStyle(config.SkipStyle),
		FileStyle:      actualStyle(config.FileStyle),
		LineStyle:      actualStyle(config.LineStyle),
		CoverThreshold: actualInt(config.CoverThreshold),
		CoveredStyle:   actualStyle(config.CoveredStyle),
		UncoveredStyle: actualStyle(config.UncoveredStyle),
		Removals:       config.Removals,
	}
}
