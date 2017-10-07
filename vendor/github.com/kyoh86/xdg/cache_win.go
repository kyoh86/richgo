// +build windows

package xdg

import (
	"os"
	"path/filepath"
)

// CacheHome returns a XDG cache directory (XDG_CACHE_HOME).
func CacheHome() string {
	home := os.Getenv(CacheHomeEnv)
	if home == "" {
		home = os.Getenv("LOCALAPPDATA")
		if home != "" {
			home = filepath.Join(home, "cache")
		}
	}
	return home
}
