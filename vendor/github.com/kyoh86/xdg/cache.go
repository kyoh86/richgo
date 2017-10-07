package xdg

const (
	// CacheHomeEnv is the name of the environment variable holding a cache directory path.
	CacheHomeEnv = "XDG_CACHE_HOME"
)

// FindCacheFile finds a file from the XDG data directory.
// If one cannot be found, an error `ErrNotFound` be returned.
func FindCacheFile(rel ...string) (string, error) {
	return findFile([]string{CacheHome()}, rel)
}
