package files

import (
	"os"
	"path/filepath"
)

type GetWdFunc func() (string, error)

func GetRoot(getwd GetWdFunc) string {
	if getwd == nil {
		getwd = os.Getwd
	}
	wd, _ := getwd()
	return filepath.Dir(filepath.Dir(wd))
}

func GetRootDefault() string {
	return GetRoot(nil)
}