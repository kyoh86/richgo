package main

import (
	"io"
	"os"
	"os/exec"
	"syscall"

	"github.com/kyoh86/richgo/config"
	"github.com/kyoh86/richgo/editor"
	"github.com/kyoh86/richgo/editor/test"
)

const testFilterCmd = "testfilter"

type factoryFunc func() editor.Editor

var lps = map[string]factoryFunc{
	"test": test.New,
}

func main() {
	config.Load()

	var cmd *exec.Cmd
	var factory factoryFunc = editor.Parrot

	// without arguments
	switch len(os.Args) {
	case 0:
		panic("no arguments")
	case 1:
		cmd = exec.Command("go")
	default:
		// This is a bit of a special case. Somebody is already
		// running `go test` for us, and just wants us to prettify the
		// output.
		if os.Args[1] == testFilterCmd {
			cmd = exec.Command("cat", "-")
			factory = test.New
		} else {
			cmd = exec.Command("go", os.Args[1:]...)
			// select a wrapper with subcommand
			if f, ok := lps[os.Args[1]]; ok {
				factory = f
			}
		}
	}

	stderr := formatWriteCloser(os.Stderr, factory)
	defer stderr.Close()

	stdout := formatWriteCloser(os.Stdout, factory)
	defer stdout.Close()

	cmd.Stderr = stderr
	cmd.Stdout = stdout
	cmd.Stdin = os.Stdin

	switch err := cmd.Run().(type) {
	case nil:
		// noop
	default:
		panic(err)
	case *exec.ExitError:
		if waitStatus, ok := err.Sys().(syscall.WaitStatus); ok {
			defer os.Exit(waitStatus.ExitStatus())
		} else {
			panic(err)
		}
	}
}

func formatWriteCloser(wc io.WriteCloser, factory factoryFunc) io.WriteCloser {
	if editor.Formattable(os.Stderr) {
		return editor.Stream(wc, factory())
	}
	return editor.Stream(wc, editor.Parrot())
}
