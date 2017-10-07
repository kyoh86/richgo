package xdg

const (
	// ConfigHomeEnv is the name of the environment variable holding a user configuration directory path.
	ConfigHomeEnv = "XDG_CONFIG_HOME"
	// ConfigDirsEnv is the name of the environment variable holding system configuration directory paths.
	ConfigDirsEnv = "XDG_CONFIG_DIRS"
)

// AllConfigDirs returns all XDG configuration directories.
func AllConfigDirs() []string {
	var dirs []string

	// XDG_CONFIG_HOME
	if home := ConfigHome(); home != "" {
		dirs = append(dirs, home)
	}

	// XDG_CONFIG_DIRS
	dirs = append(dirs, ConfigDirs()...)

	return dirs
}

// FindConfigFile finds a file from the XDG configuration directory.
// If one cannot be found, an error `ErrNotFound` be returned.
func FindConfigFile(rel ...string) (string, error) {
	return findFile(AllConfigDirs(), rel)
}
