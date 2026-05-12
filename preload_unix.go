//go:build linux

package webp

import (
	"os"
	"path/filepath"
	"strings"
)

func preloadNative() {
	home, _ := os.UserHomeDir()
	libDir := filepath.Join(home, ".manga-visor", "webp-bin")
	libPath := filepath.Join(libDir, "libwebp.so")

	if _, err := os.Stat(libPath); err != nil {
		return
	}

	current := os.Getenv("LD_LIBRARY_PATH")
	if current == "" {
		os.Setenv("LD_LIBRARY_PATH", libDir)
	} else if !strings.Contains(current, libDir) {
		os.Setenv("LD_LIBRARY_PATH", libDir+string(os.PathListSeparator)+current)
	}
}
