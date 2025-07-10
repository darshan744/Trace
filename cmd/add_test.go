package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/darshan744/Trace/configs"
)

func Test(t *testing.T) {
	tempDir := t.TempDir()

	tracePath := filepath.Join(tempDir, configs.MainDir)

	err := os.Mkdir(tracePath, 0755)

	if err != nil {
		t.Fatalf("Error in temp dir creattion %v", err)
	}

	oldDir, _ := os.Getwd()
	defer os.Chdir(oldDir)

	_ = os.Chdir(tempDir)

	args := []string{"a.txt", "b.txt"}
	staged := []string{}
	for _, file := range args {
		os.WriteFile(filepath.Join(tempDir, file), []byte("Hello text"), 0644)

	}
	handleArgFiles(args, &staged)

	if len(staged) != 2 {
		t.Fatalf("Didn't read correctly len > 2")
	}
	expectedPaths := map[string]bool{
		"a.txt": true,
		"b.txt": true,
	}
	for _, path := range staged {
		base := filepath.Base(path)
		if !expectedPaths[base] {
			t.Fatalf("Path didn't exist in creation")
		}
	}

}
