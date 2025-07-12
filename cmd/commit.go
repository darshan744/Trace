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
			Message: message,
			Files:   getFiles(),
			Time:    time,
		}
		// convert the to json -> returns []byte
		cont, err := json.Marshal(commitMessage)
		if err != nil {
			fmt.Printf("Error in converting the message to json : %v", err)
			return
		}
		// hash the []byte to get hashed value
		hashedContent := internals.Hash(cont)
		// write to hash as filename .json
		os.WriteFile(filepath.Join(configs.CommitDir, hex.EncodeToString(hashedContent[:]))+".json", cont, 0644)
		os.WriteFile(configs.IndexDir, []byte(""), 0644)
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
	commitCmd.Flags().StringVarP(&message, "message", "m", "", "Help for commiting")
}

func getFiles() []string {
	byteData, err := os.ReadFile(configs.IndexDir)
	if err != nil {
		fmt.Printf("Error while getting index.json for commit message : %v ", err)
		return nil
	}
	var files configs.IndexFiles
	json.Unmarshal(byteData, &files)

	return files.Files
}
