// +build windows

package xdg

import (
	"os"
	"path/filepath"
)

// DataHome returns a user XDG data directory (XDG_DATA_HOME).
func DataHome() string {
	return altHome(os.Getenv(DataHomeEnv), os.Getenv("LOCALAPPDATA"))
}

// DataDirs returns system XDG data directories (XDG_DATA_DIRS).
func DataDirs() []string {
	// XDG_DATA_DIRS
	return filepath.SplitList(os.Getenv(DataDirsEnv))
}
