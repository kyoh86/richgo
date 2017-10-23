package editor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParrot(t *testing.T) {
	const src = "18wRDgGPuvsy6egaevFx"
	parrot := Parrot()
	out, err := parrot.Edit(src)
	if assert.NoError(t, err) {
		assert.Equal(t, src, out)
	}
}
