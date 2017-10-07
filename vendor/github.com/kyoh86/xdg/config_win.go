// +build windows

package xdg

import (
	"os"
	"path/filepath"
)

// ConfigHome returns a user XDG configuration directory (XDG_CONFIG_HOME).
func ConfigHome() string {
	return alternate(os.Getenv(ConfigHomeEnv), os.Getenv("APPDATA"))
}

// ConfigDirs returns system XDG configuration directories (XDG_CONFIG_DIRS).
func ConfigDirs() []string {
	// XDG_CONFIG_DIRS
	return filepath.SplitList(os.Getenv(ConfigDirsEnv))
}
