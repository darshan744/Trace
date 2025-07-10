/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds the specified file to staging",
	Long:  `Adds the file or complete directory to the staging area and then when commiting you can create a history version of it`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 && args[0] == "." {

		} else {

		}
	},
}

func addFile() {

}

func traverse(dir string, stagedEntries []string) {
}
func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
