// +build !windows

package xdg

import (
	"os"
	"path/filepath"
	"strings"
)

// DataHome returns a user XDG data directory (XDG_DATA_HOME).
func DataHome() string {
	return altHome(os.Getenv(DataHomeEnv), ".local", "share")
}

// DataDirs returns system XDG data directories (XDG_DATA_DIRS).
func DataDirs() []string {
	// XDG_DATA_DIRS
	xdgDirs := alternate(
		os.Getenv(DataDirsEnv),
		strings.Join([]string{
			filepath.Join("/", "usr", "local", "share"),
			filepath.Join("/", "usr", "share"),
		}, string(filepath.ListSeparator)),
	)
	return filepath.SplitList(xdgDirs)
}
