package test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/kyoh86/richgo/config"
	"github.com/stretchr/testify/require"
)

func TestE2E(t *testing.T) {
	// t.Fatal("not implemented")
	raws := bytes.Split(MustAsset("raw.txt"), []byte("\n"))
	exps := bytes.Split(MustAsset("colored.txt"), []byte("\n"))

	config.Load()
	editor := New()
	var expi int
	for _, raw := range raws {
		act, err := editor.Edit(string(raw) + "\n")
		require.NoError(t, err)
		for _, line := range strings.Split(act, "\n") {
			if len(line) > 0 {
				require.True(t, len(exps) > expi, "should have length more than", expi)
				require.Equal(t, string(exps[expi]), line, "at line ", expi)
				expi++
			}
		}
	}
}
