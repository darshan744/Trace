package configs

import (
	"path/filepath"
	"time"
)

var (
	MainDir = ".trace"
	SubDirs = []string{
		"objects", "refs", "commits",
	}
	IgnoredDirs = map[string]bool{
		".git":   true,
		".trace": true,
	}
	ObjectDir = filepath.Join(MainDir, SubDirs[0])
	IndexDir  = filepath.Join(MainDir, SubDirs[1], "index.json")
	CommitDir = filepath.Join(MainDir, SubDirs[2])
)

type Commit struct {
	Message string    `json:"message"`
	Files   []string  `json:"files"`
	Time    time.Time `json:"time"`
}

type IndexFiles struct {
	Files []string `json:"files"`
}
