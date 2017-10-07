package xdg

const (
	// DataHomeEnv is the name of the environment variable holding a user data directory path.
	DataHomeEnv = "XDG_DATA_HOME"
	// DataDirsEnv is the name of the environment variable holding system data directory paths.
	DataDirsEnv = "XDG_DATA_DIRS"
)

// AllDataDirs returns all XDG data directories.
func AllDataDirs() []string {
	var dirs []string

	// XDG_DATA_HOME
	if home := DataHome(); home != "" {
		dirs = append(dirs, home)
	}

	// XDG_DATA_DIRS
	dirs = append(dirs, DataDirs()...)

	return dirs
}

// FindDataFile finds a file from the XDG data directory.
// If one cannot be found, an error `ErrNotFound` be returned.
func FindDataFile(rel ...string) (string, error) {
	return findFile(AllDataDirs(), rel)
}
