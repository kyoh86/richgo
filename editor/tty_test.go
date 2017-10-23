package editor

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormattable(t *testing.T) {
	if assert.NoError(t, os.Setenv(forceColorFlag, "1")) {
		assert.True(t, Formattable(nil))
	}
}
