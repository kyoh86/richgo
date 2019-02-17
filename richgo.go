package main

import (
	"os"
	"os/exec"
	"syscall"

	"github.com/kyoh86/richgo/config"
	"github.com/kyoh86/richgo/editor"
	"github.com/kyoh86/richgo/editor/test"
)

const testFilterCmd = "testfilter"
const testCmd = "test"

type factoryFunc func() editor.Editor

func main() {
	config.Load()

	cmd, factory := parseArgs(os.Args[1:])
	tear := wrapCmd(cmd, factory)
	defer func() {
		if err := tear(); err != nil {
			panic(err)
		}
	}()

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

func parseArgs(args []string) (*exec.Cmd, factoryFunc) {
	if len(args) > 0 {
		switch args[0] {
		case testFilterCmd:
			return exec.Command("cat", "-"), test.New
		case testCmd:
			return exec.Command("go", args...), test.New
		}
	}
	return exec.Command("go"), nil
}

func wrapCmd(cmd *exec.Cmd, factory factoryFunc) func() error {
	if factory == nil {
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		return func() error { return nil }
	}
	if !editor.Formattable(os.Stderr) {
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		return func() error { return nil }
	}
	stderr := editor.Stream(os.Stderr, factory())
	stdout := editor.Stream(os.Stdout, factory())
	cmd.Stderr = stderr
	cmd.Stdout = stdout
	cmd.Stdin = os.Stdin
	return func() (retErr error) {
		if err := stderr.Close(); err != nil {
			retErr = err
		}
		if err := stdout.Close(); err != nil && retErr == nil {
			retErr = err
		}
		return
	}
}
