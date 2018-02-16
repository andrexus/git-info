package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Use:  "git-info",
	Long: "generate git information",
	Run: func(cmd *cobra.Command, args []string) {
		getGitInfo(cmd, args)
	},
}

// NewRoot will add flags and subcommands to the different commands
func RootCmd() *cobra.Command {
	rootCmd.AddCommand(&gitInfoCmd, &versionCmd)
	rootCmd.PersistentFlags().StringP("mode", "m", "full", "output mode (short, full)")
	return &rootCmd
}
