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
		"Trace":  true,
	}
	ObjectDir = filepath.Join(MainDir, SubDirs[0])
	IndexDir  = filepath.Join(MainDir, SubDirs[1], "index.json")
	CommitDir = filepath.Join(MainDir, SubDirs[2])
)

type Commit struct {
	Message     string        `json:"message"`
	HashedFiles []HashedFiles `json:"hashedFiles"`
	Time        time.Time     `json:"time"`
}

type IndexFiles struct {
	Files []string `json:"files"`
}

type LatestCommit struct {
	Latest string `json:"latest"`
}
type HashedFiles struct {
	FileName string `json:"fileName"`
	Hash     string `json:"hash"`
}
