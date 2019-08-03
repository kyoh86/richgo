package test

import (
	"time"

	"github.com/deadcheat/goblet"
)

//go:generate goblet -g --ignore-dotfiles -o output_test.go -p test ../../sample/out_colored.txt ../../sample/out_raw.txt

// Assets a generated file system
var Assets = goblet.NewFS(
	map[string][]string{},
	map[string]*goblet.File{
		"/sample/out_colored.txt": goblet.NewFile("/sample/out_colored.txt", _Assets64295af663b532b72d17cd22db298a3d276a5583, 0x1a4, time.Unix(1564845186, 1564845186462872869)),
		"/sample/out_raw.txt":     goblet.NewFile("/sample/out_raw.txt", _Assetsed805497a9c9001c8e8baf583bfdc6893d5fe8f6, 0x1a4, time.Unix(1549785466, 1549785466313613757)),
	},
)

// binary data
var (
	_Assets64295af663b532b72d17cd22db298a3d276a5583 = []byte("\x1b[33m\x1b[1mBUILD| github.com/kyoh86/richgo/sample/buildfail\x1b[0m\n\x1b[33m\x1b[1m     | \x1b[36msample/buildfail/buildfail_test.go\x1b[0m:\x1b[35m6\x1b[0m:10: t.Foo undefined (type *testing.T has no field or method Foo)\x1b[0m\n\x1b[90mSTART| SampleNG\x1b[0m\n\x1b[90mSTART|   SampleNG/SubtestNG\x1b[0m\n\x1b[31m\x1b[1mFAIL | SampleNG (0.00s)\x1b[0m\n\x1b[31m\x1b[1m     | \t\x1b[36msample_ng_test.go\x1b[0m:\x1b[35m9\x1b[0m: It's not OK... :(\x1b[0m\n\x1b[31m\x1b[1mFAIL |   SampleNG/SubtestNG (0.00s)\x1b[0m\n\x1b[31m\x1b[1m     |     \t\x1b[36msample_ng_test.go\x1b[0m:\x1b[35m13\x1b[0m: It's also not OK... :(\x1b[0m\n\x1b[90mSTART| SampleOK\x1b[0m\n\x1b[90mSTART|   SampleOK/SubtestOK\x1b[0m\n\x1b[90m     | time:2017-01-01T01:01:01+09:00\x1b[0m\n\x1b[32mPASS | SampleOK (0.00s)\x1b[0m\n\x1b[32m     | \t\x1b[36msample_ok_test.go\x1b[0m:\x1b[35m11\x1b[0m: It's OK!\x1b[0m\n\x1b[32mPASS |   SampleOK/SubtestOK (0.00s)\x1b[0m\n\x1b[32m     |     \t\x1b[36msample_ok_test.go\x1b[0m:\x1b[35m15\x1b[0m: It's also OK!\x1b[0m\n\x1b[90mSTART| SampleSkip\x1b[0m\n\x1b[90mSKIP | SampleSkip (0.00s)\x1b[0m\n\x1b[90m     | \t\x1b[36msample_skip_test.go\x1b[0m:\x1b[35m6\x1b[0m: \x1b[0m\n\x1b[90mSTART| SampleSkipSub\x1b[0m\n\x1b[90mSTART|   SampleSkipSub/SubtestSkip\x1b[0m\n\x1b[32mPASS | SampleSkipSub (0.00s)\x1b[0m\n\x1b[90mSKIP |   SampleSkipSub/SubtestSkip (0.00s)\x1b[0m\n\x1b[90m     |     \t\x1b[36msample_skip_test.go\x1b[0m:\x1b[35m11\x1b[0m: \x1b[0m\n\x1b[90mSTART| SampleTimeout\x1b[0m\n\x1b[90mSTART|   SampleTimeout/SubtestTimeout\x1b[0m\n\x1b[32mPASS | SampleTimeout (3.00s)\x1b[0m\n\x1b[32mPASS |   SampleTimeout/SubtestTimeout (3.00s)\x1b[0m\n\x1b[33m\x1b[1mCOVER| 0.0% [__________]\x1b[0m\n\x1b[31m\x1b[1mFAIL | \tgithub.com/kyoh86/richgo/sample\t3.014s\x1b[0m\n\x1b[31m\x1b[1mFAIL | \tgithub.com/kyoh86/richgo/sample/buildfail [build failed]\x1b[0m\n\x1b[90mSTART| Cover05\x1b[0m\n\x1b[32mPASS | Cover05 (0.00s)\x1b[0m\n\x1b[32mCOVER| 50.0% [#####_____]\x1b[0m\n\x1b[32mPASS | github.com/kyoh86/richgo/sample/cover05 0.015s\x1b[0m\n\x1b[32mCOVER| 50.0% [#####_____]\x1b[0m\n\x1b[90mSTART| CoverAll\x1b[0m\n\x1b[32mPASS | CoverAll (0.00s)\x1b[0m\n\x1b[32mCOVER| 100.0% [##########]\x1b[0m\n\x1b[32mPASS | github.com/kyoh86/richgo/sample/coverall 0.008s\x1b[0m\n\x1b[32mCOVER| 100.0% [##########]\x1b[0m\n\x1b[32m     | testing: warning: no tests to run\x1b[0m\n\x1b[33m\x1b[1mCOVER| 0.0% [__________]\x1b[0m\n\x1b[32mPASS | github.com/kyoh86/richgo/sample/emptytest 0.013s\x1b[0m\n\x1b[33m\x1b[1mCOVER| 0.0% [__________]\x1b[0m\n\x1b[90mSTART| Nocover\x1b[0m\n\x1b[32mPASS | Nocover (0.00s)\x1b[0m\n\x1b[32m     | \t\x1b[36mnocover_test.go\x1b[0m:\x1b[35m6\x1b[0m: accept\x1b[0m\n\x1b[33m\x1b[1mCOVER| 0.0% [__________]\x1b[0m\n\x1b[32mPASS | github.com/kyoh86/richgo/sample/nocover 0.007s\x1b[0m\n\x1b[33m\x1b[1mCOVER| 0.0% [__________]\x1b[0m\n\x1b[90mSKIP | github.com/kyoh86/richgo/sample/notest\t[no test files]\x1b[0m\n")
	_Assetsed805497a9c9001c8e8baf583bfdc6893d5fe8f6 = []byte("# github.com/kyoh86/richgo/sample/buildfail\nsample/buildfail/buildfail_test.go:6:10: t.Foo undefined (type *testing.T has no field or method Foo)\n=== RUN   TestSampleNG\n=== RUN   TestSampleNG/SubtestNG\n--- FAIL: TestSampleNG (0.00s)\n\tsample_ng_test.go:9: It's not OK... :(\n    --- FAIL: TestSampleNG/SubtestNG (0.00s)\n    \tsample_ng_test.go:13: It's also not OK... :(\n=== RUN   TestSampleOK\n=== RUN   TestSampleOK/SubtestOK\ntime:2017-01-01T01:01:01+09:00\n--- PASS: TestSampleOK (0.00s)\n\tsample_ok_test.go:11: It's OK!\n    --- PASS: TestSampleOK/SubtestOK (0.00s)\n    \tsample_ok_test.go:15: It's also OK!\n=== RUN   TestSampleSkip\n--- SKIP: TestSampleSkip (0.00s)\n\tsample_skip_test.go:6: \n=== RUN   TestSampleSkipSub\n=== RUN   TestSampleSkipSub/SubtestSkip\n--- PASS: TestSampleSkipSub (0.00s)\n    --- SKIP: TestSampleSkipSub/SubtestSkip (0.00s)\n    \tsample_skip_test.go:11: \n=== RUN   TestSampleTimeout\n=== RUN   TestSampleTimeout/SubtestTimeout\n--- PASS: TestSampleTimeout (3.00s)\n    --- PASS: TestSampleTimeout/SubtestTimeout (3.00s)\nFAIL\ncoverage: 0.0% of statements\nFAIL\tgithub.com/kyoh86/richgo/sample\t3.014s\nFAIL\tgithub.com/kyoh86/richgo/sample/buildfail [build failed]\n=== RUN   TestCover05\n--- PASS: TestCover05 (0.00s)\nPASS\ncoverage: 50.0% of statements\nok  \tgithub.com/kyoh86/richgo/sample/cover05\t0.015s\tcoverage: 50.0% of statements\n=== RUN   TestCoverAll\n--- PASS: TestCoverAll (0.00s)\nPASS\ncoverage: 100.0% of statements\nok  \tgithub.com/kyoh86/richgo/sample/coverall\t0.008s\tcoverage: 100.0% of statements\ntesting: warning: no tests to run\nPASS\ncoverage: 0.0% of statements\nok  \tgithub.com/kyoh86/richgo/sample/emptytest\t0.013s\tcoverage: 0.0% of statements [no tests to run]\n=== RUN   TestNocover\n--- PASS: TestNocover (0.00s)\n\tnocover_test.go:6: accept\nPASS\ncoverage: 0.0% of statements\nok  \tgithub.com/kyoh86/richgo/sample/nocover\t0.007s\tcoverage: 0.0% of statements\n?   \tgithub.com/kyoh86/richgo/sample/notest\t[no test files]\n")
)
