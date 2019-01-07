// +build !windows

package search

import (
	"os"
)

func isExecutable(path string, info os.FileInfo) bool {
	return info.Mode()&0111 != 0 && !info.IsDir()
}
