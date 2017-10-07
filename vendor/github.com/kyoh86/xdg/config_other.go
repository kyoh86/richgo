// +build !windows

package xdg

import (
	"os"
	"path/filepath"
)

// ConfigHome returns a user XDG configuration directory (XDG_CONFIG_HOME).
func ConfigHome() string {
	return altHome(os.Getenv(ConfigHomeEnv), ".config")
}

// ConfigDirs returns system XDG configuration directories (XDG_CONFIG_DIRS).
func ConfigDirs() []string {
	// XDG_CONFIG_DIRS
	xdgDirs := alternate(
		os.Getenv(ConfigDirsEnv),
		filepath.Join("/", "etc", "xdg"),
	)
	return filepath.SplitList(xdgDirs)
}
