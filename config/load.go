package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

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

	loaded Config
)

func loadableSources() []string {
	dirs := []string{}

	if dir, err := os.Getwd(); err == nil {
		dirs = append(dirs, dir)
	}

	localOnly := os.Getenv(LocalOnlyEnvName)
	if localOnly != "1" {
		dirs = appendIndirect(dirs, getEnvPath("GOPATH"))
		dirs = appendIndirect(dirs, getEnvPath("GOROOT"))
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

// Load configurations from file
func Load() {
	paths := loadableSources()
	c := &defaultConfig
	for _, p := range paths {
		data, err := ioutil.ReadFile(p)
		if err != nil {
			continue
		}
		if err := yaml.Unmarshal(data, &loaded); err != nil {
			continue
		}
		c = concatConfig(&loaded, c)
	}
	C = *actualConfig(c)
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
	if stat, err := os.Stat(envPath); err == nil && stat.IsDir() {
		return &envPath
	}
	return nil
}
