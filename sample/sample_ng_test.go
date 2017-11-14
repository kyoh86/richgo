// +build sample

package sample

import "testing"

func TestSampleNG(t *testing.T) {
	t.Fail()
	t.Log("It's not OK... :(")

	t.Run("SubtestNG", func(t *testing.T) {
		t.Fail()
		t.Log("It's also not OK... :(")
	})
}
