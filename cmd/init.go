/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	mainDir = ".trace"
	subDirs = []string{
		"objects", "refs",
	}
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a trace repository",
	Long:  `Init initializes a empty trace repository `,
	Run: func(cmd *cobra.Command, args []string) {
		initializeRepo()
	},
}

func initializeRepo() {
	currentDir, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
		return
	}

	tracePath := filepath.Join(currentDir, mainDir)
	if dirExists(tracePath) {
		fmt.Println("Already Initialized", tracePath)
		return
	}

	err = os.Mkdir(".trace", 0755)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, subdir := range subDirs {
		subpath := filepath.Join(tracePath, subdir)

		if err := os.Mkdir(subpath, 0755); err != nil {
			fmt.Printf("Failed to create %s : %v ", subdir, err)
			return
		}
	}

	fmt.Println("Successfully initialized Repository")

}

func dirExists(dir string) bool {
	info, err := os.Stat(dir)
	return err == nil && info.IsDir()
}
func init() {
	rootCmd.AddCommand(initCmd)
}
