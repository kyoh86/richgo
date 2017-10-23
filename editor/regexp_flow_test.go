package editor

import (
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaces(t *testing.T) {
	assert.Equal(t,
		"A!A!A!-D!E!L!A!X!-C!C!C!",
		Replaces(
			"aaabbbccc",
			RegexRepl{
				Exp:  regexp.MustCompile(`b{3}`),
				Repl: "-delax-",
			},
			RegexRepl{
				Exp: regexp.MustCompile(`\w`),
				Func: func(s string) string {
					return strings.ToUpper(s)
				},
			},
			RegexRepl{
				Exp:  regexp.MustCompile(`\w`),
				Repl: "x",
				Func: func(s string) string {
					return s + "!"
				},
			},
		),
	)
}
