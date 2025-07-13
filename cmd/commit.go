/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/darshan744/Trace/configs"
	"github.com/darshan744/Trace/internals"
	"github.com/spf13/cobra"
)

var message string

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit your changes",
	Long:  `Create a check point of your project`,
	Run: func(cmd *cobra.Command, args []string) {
		var time time.Time = time.Now()
		// create a commit object
		commitMessage := configs.Commit{
			Message:     message,
			HashedFiles: internals.GetIndexFileData(),
			Time:        time,
		}
		// convert the to json -> returns []byte
		cont, err := json.Marshal(commitMessage)
		if err != nil {
			fmt.Printf("Error in converting the message to json : %v", err)
			return
		}
		// hash the []byte to get hashed value
		hashedContent := internals.Hash(cont)
		// convert byte content to hashed string like
		hexHashed := hex.EncodeToString(hashedContent[:])
		latestCommitObject := configs.LatestCommit{
			Latest: hexHashed,
		}
		latestCommitJson, errJ := json.Marshal(latestCommitObject)
		if errJ != nil {
			fmt.Printf("Error in converting latestCommitObject to JSON : %v", err)
			return
		}
		// write to hash as filename .json
		os.WriteFile(filepath.Join(configs.CommitDir, hexHashed)+".json", cont, 0644)
		// change the index to point to latest commit
		os.WriteFile(configs.IndexDir, latestCommitJson, 0644)
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
	commitCmd.Flags().StringVarP(&message, "message", "m", "", "Help for commiting")
}
