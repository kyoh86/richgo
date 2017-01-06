package editor

import (
	"github.com/mattn/go-isatty"
)

// Formattable judge whether a descriptor (like a os.Stdout, os.Stderr or *os.File...) has the TTY or not.
func Formattable(descriptor interface {
	Fd() uintptr
}) bool {
	return isatty.IsTerminal(descriptor.Fd())
}
