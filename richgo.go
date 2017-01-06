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
		cmd = exec.Command("go", os.Args[1:]...)
		// select a wrapper with subcommand
		if f, ok := lps[os.Args[1]]; ok {
			factory = f
		}
	}

	var stderr io.WriteCloser
	if editor.Formattable(os.Stderr) {
		stderr = editor.Stream(os.Stderr, factory())
	} else {
		stderr = editor.Stream(os.Stderr, editor.Parrot())
	}
	defer stderr.Close()

	var stdout io.WriteCloser
	if editor.Formattable(os.Stdout) {
		stdout = editor.Stream(os.Stdout, factory())
	} else {
		stdout = editor.Stream(os.Stdout, editor.Parrot())
	}
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
