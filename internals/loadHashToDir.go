package internals

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/darshan744/Trace/configs"
)

func RestoreHistory(commitName string) {

	filePath := filepath.Join(configs.CommitDir, commitName) + ".json"
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Theres no commit as : %s : %v", commitName, err)
		} else {
			fmt.Printf("Error in reading commit %v ", err)
		}
	}
	writeToDir(fileInfo.Name())
}

func writeToDir(fileName string) {
	file, err := os.ReadFile(filepath.Join(configs.CommitDir, fileName))
	if err != nil {
		fmt.Printf("Error in reading file for commit : %s  , err : %v ", fileName, err)
		return
	}
	var fileContents configs.Commit

	err = json.Unmarshal(file, &fileContents)
	if err != nil {
		fmt.Printf("Error in decoding json : %v ", err)
		return
	}
	fmt.Printf("%s", fileContents)
}
