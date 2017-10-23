package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wacul/ptr"
)

func TestConcatConfig(t *testing.T) {
	assert.Nil(t, concatConfig(nil, nil), "concat(NIL,NIL)")
	assert.Equal(t, &Config{CoverThreshold: ptr.Int(1)}, concatConfig(&Config{CoverThreshold: ptr.Int(1)}, nil), "concat(VAL, NIL)")
	assert.Equal(t, &Config{CoverThreshold: ptr.Int(1)}, concatConfig(&Config{CoverThreshold: ptr.Int(1)}, &Config{CoverThreshold: ptr.Int(2)}), "concat(VAL, VAL)")
	assert.Equal(t, &Config{CoverThreshold: ptr.Int(2)}, concatConfig(nil, &Config{CoverThreshold: ptr.Int(2)}), "concat(NIL, VAL)")
}

func TestActualConfig(t *testing.T) {
	t.Run("initial actual config", func(t *testing.T) {
		act := actualConfig(nil)
		if act == nil {
			t.Error("expected that the actual config never be nil, but not")
			t.FailNow()
		}
		if act.BuildStyle == nil {
			t.Error("expect that any property of the actual config never be nil, but not")
		}
	})

	t.Run("actual config will save a value", func(t *testing.T) {
		act := actualConfig(&Config{
			CoverThreshold: ptr.Int(10),
		})
		if act == nil {
			t.Error("expect that the actual config never be nil, but not")
			t.FailNow()
		}
		if act.CoverThreshold == nil {
			t.Error("expect that any property of the actual config never be nil, but not")
			t.FailNow()
		}
		if *act.CoverThreshold != 10 {
			t.Error("expect that the actual config will save a value, but not")
		}
	})

}

func TestConcatColor(t *testing.T) {
	color := func(n uint8) *Color {
		return &Color{Type: ColorType8Bit, Value8: n}
	}
	assert.Nil(t, concatColor(nil, nil), "concat(NIL,NIL)")
	assert.Equal(t, color(1), concatColor(color(1), nil), "concat(VAL, NIL)")
	assert.Equal(t, color(1), concatColor(color(1), color(2)), "concat(VAL, VAL)")
	assert.Equal(t, color(2), concatColor(nil, color(2)), "concat(NIL, VAL)")
}

func TestActualColor(t *testing.T) {
	color := func(n uint8) *Color {
		return &Color{Type: ColorType8Bit, Value8: n}
	}
	assert.Equal(t, &Color{Type: ColorTypeNone}, actualColor(nil), "actual(NIL)")
	assert.Equal(t, color(1), actualColor(color(1)), "actual(VAL)")
}

func TestConcatLabelType(t *testing.T) {
	foo := (*LabelType)(ptr.String("foo"))
	bar := (*LabelType)(ptr.String("bar"))
	assert.Nil(t, concatLabelType(nil, nil), "concat(NIL,NIL)")
	assert.Equal(t, foo, concatLabelType(foo, nil), "concat(VAL, NIL)")
	assert.Equal(t, foo, concatLabelType(foo, bar), "concat(VAL, VAL)")
	assert.Equal(t, bar, concatLabelType(nil, bar), "concat(NIL, VAL)")
}

func TestActualLabelType(t *testing.T) {
	long := LabelTypeLong
	foo := (*LabelType)(ptr.String("foo"))
	assert.Equal(t, &long, actualLabelType(nil), "actual(NIL)")
	assert.Equal(t, foo, actualLabelType(foo), "actual(VAL)")
}

func TestConcatStyle(t *testing.T) {
	style := func(n uint8) *Style {
		return &Style{Foreground: &Color{Type: ColorType8Bit, Value8: n}}
	}
	assert.Nil(t, concatStyle(nil, nil), "concat(NIL,NIL)")
	assert.Equal(t, style(1), concatStyle(style(1), nil), "concat(VAL, NIL)")
	assert.Equal(t, style(1), concatStyle(style(1), style(2)), "concat(VAL, VAL)")
	assert.Equal(t, style(2), concatStyle(nil, style(2)), "concat(NIL, VAL)")
}

func TestConcatInt(t *testing.T) {
	assert.Nil(t, concatInt(nil, nil), "concat(NIL,NIL)")
	assert.Equal(t, ptr.Int(1), concatInt(ptr.Int(1), nil), "concat(VAL, NIL)")
	assert.Equal(t, ptr.Int(1), concatInt(ptr.Int(1), ptr.Int(2)), "concat(VAL, VAL)")
	assert.Equal(t, ptr.Int(2), concatInt(nil, ptr.Int(2)), "concat(NIL, VAL)")
}

func TestActualInt(t *testing.T) {
	assert.Equal(t, ptr.Int(0), actualInt(nil), "actual(NIL)")
	assert.Equal(t, ptr.Int(1), actualInt(ptr.Int(1)), "actual(VAL)")
}

func TestConcatBool(t *testing.T) {
	assert.Nil(t, concatBool(nil, nil), "concat(NIL,NIL)")
	assert.Equal(t, ptr.Bool(true), concatBool(ptr.Bool(true), nil), "concat(VAL, NIL)")
	assert.Equal(t, ptr.Bool(true), concatBool(ptr.Bool(true), ptr.Bool(false)), "concat(VAL, VAL)")
	assert.Equal(t, ptr.Bool(false), concatBool(nil, ptr.Bool(false)), "concat(NIL, VAL)")
}

func TestActualBool(t *testing.T) {
	assert.Equal(t, ptr.Bool(false), actualBool(nil), "actual(NIL)")
	assert.Equal(t, ptr.Bool(true), actualBool(ptr.Bool(true)), "actual(VAL)")
}
