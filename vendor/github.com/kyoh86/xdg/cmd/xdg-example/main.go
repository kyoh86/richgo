package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/kyoh86/xdg"
)

func main() {
	fmt.Println("ConfigHome() will get a user level configuration directory path: ", xdg.ConfigHome())

	os.Setenv(xdg.ConfigHomeEnv, "/.config")
	fmt.Println("...And `XDG_CONFIG_HOME` envar will be supported: ", xdg.ConfigHome())

	fmt.Println(strings.Repeat("-", 48))

	fmt.Println("ConfigDirs() will get system level configuration directories: ", xdg.ConfigDirs())

	os.Setenv(xdg.ConfigDirsEnv, "/.config:/.config2")
	fmt.Println("...And `XDG_CONFIG_DIRS` envar will be able to change it:", xdg.ConfigDirs())

	fmt.Println(strings.Repeat("-", 48))

	fmt.Println("DataHome() will get a user level data directory path: ", xdg.DataHome())

	os.Setenv(xdg.DataHomeEnv, "/.data")
	fmt.Println("...And `XDG_DATA_HOME` envar will be supported: ", xdg.DataHome())

	fmt.Println(strings.Repeat("-", 48))

	fmt.Println("DataDirs() will get system level data directories: ", xdg.DataDirs())

	os.Setenv(xdg.DataDirsEnv, "/.data:/.data2")
	fmt.Println("...And `XDG_DATA_DIRS` envar will be able to change it:", xdg.DataDirs())

	fmt.Println(strings.Repeat("-", 48))

	fmt.Println("CacheHome() will get a cache directory path: ", xdg.CacheHome())

	os.Setenv(xdg.CacheHomeEnv, "/.cache")
	fmt.Println("...And `XDG_CACHE_HOME` envar will be supported: ", xdg.CacheHome())

	fmt.Println(strings.Repeat("-", 48))

	fmt.Println("RuntimeDir() will get a runtime directory path: ", xdg.RuntimeDir())

	os.Setenv(xdg.RuntimeDirEnv, "/.runtime")
	fmt.Println("...And `XDG_RUNTIME_DIR` envar will be supported: ", xdg.RuntimeDir())
}
