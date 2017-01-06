package editor

import "regexp"

// RegexRepl : replacer
type RegexRepl struct {
	Exp  *regexp.Regexp
	Repl string
	Func func(string) string
}

// Replaces string with
func Replaces(str string, rs ...RegexRepl) string {
	for _, r := range rs {
		x := r.Exp
		switch {
		case r.Func != nil:
			str = x.ReplaceAllStringFunc(str, r.Func)
		default:
			str = x.ReplaceAllString(str, r.Repl)
		}
	}
	return str
}
