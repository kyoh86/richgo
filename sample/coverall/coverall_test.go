package coverall

import "testing"

func TestCoverAll(t *testing.T) {
	if CoverAll() != "CoverAll" {
		t.Error("not covered all")
	}
}
