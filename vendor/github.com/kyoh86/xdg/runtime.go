package xdg

import "os"

const (
	// RuntimeDirEnv is the name of the environment variable holding a runtime directory path.
	RuntimeDirEnv = "XDG_RUNTIME_DIR"
)

// RuntimeDir returns XDG runtime directory.
func RuntimeDir() string {
	// XDG_RUNTIME_DIR
	return alternate(os.Getenv(RuntimeDirEnv), os.TempDir())
}

// FindRuntimeFile finds a file from the XDG runtime directory.
// If one cannot be found, an error `ErrNotFound` be returned.
func FindRuntimeFile(rel ...string) (string, error) {
	return findFile([]string{RuntimeDir()}, rel)
}
