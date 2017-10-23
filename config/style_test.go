package config

import (
	"testing"

	"github.com/morikuni/aec"
	"github.com/wacul/ptr"
)

func TestApply(t *testing.T) {
	const base = "abc"
	for exp, style := range map[string]*Style{
		aec.RedF.Apply(base): actualStyle(&Style{
			Foreground: &Color{Type: ColorTypeName, Name: Red},
		}),
		aec.Bold.Apply(base): actualStyle(&Style{
			Bold: ptr.Bool(true),
		}),
		"": actualStyle(&Style{
			Hide: ptr.Bool(true),
		}),
		base: actualStyle(nil),
	} {
		act := style.Apply(base)
		if exp != act {
			t.Errorf("expect that a style applied as %q, but %q", exp, act)
		}
	}
}
