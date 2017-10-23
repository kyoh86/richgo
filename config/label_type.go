package config

import "fmt"

// LabelType is the type of line-labels
type LabelType string

const (
	// LabelTypeNone suppress line-labels
	LabelTypeNone = LabelType("none")
	// LabelTypeShort prints single-character line-label
	LabelTypeShort = LabelType("short")
	// LabelTypeLong prints text line-label
	LabelTypeLong = LabelType("long")
)

func (l LabelType) String() string { return string(l) }

// MarshalJSON implements Marshaler
func (l LabelType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", l)), nil
}

// UnmarshalJSON implements Unmarshaler
func (l *LabelType) UnmarshalJSON(raw []byte) error {
	switch str := string(raw); str {
	case `"none"`:
		*l = LabelTypeNone
	case `"short"`:
		*l = LabelTypeShort
	case `"long"`:
		*l = LabelTypeLong
	default:
		return fmt.Errorf("invalid LabelType %s", str)
	}
	return nil
}

// LabelTypes defines the possible values of LabelType
func LabelTypes() []LabelType {
	return []LabelType{
		LabelTypeNone,
		LabelTypeShort,
		LabelTypeLong,
	}
}
