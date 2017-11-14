// +build sample

package sample

import "testing"

func TestSampleSkip(t *testing.T) {
	t.Skip()
}

func TestSampleSkipSub(t *testing.T) {
	t.Run("SubtestSkip", func(t *testing.T) {
		t.Skip()
		t.Log("It's maybe skipped... :(")
	})
}
