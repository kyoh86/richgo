package xdg

import (
	"errors"
	"os"
	"path/filepath"
)

// ErrNotFound indicates that a file cannot be found in any directory.
var ErrNotFound = errors.New("not found")

func alternate(str, alt string) string {
	if str == "" {
		return alt
	}
	return str
}

func altHome(path string, suffix ...string) string {
	if path != "" {
		return path
	}
	home := os.Getenv("HOME")
	if home != "" {
		return filepath.Join(append([]string{home}, suffix...)...)
	}
	return ""
}

func findFile(paths []string, rel []string) (string, error) {
	for _, dir := range paths {
		if !filepath.IsAbs(dir) {
			// XDG Base Directory Specification supports only Absolute paths
			continue
		}

		fpath := filepath.Join(append([]string{dir}, rel...)...)
		if isFile(fpath) {
			return fpath, nil
		}
	}

	return "", ErrNotFound
}

func isFile(p string) bool {
	stat, err := os.Stat(p)
	return err == nil && !stat.IsDir()
}
