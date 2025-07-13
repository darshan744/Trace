package internals

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/darshan744/Trace/configs"
)

func RestoreHistory(commitName string) {
	// filepath to the json commit name
	filePath := filepath.Join(configs.CommitDir, commitName) + ".json"
	// get the stat of the file
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Theres no commit as : %s : %v", commitName, err)
		} else {
			fmt.Printf("Error in reading commit %v ", err)
		}
	}
	// read the json file
	readHashJson(fileInfo.Name())
}

func readHashJson(fileName string) {
	// reading the file
	file, err := os.ReadFile(filepath.Join(configs.CommitDir, fileName))
	if err != nil {
		fmt.Printf("Error in reading file for commit : %s  , err : %v ", fileName, err)
		return
	}
	var fileContents configs.Commit
	// converting json to struct
	err = json.Unmarshal(file, &fileContents)
	if err != nil {
		fmt.Printf("Error in decoding json : %v ", err)
		return
	}
	// restoring to repo
	restoreRepo(fileContents.HashedFiles)
}

func restoreRepo(files []configs.HashedFiles) {
	for _, file := range files {
		// file to be written to
		fileName := file.FileName
		// hashed file name where the content resides
		hashedContentFileName := file.Hash
		// content to be written
		content := readHashObject(hashedContentFileName)
		// write to file
		os.WriteFile(fileName, content, 0644)
	}
}

func readHashObject(hashFileName string) []byte {
	// hash file path
	hashPath := filepath.Join(configs.ObjectDir, hashFileName)
	// content
	hashFileContentAsByte, err := os.ReadFile(hashPath)
	if err != nil {
		fmt.Printf("Error in reading hash file %s : %v ", hashFileName, err)
		return nil
	}
	return hashFileContentAsByte
}
