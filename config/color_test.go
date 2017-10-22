package config

import (
	"encoding/json"
	"fmt"
	"testing"

	yaml "gopkg.in/yaml.v2"
)

type testStruct struct {
	C Color `json:"c" yaml:"c"`
}

func TestMarshalYAML(t *testing.T) {
	for exp, color := range map[string]testStruct{
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
			t.Errorf("expected a marshaled color %q, but %q", exp, act)
		}
	}

	_, err := yaml.Marshal(testStruct{Color{Type: ColorType("foobar")}})
	if err == nil {
		t.Errorf("invalid ColorType %q accepted in marshaling", "foobar")
	}
}

func TestMarshalJSON(t *testing.T) {
	for exp, color := range map[string]testStruct{
		`{"c":"black"}`:   {C: Color{Type: ColorTypeName, Name: Black}},
		`{"c":"#ffeedd"}`: {C: Color{Type: ColorType24Bit, ValueR: 0xff, ValueG: 0xee, ValueB: 0xdd}},
		`{"c":31}`:        {C: Color{Type: ColorType8Bit, Value8: 31}},
	} {
		buf, err := json.Marshal(&color)
		if err != nil {
			t.Errorf("failed to marshal a color %q to json with error %q", "black", err)
			t.FailNow()
		}
		act := string(buf)
		if exp != act {
			t.Errorf("expected a marshaled color %q, but %q", exp, act)
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
		`0x1F`:                  Color{Type: ColorType8Bit, Value8: 0x1F},
		`"0x1F"`:                Color{Type: ColorType8Bit, Value8: 0x1F},
		`25`:                    Color{Type: ColorType8Bit, Value8: 25},
		`"red"`:                 Color{Type: ColorTypeName, Name: Red},
		`red`:                   Color{Type: ColorTypeName, Name: Red},
		`"#ffeedd"`:             Color{Type: ColorType24Bit, ValueR: 0xff, ValueG: 0xee, ValueB: 0xdd},
		`"rGb(255, 0xEe, 127)"`: Color{Type: ColorType24Bit, ValueR: 255, ValueG: 0xEE, ValueB: 127},
		`rGb(255, 0XEe, 127)`:   Color{Type: ColorType24Bit, ValueR: 255, ValueG: 0xEE, ValueB: 127},
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
			t.Errorf("expected a marshaled color %#v, but %#v", exp, obj.Act)
		}
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
		`"0x1F"`:                Color{Type: ColorType8Bit, Value8: 0x1F},
		`25`:                    Color{Type: ColorType8Bit, Value8: 25},
		`"red"`:                 Color{Type: ColorTypeName, Name: Red},
		`"#ffeedd"`:             Color{Type: ColorType24Bit, ValueR: 0xff, ValueG: 0xee, ValueB: 0xdd},
		`"rGb(255, 0xEe, 127)"`: Color{Type: ColorType24Bit, ValueR: 255, ValueG: 0xEE, ValueB: 127},
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
			t.Errorf("expected a marshaled color %#v, but %#v", exp, obj.Act)
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
