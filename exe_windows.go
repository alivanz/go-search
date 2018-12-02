package search

import (
	"os"
	"strings"
)

func isExecutable(path string, info os.FileInfo) bool {
	return strings.HasSuffix(path, ".exe")
}
