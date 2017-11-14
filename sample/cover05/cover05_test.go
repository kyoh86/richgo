// +build sample

package cover05

import "testing"

func TestCover05(t *testing.T) {
	case0 := Cover05(0)
	if case0 != "others" {
		t.Error("0 is not left out")
	}
}
