package test

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/kyoh86/richgo/config"
	"github.com/kyoh86/richgo/editor"
	"github.com/wacul/ptr"
)

// New will format lines as `go test` output
func New() editor.Editor {
	removals := make([]editor.RegexRepl, 0, len(config.C.Removals))
	for _, r := range config.C.Removals {
		removals = append(removals, editor.RegexRepl{
			Exp: regexp.MustCompile(r),
		})
	}
	return &test{
		additional: removals,
	}
}

// test through output raw.
type test struct {
	prevLineStyle *config.Style
	additional    []editor.RegexRepl
}

const noTestPattern = `[ \t]+\[(?:no test files|no tests to run)\]`

var (
	runhead    = regexp.MustCompile(`(?m)^=== RUN   Test.*`)
	passtail   = regexp.MustCompile(`(?m)^([ \t]*)--- PASS: Test.*`)
	skiptail   = regexp.MustCompile(`(?m)^([ \t]*)--- SKIP: Test.*`)
	failtail   = regexp.MustCompile(`(?m)^([ \t]*)--- FAIL: Test.*`)
	passlonely = regexp.MustCompile(`(?m)^PASS[ \t]*$`)
	faillonely = regexp.MustCompile(`(?m)^FAIL[ \t]*$`)

	okPath     = regexp.MustCompile(`(?m)^ok[ \t]+([^ \t]+)[ \t]*(?:[\d\.]+\w+|\(cached\))?[ \t]*(?:[ \t]+(coverage:[ \t]+\d+\.\d+% of statements)[ \t]*)?(?:` + noTestPattern + `)?$`)
	failPath   = regexp.MustCompile(`(?m)^FAIL[ \t]+[^ \t]+[ \t]+(?:[\d\.]+\w+|\[build failed\])$`)
	notestPath = regexp.MustCompile(`(?m)^\?[ \t]+[^ \t]+` + noTestPattern + `$`)

	coverage = regexp.MustCompile(`(?m)^coverage: ((\d+)\.\d)+% of statements?$`)

	filename   = regexp.MustCompile(`(?m)([^\s:]+\.go)((?::\d+){1,2})`)
	emptyline  = regexp.MustCompile(`(?m)^[ \t]*\r?\n`)
	importpath = regexp.MustCompile(`(?m)^# ([^ ]+)(?: \[[^ \[\]]+\])?$`)

	any = regexp.MustCompile(`.*`)
)

