package editor

import (
	"os"

	"github.com/mattn/go-isatty"
)

const forceColorFlag = "RICHGO_FORCE_COLOR"

// Formattable judge whether a descriptor (like a os.Stdout, os.Stderr or *os.File...)
// is capable of colorization or not. The default behavior is to detect whether
// it is connected to a TTY, but may be overridden by setting the environment
// variable `RICHGO_FORCE_COLOR` to a non-empty value.
func Formattable(descriptor interface {
	Fd() uintptr
}) bool {
	return os.Getenv(forceColorFlag) != "" || isatty.IsTerminal(descriptor.Fd())
}
