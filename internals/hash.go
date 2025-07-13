package internals

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"

	"github.com/darshan744/Trace/configs"
)

func HashFiles(files []string) {
	var hashedFiles []string = make([]string, 0)
	var hashFileObj []configs.HashedFiles = make([]configs.HashedFiles, 0)
	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			fmt.Printf("Error in hashing file %s : %v", file, err)
			return
		}
		// its a byte array of 20 (meaning its not a slice )
		var hashedValue [20]byte = Hash(content)
		// Reasong for [:]
		// EncodeToString expectes a slice not fixed size array
		// To get a slice we do [:]
		var hexCodeStringOfHash string = hex.EncodeToString(hashedValue[:])
		var fileObj configs.HashedFiles = configs.HashedFiles{FileName: file, Hash: hexCodeStringOfHash}
		hashFileObj = append(hashFileObj, fileObj)
		// stores somethign like "blob <contentLen>\0<filecontent>" git stores like this hence we do the same
		hashedFileDir := configs.ObjectDir + "/" + hexCodeStringOfHash
		hashedFiles = append(hashedFiles, hashedFileDir)
		os.WriteFile(hashedFileDir, content, 0644)
	}
	writeToIndex(hashFileObj)
}

func writeToIndex(fileDirs []configs.HashedFiles) {
	byteData, err := json.Marshal(fileDirs)

	if err != nil {
		fmt.Printf("Error while writing to index : %v", err)
	}
	os.WriteFile(configs.IndexDir, byteData, 0644)
}

func Hash(content []byte) [20]byte {
	return sha1.Sum(content)
}
