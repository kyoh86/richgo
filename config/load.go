package config

import (
	"go/build"
	"log"
	"os"
	"path/filepath"

	"github.com/kyoh86/xdg"
	yaml "gopkg.in/yaml.v2"
)

const (
	// Filename is filename of configurations for this app.
	Filename = ".richstyle"
	// LocalOnlyEnvName is the name of environment variable
	// to stop searching configuration files excepting current directory.
	LocalOnlyEnvName = "RICHGO_LOCAL"
)

var (
	// Extensions is extension choices of configurations for this app.
	Extensions = []string{
		"",
		".yaml",
		".yml",
	}
	// C is global configuration
	C Config
)

func loadableSources() []string {
	dirs := []string{}

	if dir, err := os.Getwd(); err == nil {
		dirs = append(dirs, dir)
	}

	localOnly := os.Getenv(LocalOnlyEnvName)
	if localOnly != "1" {
		dirs = append(dirs, build.Default.GOPATH)
		dirs = appendIndirect(dirs, getEnvPath("GOROOT"))
		if xdgHome := xdg.ConfigHome(); xdgHome != "" {
			dirs = append(dirs, xdgHome)
		}
		dirs = appendIndirect(dirs, getEnvPath("HOME"))
	}

	paths := make([]string, 0, len(dirs)*len(Extensions))
	for _, d := range dirs {
		for _, e := range Extensions {
			paths = append(paths, filepath.Join(d, Filename+e))
		}
	}
	return paths
}

var loadForTest func(path string) ([]byte, error)

func load(path string) ([]byte, error) {
	if loadForTest != nil {
		return loadForTest(path)
	}
	return os.ReadFile(path)
}

// Load configurations from file
func Load() {
	paths := loadableSources()
	c := &defaultConfig
	for _, p := range paths {
		data, err := load(p)
		if err != nil {
			if !os.IsNotExist(err) {
				log.Println("error reading from", p, ": ", err)
			}
			continue
		}
		var loaded Config
		if err := yaml.Unmarshal(data, &loaded); err != nil {
			log.Println("error unmarshaling yaml from", p, ": ", err)
			continue
		}
		c = concatConfig(&loaded, c)
	}
	C = *actualConfig(c)
}

// Default is the default configuration
func Default() {
	C = *actualConfig(&defaultConfig)
}

func appendIndirect(arr []string, ptr *string) []string {
	if ptr != nil {
		return append(arr, *ptr)
	}
	return arr
}

func getEnvPath(envName string) *string {
	envPath := os.Getenv(envName)
	if envPath == "" {
		return nil
	}
	if isDir(envPath) {
		return &envPath
	}
	return nil
}

var isDirForTest func(path string) bool

func isDir(path string) bool {
	if isDirForTest != nil {
		return isDirForTest(path)
	}
	if stat, err := os.Stat(path); err == nil && stat.IsDir() {
		return true
	}
	return false
}
