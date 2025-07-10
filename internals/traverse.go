package internals

import (
	"fmt"
	"os"
	"path"

	"github.com/darshan744/Trace/configs"
)

func Traverse(dir string, stagedEntries *[]string) {
	entries, err := os.ReadDir(dir)

	if err != nil {
		fmt.Printf("Error in staging files %v ", err)
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			if !configs.IgnoredDirs[entry.Name()] {
				Traverse(path.Join(dir, entry.Name()), stagedEntries)
			}
		} else {
			*stagedEntries = append(*stagedEntries, path.Join(dir, entry.Name()))
		}
	}
}
