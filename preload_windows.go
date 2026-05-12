//go:build windows

package webp

import (
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"unsafe"
)

func preloadNative() {
	home, _ := os.UserHomeDir()
	libPath := filepath.Join(home, ".manga-visor", "webp-bin", "libwebp.dll")
	if _, err := os.Stat(libPath); err != nil {
		return
	}

	libDir := filepath.Dir(libPath)
	entries, _ := os.ReadDir(libDir)
	loadLibExW := syscall.NewLazyDLL("kernel32.dll").NewProc("LoadLibraryExW")

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(strings.ToLower(entry.Name()), ".dll") {
			continue
		}
		dllPath := filepath.Join(libDir, entry.Name())
		pathPtr, _ := syscall.UTF16PtrFromString(dllPath)
		loadLibExW.Call(uintptr(unsafe.Pointer(pathPtr)), 0, 0x100)
	}
}
