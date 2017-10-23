package config

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLabelTypeMarshaling(t *testing.T) {
	type container struct {
		L LabelType
	}
	for _, l := range LabelTypes() {
		var d container
		raw, err := json.Marshal(container{l})
		if assert.NoError(t, err, "marshaling", l) {
			if err := json.Unmarshal(raw, &d); assert.NoError(t, err, "unmarshaling", l) {
				assert.Equal(t, l, d.L, "idenpocy of", l)
			}
		}
	}
	assert.Error(t, json.Unmarshal([]byte(`{"L":null}`), &container{}), "nil body")
	assert.Error(t, json.Unmarshal([]byte(`{"L":""}`), &container{}), "empty body")
}
