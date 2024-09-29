package files

import (
	"os"
	"path/filepath"
)

func GetRoot() string {
	wd, _ := os.Getwd()
	return filepath.Dir(filepath.Dir(wd))
}