package test

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/kyoh86/richgo/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestE2E(t *testing.T) {
	mustAsset := func(t *testing, T, name string) []byte {
		file, err := Assets.File(name)
		if err != nil {
			t.Fatalf("failed to find %s: %s", name, err)
		}
		buf, err := ioutil.ReadAll(file)
		if err != nil {
			t.Fatalf("failed to load %s: %s", name, err)
		}
		return buf
	}
	raws := bytes.Split(mustAsset(t, "/sample/out_raw.txt"), []byte("\n"))
	exps := bytes.Split(mustAsset(t, "/sample/out_colored.txt"), []byte("\n"))

	config.Default()
	editor := New()
	var expi int
	for _, raw := range raws {
		act, err := editor.Edit(string(raw))
		require.NoError(t, err)
		for _, line := range strings.Split(act, "\n") {
			if len(line) > 0 {
				require.True(t, len(exps) > expi, "should have length more than", expi)
				assert.Equal(t, string(exps[expi]), line, "at line %d", expi+1)
				expi++
			}
		}
	}
}
