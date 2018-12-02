package search

import (
	"os"
	"path/filepath"
	"strings"
)

// Criteria Unit filter
type Criteria func(path string, info os.FileInfo) bool

//Must join multiple criteria, must be all valid
func Must(criterias ...Criteria) Criteria {
	return func(path string, info os.FileInfo) bool {
		for _, criteria := range criterias {
			if !criteria(path, info) {
				return false
			}
		}
		return true
	}
}

// Any join multiple criteria, invalid if all criterias unfulfilled.
// Note that zero criterias always valid
func Any(criterias ...Criteria) Criteria {
	return func(path string, info os.FileInfo) bool {
		if len(criterias) == 0 {
			return true
		}
		for _, criteria := range criterias {
			if criteria(path, info) {
				return true
			}
		}
		return false
	}
}

func IsFolder(path string, info os.FileInfo) bool {
	return info.Mode().IsDir()
}

func IsExecutable(path string, info os.FileInfo) bool {
	return isExecutable(path, info)
}

func Filename(filename string) Criteria {
	return func(path string, info os.FileInfo) bool {
		return filepath.Base(path) == filename
	}
}

func FilenamePrefix(prefix string) Criteria {
	return func(path string, info os.FileInfo) bool {
		return strings.HasPrefix(
			filepath.Base(path),
			prefix,
		)
	}
}
