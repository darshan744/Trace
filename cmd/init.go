package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/darshan744/Trace/configs"
	"github.com/darshan744/Trace/internals"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a trace repository",
	Long:  `Init initializes a empty trace repository `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		currentDir, err := os.Getwd()

		if err != nil {
			fmt.Println(err)
			return
		}
		if err = initializeRepo(currentDir); err != nil {
			fmt.Println("Error", err)
			return
		}
		fmt.Println("Initialized empty trace repository", currentDir)
	},
}

func initializeRepo(currentDir string) error {

	tracePath := filepath.Join(currentDir, configs.MainDir)
	if internals.DirExists(tracePath) {
		return fmt.Errorf("Already Initialized %s ", tracePath)
	}

	err := os.Mkdir(tracePath, 0755)

	if err != nil {
		return fmt.Errorf("Could not create .trace %v", err)
	}

	for _, subdir := range configs.SubDirs {
		subpath := filepath.Join(tracePath, subdir)

		if err := os.Mkdir(subpath, 0755); err != nil {
			return fmt.Errorf("Failed to create %s : %v ", subdir, err)

		}
	}

	return nil
}

func init() {
	rootCmd.AddCommand(initCmd)
}