func (e *test) Edit(line string) (string, error) {
	var processed bool
	var style *config.Style
	edited := editor.Replaces(line,
		editor.RegexRepl{
			Exp: importpath,
			Func: func(s string) string {
				s = strings.TrimPrefix(s, `# `)
				processed = true
				style = config.C.BuildStyle
				return style.Apply(labels().Build() + s)
			},
		},

		editor.RegexRepl{
			Exp: runhead,
			Func: func(s string) string {
				if *config.C.LeaveTestPrefix {
					s = strings.TrimPrefix(s, `=== RUN   `)
				} else {
					s = strings.TrimPrefix(s, `=== RUN   Test`)
				}
				floors := strings.Split(s, `/`)
				processed = true

				clone := *config.C.StartStyle
				clone.Hide = ptr.Bool(false)
				style = &clone
				return config.C.StartStyle.Apply(labels().Start() + strings.Repeat("  ", len(floors)-1) + s)
			},
		},
		editor.RegexRepl{
			Exp: passtail,
			Func: func(s string) string {
				s = strings.TrimLeft(s, " ")
				if *config.C.LeaveTestPrefix {
					s = strings.TrimPrefix(s, `--- PASS: `)
				} else {
					s = strings.TrimPrefix(s, `--- PASS: Test`)
				}
				floors := strings.Split(s, `/`)
				processed = true
				style = config.C.PassStyle
				return style.Apply(labels().Pass() + strings.Repeat("  ", len(floors)-1) + s)
			},
		},
		editor.RegexRepl{
			Exp: failtail,
			Func: func(s string) string {
				s = strings.TrimLeft(s, " ")
				if *config.C.LeaveTestPrefix {
					s = strings.TrimPrefix(s, `--- FAIL: `)
				} else {
					s = strings.TrimPrefix(s, `--- FAIL: Test`)
				}
				floors := strings.Split(s, `/`)
				processed = true
				style = config.C.FailStyle
				return style.Apply(labels().Fail() + strings.Repeat("  ", len(floors)-1) + s)
			},
		},
		editor.RegexRepl{
			Exp: skiptail,
			Func: func(s string) string {
				s = strings.TrimLeft(s, " ")
				if *config.C.LeaveTestPrefix {
					s = strings.TrimPrefix(s, `--- SKIP: `)
				} else {
					s = strings.TrimPrefix(s, `--- SKIP: Test`)
				}
				floors := strings.Split(s, `/`)
				processed = true
				style = config.C.SkipStyle
				return style.Apply(labels().Skip() + strings.Repeat("  ", len(floors)-1) + s)
			},
		},

		editor.RegexRepl{
			Exp: okPath,
			Func: func(s string) string {
				matches := okPath.FindStringSubmatch(s)
				processed = true
				style = config.C.PassStyle

				ret := style.Apply(labels().Pass() + strings.Join(matches[1:3], " "))
				if len(matches) == 4 {
					ret += "\n" + matches[3]
				}
				return ret
			},
		},
		editor.RegexRepl{
			Exp: failPath,
			Func: func(s string) string {
				s = strings.TrimPrefix(strings.TrimLeft(s, " \t"), `FAIL`)
				processed = true
				style = config.C.FailStyle
				return style.Apply(labels().Fail() + s)
			},
		},
		editor.RegexRepl{
			Exp: notestPath,
			Func: func(s string) string {
				s = strings.TrimLeft(s, " \t?")
				processed = true
				style = config.C.SkipStyle
				return style.Apply(labels().Skip() + s)
			},
		},

		editor.RegexRepl{
			Exp: coverage,
			Func: func(s string) string {
				matches := coverage.FindStringSubmatch(s)
				fill, err := strconv.Atoi(matches[2])
				if err != nil {
					panic(err)
				}
				s = fmt.Sprintf("%s%% [%s%s]", matches[1], strings.Repeat("#", fill/10), strings.Repeat("_", 10-fill/10))
				coverStyle := config.C.CoveredStyle
				if fill < *config.C.CoverThreshold {
					coverStyle = config.C.UncoveredStyle
				}
				processed = true
				return coverStyle.Apply(labels().Cover()+s) + "\n"
			},
		},

		editor.RegexRepl{
			Exp: passlonely,
			Func: func(s string) string {
				processed = true
				return config.C.PassPackageStyle.Apply("PASS")
			},
		},
		editor.RegexRepl{
			Exp: faillonely,
			Func: func(s string) string {
				processed = true
				return config.C.FailPackageStyle.Apply("FAIL")
			},
		},

		editor.RegexRepl{
			Exp: filename,
			Func: func(s string) string {
				return filename.ReplaceAllString(s, config.C.FileStyle.Apply(`$1`)+config.C.LineStyle.Apply(`$2`)) + e.prevLineStyle.ANSI().String()
			},
		},
	)

	edited = editor.Replaces(edited, e.additional...)

	if !processed {
		edited = editor.Replaces(edited,
			editor.RegexRepl{
				Exp: any,
				Func: func(s string) string {
					if s == "" {
						return ""
					}
					return e.prevLineStyle.Apply(labels().Anonym() + s)
				},
			},
		)
	}

	edited = editor.Replaces(edited,
		editor.RegexRepl{Exp: emptyline},
	)

	if style != nil {
		e.prevLineStyle = style
	}

	return edited, nil
}
