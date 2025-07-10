package cmd

import (
	"path/filepath"
	"testing"

	"github.com/darshan744/Trace/internals"
)

func TestInitializeRepo(t *testing.T) {

	tempDir := t.TempDir()

	err := initializeRepo(tempDir)

	if err != nil {
		t.Fatalf("initialization failed %v", err)
	}

	tracePath := filepath.Join(tempDir, ".trace")
	objPath := filepath.Join(tracePath, "objects")
	refPath := filepath.Join(tracePath, "refs")

	if !internals.DirExists(objPath) {
		t.Errorf("Creation of .trace/objects folder failed")
	}
	if !internals.DirExists(refPath) {
		t.Errorf("Creation of .trace/ref folder failed")
	}

	err = initializeRepo(tempDir)
	if err == nil {
		t.Errorf("Expected error but worked")
	}
}
