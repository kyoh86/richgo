package config

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

	PassPackageStyle *Style `json:"passPackageStyle,omitempty" yaml:"passPackageStyle,omitempty"`
	FailPackageStyle *Style `json:"failPackageStyle,omitempty" yaml:"failPackageStyle,omitempty"`

	CoverThreshold *int   `json:"coverThreshold,omitempty" yaml:"coverThreshold,omitempty"`
	CoveredStyle   *Style `json:"coveredStyle,omitempty" yaml:"coveredStyle,omitempty"`
	UncoveredStyle *Style `json:"uncoveredStyle,omitempty" yaml:"uncoveredStyle,omitempty"`

	Removals []string `json:"removals,omitempty" yaml:"removals,omitempty"`

	LeaveTestPrefix *bool `json:"leaveTestPrefix,omitempty" yaml:"leaveTestPrefix,omitempty"`
}
