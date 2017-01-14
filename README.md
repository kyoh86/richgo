# Rich-Go

Rich-Go will enrich `go test` outputs with text decorations.

[![asciicast](https://asciinema.org/a/99292.png)](https://asciinema.org/a/99292)

# Installation

```
go get -u github.com/kyoh86/richgo
```

# Usage

```sh
richgo test ./...
```

## Configure

### Configuration file paths

It's possible to change styles with the preference file.
Rich-Go loads preferences from the files in the following order.

* `${CWD}/.richstyle`
* `${CWD}/.richstyle.yaml`
* `${CWD}/.richstyle.yml`
* `${GOPATH}/.richstyle`
* `${GOPATH}/.richstyle.yaml`
* `${GOPATH}/.richstyle.yml`
* `${GOROOT}/.richstyle`
* `${GOROOT}/.richstyle.yaml`
* `${GOROOT}/.richstyle.yml`
* `${HOME}/.richstyle`
* `${HOME}/.richstyle.yaml`
* `${HOME}/.richstyle.yml`

Setting the environment variable `RICHGO_LOCAL` to 1, Rich-Go loads only `${CWD}/.richstyle*`.

### Configuration file format

**Now Rich-Go supports only YAML formatted.**

```yaml
# Type of the label that notes a kind of each lines.
labelType: (long | short | none)

# Style of "Build" lines.
buildStyle:
  # Hide lines
  hide: (true | false)
  # Bold or increased intensity.
  bold: (true | false)
  faint: (true | false)
  italic: (true | false)
  underline: (true | false)
  blinkSlow: (true | false)
  blinkRapid: (true | false)
  # Swap the foreground color and background color.
  inverse: (true | false)
  conceal: (true | false)
  crossOut: (true | false)
  frame: (true | false)
  encircle: (true | false)
  overline: (true | false)
  # Fore-color of text
  foreground: (#xxxxxx | rgb(0-256,0-256,0-256) | rgb(0x00-0xFF,0x00-0xFF,0x00-0xFF) | (name of colors))
  # Back-color of text
  background: # Same format as `foreground`

# Style of the "Start" lines.
startStyle:
  # Same format as `buildStyle`

# Style of the "Pass" lines.
passStyle:
  # Same format as `buildStyle`

# Style of the "Fail" lines.
failStyle:
  # Same format as `buildStyle`

# Style of the "Skip" lines.
skipStyle:
  # Same format as `buildStyle`

# Style of the "File" lines.
fileStyle:
  # Same format as `buildStyle`

# Style of the "Line" lines.
lineStyle:
  # Same format as `buildStyle`

# If you want to delete lines, write the regular expressions.
removals:
  - (regexp)
```

### Line categories

Rich-Go separate the output-lines in following categories.

* Build
  When the Go fails to build, it prints errors like this:

  <pre><code># github.com/kyoh86/richgo/sample/buildfail
  sample/buildfail/buildfail_test.go:6: t.Foo undefined (type testing.T has no field or method Foo)</code></pre>

* Start
  In the top of test, Go prints that name like this:

  <pre><code>=== RUN   TestSampleOK/SubtestOK</code></pre>

* Pass
  When a test is successed, Go prints that name like this:

  <pre><code>    ---PASS: TestSampleOK/SubtestOK</code></pre>

* Fail
  When a test is failed, Go prints that name like this:

  <pre><code>--- FAIL: TestSampleNG (0.00s)
  sample_ng_test.go:9: It's not OK... :(</code></pre>

* Skip
  If there is no test files in directory or a test is skipped, Go prints that path or the name like this:

  <pre><code>--- SKIP: TestSampleSkip (0.00s)
  sample_skip_test.go:6:
?     github.com/kyoh86/richgo/sample/notest  [no test files]</code></pre>

Each categories can be styled seperately.

### Label types

* Long
  * Build: "BUILD"
  * Start: "START"
  * Pass: "PASS"
  * Fail: "FAIL"
  * Skip: "SKIP"

* Short
  * Build: "!!"
  * Start: ">"
  * Pass: "o"
  * Fail: "x"
  * Skip: "-"

* None
  Rich-Go will never output labels.

### Default

```yaml
labelType: long
buildStyle:
  bold: true
  foreground: yellow
startStyle:
  foreground: lightBlack
passStyle:
  foreground: green
failStyle:
  bold: true
  foreground: red
skipStyle:
  foreground: lightBlack
fileStyle:
  foreground: cyan
lineStyle:
  foreground: magenta
```

# License

[The Unlicense](http://unlicense.org)
