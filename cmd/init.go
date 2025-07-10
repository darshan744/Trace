/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"

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
	alreadyInitialized := isInitializedAlready(currentDir)

	if alreadyInitialized {
		fmt.Println("Repo Already Initialized")
		return
	}

	err = os.Mkdir(".trace", 0755)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successfully initialized a repository")
}

func isInitializedAlready(currDir string) bool {
	parentExist, err := dirExists(currDir)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	currentPath := strings.Join([]string{currDir, mainDir}, string(os.PathListSeparator))

	childExist, err := dirExists(currentPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return parentExist && childExist
}
func dirExists(dir string) (bool, error) {
	_, err := os.Stat(dir)

	if err == nil {
		return true, nil
	} else if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}

	return false, err
}
func init() {
	rootCmd.AddCommand(initCmd)
}
