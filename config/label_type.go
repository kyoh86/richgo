package config

import "fmt"

// LabelType is the type of line-labels
type LabelType string

const (
	// LabelTypeNone suppress line-labels
	LabelTypeNone = LabelType("none")
	// LabelTypeShort prints single-charactor line-label
	LabelTypeShort = LabelType("short")
	// LabelTypeLong prints text line-label
	LabelTypeLong = LabelType("long")
)

func (l LabelType) String() string { return string(l) }

// MarshalJSON implements Marshaler
func (l LabelType) MarshalJSON() ([]byte, error) {
	return []byte(l.String()), nil
}

// UnmarshalJSON implements Unmarshaler
func (l *LabelType) UnmarshalJSON(raw []byte) error {
	switch str := string(raw); str {
	case "none":
		*l = LabelTypeNone
	case "short":
		*l = LabelTypeShort
	case "long":
		*l = LabelTypeLong
	default:
		return fmt.Errorf("invalid LabelType %s", str)
	}
	return nil
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
