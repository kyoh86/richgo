package config

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	yaml "gopkg.in/yaml.v2"
)

type testStruct struct {
	C Color `json:"c" yaml:"c,omitempty"`
}

func TestMarshalYAML(t *testing.T) {
	for exp, color := range map[string]testStruct{
		`{}`:           {},
		`c: ""`:        {Color{Type: ColorTypeNone}},
		`c: black`:     {Color{Type: ColorTypeName, Name: Black}},
		`c: '#ffeedd'`: {Color{Type: ColorType24Bit, ValueR: 0xff, ValueG: 0xee, ValueB: 0xdd}},
		`c: 31`:        {Color{Type: ColorType8Bit, Value8: 31}},
	} {
		buf, err := yaml.Marshal(&color)
		if err != nil {
			t.Errorf("failed to marshal a color %q to yaml with error %q", "black", err)
			t.FailNow()
		}
		exp += "\n" // NOTE: yaml will have trailing newline
		act := string(buf)
		if exp != act {
			t.Errorf("expect that a marshaled color be %q, but %q", exp, act)
		}
	}

	_, err := yaml.Marshal(testStruct{Color{Type: ColorType("foobar")}})
	if err == nil {
		t.Errorf("invalid ColorType %q accepted in marshaling", "foobar")
	}
}

func TestMarshalJSON(t *testing.T) {
	for exp, color := range map[string]testStruct{
		`{"c":""}`:        {C: Color{Type: ColorTypeNone}},
		`{"c":"black"}`:   {C: Color{Type: ColorTypeName, Name: Black}},
		`{"c":"#ffeedd"}`: {C: Color{Type: ColorType24Bit, ValueR: 0xff, ValueG: 0xee, ValueB: 0xdd}},
		`{"c":31}`:        {C: Color{Type: ColorType8Bit, Value8: 31}},
	} {
		buf, err := json.Marshal(&color)
		if err != nil {
			t.Errorf("failed to marshal a color to json with error %q", err)
			t.FailNow()
		}
		act := string(buf)
		if exp != act {
			t.Errorf("expect that a marshaled color be %q, but %q", exp, act)
		}
	}

	_, err := json.Marshal(testStruct{Color{Type: ColorType("foobar")}})
	if err == nil {
		t.Errorf("invalid ColorType %q accepted in marshaling", "foobar")
	}
}

func TestUnmarshalYAML(t *testing.T) {
	const yamlTemplate = "act: %s"
	for value, exp := range map[string]Color{
		`""`:                    {Type: ColorTypeNone},
		`0x1F`:                  {Type: ColorType8Bit, Value8: 0x1F},
		`"0x1F"`:                {Type: ColorType8Bit, Value8: 0x1F},
		`25`:                    {Type: ColorType8Bit, Value8: 25},
		`"red"`:                 {Type: ColorTypeName, Name: Red},
		`red`:                   {Type: ColorTypeName, Name: Red},
		`"#ffeedd"`:             {Type: ColorType24Bit, ValueR: 0xff, ValueG: 0xee, ValueB: 0xdd},
		`"rGb(255, 0xEe, 127)"`: {Type: ColorType24Bit, ValueR: 255, ValueG: 0xEE, ValueB: 127},
		`rGb(255, 0XEe, 127)`:   {Type: ColorType24Bit, ValueR: 255, ValueG: 0xEE, ValueB: 127},
	} {
		var obj struct {
			Act Color `yaml:"act"`
		}
		err := yaml.Unmarshal([]byte(fmt.Sprintf(yamlTemplate, value)), &obj)
		if err != nil {
			t.Errorf("failed to unmarshal a color %q as yaml with error %q", value, err)
			continue
		}
		if obj.Act != exp {
			t.Errorf("expect that a marshaled color be %#v, but %#v", exp, obj.Act)
		}
	}

	var obj struct {
		Act *Color `yaml:"act"`
	}
	err := yaml.Unmarshal([]byte{}, &obj)
	if err != nil {
		t.Errorf("failed to unmarshal a empty color as yaml with error %q", err)
	}
	if obj.Act != nil {
		t.Errorf("expect that a marshaled color be nil, but %#v", obj.Act)
	}

	for _, value := range []string{
		`{}`,
		`[]`,
		`256`,
		`rgb(256,0,0)`,
		`rgb(0,256,0)`,
		`rgb(0,0,256)`,
	} {
		var obj struct {
			Act Color `yaml:"act"`
		}
		err := yaml.Unmarshal([]byte(fmt.Sprintf(yamlTemplate, value)), &obj)
		if err == nil {
			t.Errorf("invalid color %q accepted in unmarshaling as %#v", value, obj)
		}
	}
}

