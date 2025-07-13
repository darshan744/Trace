package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var commitName string

// loadCmd represents the load command
var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Move to a specified commit",
	Long:  `This command helps the user to move out to the specified commit done by the user `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(commitName)
	},
}

func init() {
	rootCmd.AddCommand(loadCmd)
	loadCmd.Flags().StringVarP(&commitName, "commitHash", "c", "", "Specify the commit's hash to get to that point")
}
