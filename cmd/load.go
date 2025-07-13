package cmd

import (
	"fmt"

	"github.com/darshan744/Trace/internals"
	"github.com/spf13/cobra"
)

var commitName string

// loadCmd represents the load command
var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Move to a specified commit",
	Long:  `This command helps the user to move out to the specified commit done by the user `,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Only one hash is allowed")
			return
		}
		internals.RestoreHistory(args[0])
	},
}

func init() {
	rootCmd.AddCommand(loadCmd)
}