func TestUnmarshalJSON(t *testing.T) {
	const jsonTemplate = `{"act":%s}`
	for value, exp := range map[string]Color{
		`""`:                    {Type: ColorTypeNone},
		`"0x1F"`:                {Type: ColorType8Bit, Value8: 0x1F},
		`25`:                    {Type: ColorType8Bit, Value8: 25},
		`"red"`:                 {Type: ColorTypeName, Name: Red},
		`"#ffeedd"`:             {Type: ColorType24Bit, ValueR: 0xff, ValueG: 0xee, ValueB: 0xdd},
		`"rGb(255, 0xEe, 127)"`: {Type: ColorType24Bit, ValueR: 255, ValueG: 0xEE, ValueB: 127},
	} {
		var obj struct {
			Act Color `json:"act"`
		}
		err := json.Unmarshal([]byte(fmt.Sprintf(jsonTemplate, value)), &obj)
		if err != nil {
			t.Errorf("failed to unmarshal a color %q as json with error %q", value, err)
			continue
		}
		if obj.Act != exp {
			t.Errorf("expect that a marshaled color be %#v, but %#v", exp, obj.Act)
		}
	}
	for _, value := range []string{
		`{}`,
		`[]`,
		`256`,
		`"rgb(256,0,0)"`,
		`"rgb(0,256,0)"`,
		`"rgb(0,0,256)"`,
	} {
		var obj struct {
			Act Color `json:"act"`
		}
		err := json.Unmarshal([]byte(fmt.Sprintf(jsonTemplate, value)), &obj)
		if err == nil {
			t.Errorf("invalid color %q accepted in unmarshaling as %#v", value, obj)
		}
	}
}

func TestB(t *testing.T) {
	const esc = "\x1b["
	for _, c := range []Color{
		{Type: ColorTypeName, Name: Black},
		{Type: ColorType8Bit, Value8: 11},
		{Type: ColorType24Bit, ValueR: 33, ValueG: 22, ValueB: 11},
	} {
		// B() func must return Back Color (formed ESC+[+3x)
		if !strings.HasPrefix(c.B().String(), esc+"4") {
			t.Errorf("invalid front color: %s", strings.TrimPrefix(c.B().String(), esc))
		}
	}

	for _, c := range []Color{
		{},
		{Type: ColorTypeName, Name: "invalidColor"},
	} {
		c := c
		assert.Equal(t, "", c.B().String())
	}
}

func TestF(t *testing.T) {
	const esc = "\x1b["
	for _, c := range []Color{
		{Type: ColorTypeName, Name: Black},
		{Type: ColorType8Bit, Value8: 11},
		{Type: ColorType24Bit, ValueR: 33, ValueG: 22, ValueB: 11},
	} {
		// F() func must return Front Color (formed ESC+[+3x)
		if !strings.HasPrefix(c.F().String(), esc+"3") {
			t.Errorf("invalid front color: %s", strings.TrimPrefix(c.F().String(), esc))
		}
	}

	for _, c := range []Color{
		{},
		{Type: ColorTypeName, Name: "invalidColor"},
	} {
		c := c
		assert.Equal(t, "", c.F().String())
	}
}
