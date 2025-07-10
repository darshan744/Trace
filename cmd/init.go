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

		currentDir, err := os.Getwd()

		if err != nil {
			fmt.Println(err)
			return
		}
		if err = initializeRepo(currentDir); err != nil {
			fmt.Println("Error", err)
			return
		}

		fmt.Println("Initialized empty trace repository")
	},
}

func initializeRepo(currentDir string) error {

	tracePath := filepath.Join(currentDir, mainDir)
	if dirExists(tracePath) {
		return fmt.Errorf("Already Initialized %s ", tracePath)
	}

	err := os.Mkdir(tracePath, 0755)

	if err != nil {
		return fmt.Errorf("Could not create .trace %v", err)
	}

	for _, subdir := range subDirs {
		subpath := filepath.Join(tracePath, subdir)

		if err := os.Mkdir(subpath, 0755); err != nil {
			return fmt.Errorf("Failed to create %s : %v ", subdir, err)

		}
	}

	return nil
}

func dirExists(dir string) bool {
	info, err := os.Stat(dir)
	return err == nil && info.IsDir()
}
func init() {
	rootCmd.AddCommand(initCmd)
}
