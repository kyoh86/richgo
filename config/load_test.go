package config

import (
	"bytes"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/wacul/ptr"
	yaml "gopkg.in/yaml.v2"
)

func TestInGlobal(t *testing.T) {
	os.Setenv("HOME", "/")
	os.Setenv("GOPATH", "/")
	os.Setenv("GOROOT", "/foo/bar/path/not/exists")
	os.Setenv("PWD", "/home/kyoh86/go/src/github.com/kyoh86/richgo")
	err := os.Setenv(LocalOnlyEnvName, "0")
	if err != nil {
		t.Errorf("failed to set env: %s", err)
		t.FailNow()
	}
	sources := loadableSources()
	if len(sources) == 0 {
		t.Errorf("failed to get loadable sources")
	}
}

func TestInGlobalWithNotCoveredEnv(t *testing.T) {
	os.Clearenv()
	os.Setenv("HOME", "/home/kyoh86")
	os.Setenv("GOPATH", "/home/kyoh86/go")
	os.Unsetenv("GOROOT")
	os.Setenv("PWD", "/home/kyoh86/go/src/github.com/kyoh86/richgo")
	isDirForTest = func(path string) bool {
		return len(path) > 0
	}
	defer func() { isDirForTest = nil }()
	err := os.Setenv(LocalOnlyEnvName, "0")
	if err != nil {
		t.Errorf("failed to set env: %s", err)
		t.FailNow()
	}
	sources := loadableSources()
	if len(sources) == 0 {
		t.Errorf("failed to get loadable sources")
	}
}

func TestInLocal(t *testing.T) {
	err := os.Setenv(LocalOnlyEnvName, "1")
	if err != nil {
		t.Errorf("failed to set env: %s", err)
		t.FailNow()
	}
	sources := loadableSources()
	if len(sources) == 0 {
		t.Errorf("failed to get loadable sources")
	}
	wd, err := os.Getwd()
	if err != nil {
		t.Errorf("failed to get working directory: %q", err)
		t.FailNow()
	}
	for _, s := range sources {
		rel, err := filepath.Rel(wd, s)
		if err != nil {
			t.Errorf("failed to get relative path to %q from %q: %q", s, wd, err)
		}
		if strings.HasPrefix(rel, "..") {
			t.Errorf("expect that any source will be in local, but not (%q)", s)
		}
	}
}

func TestLoad(t *testing.T) {
	t.Run("no trick", func(t *testing.T) {
		if err := os.Setenv("PWD", "/home/kyoh86/go/src/github.com/kyoh86/richgo"); err != nil {
			t.Errorf("failed to set env: %s", err)
			t.FailNow()
		}

		if err := os.Setenv(LocalOnlyEnvName, "1"); err != nil {
			t.Errorf("failed to set env: %s", err)
			t.FailNow()
		}
		Load()
	})

	t.Run("with valid files", func(t *testing.T) {
		if err := os.Setenv("PWD", "/home/kyoh86/go/src/github.com/kyoh86/richgo"); err != nil {
			t.Errorf("failed to set env: %s", err)
			t.FailNow()
		}

		if err := os.Setenv(LocalOnlyEnvName, "1"); err != nil {
			t.Errorf("failed to set env: %s", err)
			t.FailNow()
		}

		loadForTest = func(p string) ([]byte, error) {
			return yaml.Marshal(&Config{
				CoverThreshold: ptr.Int(10),
			})
		}

		Load()
		if C.CoverThreshold == nil {
			t.Error("expect that a config loaded correctly but 'CoverThreshold' is nil")
			t.FailNow()
		}
		if *C.CoverThreshold != 10 {
			t.Errorf("expect that a 'CoverThreshold' is 10, but %d", *C.CoverThreshold)
		}
	})

	t.Run("with invalid file", func(t *testing.T) {
		loadForTest = func(p string) ([]byte, error) {
			return nil, errors.New("test error")
		}

		w := bytes.Buffer{}
		log.SetFlags(0)
		log.SetOutput(&w)
		Load()
		if !strings.HasPrefix(w.String(), "error reading from") {
			t.Error("expect that a Load func puts error in reading a file")
		}
	})

	t.Run("with invalid yaml", func(t *testing.T) {
		loadForTest = func(p string) ([]byte, error) {
			return []byte(`":`), nil
		}

		w := bytes.Buffer{}
		log.SetFlags(0)
		log.SetOutput(&w)
		Load()
		if !strings.HasPrefix(w.String(), "error unmarshaling yaml from") {
			t.Error("expect that a Load func puts error in unmarshaling a file")
		}
	})
}
