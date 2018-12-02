package search

import (
	"os"
	"path/filepath"
)

func SearchFile(root string, criteria Criteria) []string {
	out := make([]string, 0)
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if criteria(path, info) {
			out = append(out, path)
		}
		return nil
	})
	return out
}
