package fsx

import "os"

// FileExists returns true IFF a non-directory file exists at the provided path.
func FileExists(path string) bool {
	info, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	if info.IsDir() {
		return false
	}

	return true
}
