package cmd

import (
	"github.com/darshan744/Trace/internals"
	"github.com/spf13/cobra"
)

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Lists out the commits done by the user",
	Long:  `Logs the commit done by the user which he can use to switch back to a certain snapshot`,
	Run: func(cmd *cobra.Command, args []string) {
		internals.LogAllCommit()
	},
}

func init() {
	rootCmd.AddCommand(logCmd)
}
