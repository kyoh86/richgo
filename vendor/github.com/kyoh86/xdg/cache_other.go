// +build !windows

package xdg

import "os"

// CacheHome returns a XDG cache directory (XDG_CACHE_HOME).
func CacheHome() string {
	return altHome(os.Getenv(CacheHomeEnv), ".cache")
}
