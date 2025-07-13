package internals

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/darshan744/Trace/configs"
)

var dir string = configs.CommitDir

func LogAllCommit() {
	// read the filename of the dir
	dirs, err := os.ReadDir(dir)

	if err != nil {
		fmt.Printf("Error in reading commts for logging : %v", err)
		return
	}
	// if dirs is 0 then no commit happend
	if len(dirs) == 0 {
		fmt.Println("You haven't commited anything yet !!! ")
		return
	}

	for _, files := range dirs {
		// get the file path
		filePath := filepath.Join(dir, files.Name())
		// read the json content of the file(hashed commit entry) file name is only hashed
		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error in reading file content for logging commits %s : %v ", files.Name(), err)
			return
		}
		// returns the struct of the json
		commit := SingleCommit(fileContent)
		// logs the struct in a format
		logCommit(commit, files.Name())
	}
}

func SingleCommit(data []byte) (commit configs.Commit) {
	// converting byte json to struct
	err := json.Unmarshal(data, &commit)
	if err != nil {
		fmt.Printf("Error in converting bytes to json in commit : %v", err)
		return
	}
	return
}
func logCommit(commit configs.Commit, commitName string) {
	// splits the file name
	splits := strings.Split(commitName, ".json")
	// first one is the hash
	fmt.Println("Commit ", splits[0])
	// commits message
	fmt.Println("Commit message : ", commit.Message)
	// easy representation format logging
	formattedTime := commit.Time.Format("02 Jan, 2006 at 3:04 PM")
	fmt.Println("Timestamp : ", formattedTime)
}
