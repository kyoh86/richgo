// +build sample

package sample

import (
	"fmt"
	"testing"
)

func TestSampleOK(t *testing.T) {
	t.Log("It's OK!")

	t.Run("SubtestOK", func(t *testing.T) {
		fmt.Println("time:2017-01-01T01:01:01+09:00")
		t.Log("It's also OK!")
	})
}
