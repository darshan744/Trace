/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/darshan744/Trace/configs"
	"github.com/darshan744/Trace/internals"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds the specified file to staging",
	Long:  `Adds the file or complete directory to the staging area and then when commiting you can create a history version of it`,
	Run: func(cmd *cobra.Command, args []string) {
		if !internals.DirExists(configs.MainDir) {
			fmt.Println("This is not a trace repository ")
			return
		}
		stagedEntries := make([]string, 0)
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error in geting current Directory %v ", err)
			return
		}
		if len(args) == 1 && args[0] == "." {
			internals.Traverse(currentDir, &stagedEntries)
		} else {
			handleArgFiles(args, &stagedEntries)
		}
		internals.HashFiles(stagedEntries)
	},
}

func handleArgFiles(args []string, stagedEntries *[]string) {
	for _, arg := range args {
		info, err := os.Stat(arg)

		if err != nil {
			fmt.Printf("Error in reading file or directory %s : %v ", arg, err)
			return
		}

		if info.IsDir() {
			internals.Traverse(arg, stagedEntries)
		} else {
			*stagedEntries = append(*stagedEntries, arg)
		}
	}
}
func init() {
	rootCmd.AddCommand(addCmd)
}
