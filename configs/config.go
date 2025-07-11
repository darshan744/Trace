package configs

import "path/filepath"

var (
	MainDir = ".trace"
	SubDirs = []string{
		"objects", "refs",
	}
	IgnoredDirs = map[string]bool{
		".git":   true,
		".trace": true,
	}
	ObjectDir = filepath.Join(MainDir, SubDirs[0])
)
