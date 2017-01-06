package sample

import (
	"testing"
	"time"
)

func TestSampleTimeout(t *testing.T) {
	t.Skip() //COMMENT: Comment-out this line to get sample
	t.Run("SubtestTimeout", func(t *testing.T) {
		// Trying a command `\go test ./sample/... -test.timeout 1s`, fail with timeout
		time.Sleep(3 * time.Second)
	})
}
