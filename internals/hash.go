package internals

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"

	"github.com/darshan744/Trace/configs"
)

func HashContent(files []string) {
	for _, file := range files {
		content, err := os.ReadFile(file)
		fmt.Println(file)
		if err != nil {
			fmt.Printf("Error in hashing file %s : %v", file, err)
			return
		}
		// for imitating git
		var contentLen int = len(content)
		var contentStr string = string(content)
		fmt.Println(contentStr)
		// sha1.Sum expects a []byte
		var blob []byte = []byte("blob " + strconv.Itoa(contentLen) + "\000" + contentStr)
		// its a byte array of 20 (meaning its not a slice )
		var hashedValue [20]byte = sha1.Sum(blob)
		// Reasong for [:]
		// EncodeToString expectes a slice not fixed size array
		// To get a slice we do [:]
		var hexCodeStringOfHash string = hex.EncodeToString(hashedValue[:])
		// stores somethign like "blob <contentLen>\0<filecontent>" git stores like this hence we do the same
		os.WriteFile(configs.ObjectDir+"/"+hexCodeStringOfHash, blob, 0644)
		fmt.Println("-------------------------------------------")
	}
}
