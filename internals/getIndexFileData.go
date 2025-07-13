package internals

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/darshan744/Trace/configs"
)

func GetIndexFileData() []configs.HashedFiles {
	byteData, err := os.ReadFile(configs.IndexDir)
	if err != nil {
		fmt.Printf("Error while getting index.json for commit message : %v ", err)
		return nil
	}
	var files []configs.HashedFiles
	json.Unmarshal(byteData, &files)

	return files
}
