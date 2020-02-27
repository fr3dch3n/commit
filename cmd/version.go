package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

const Version = "v0.0.1"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Commit",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Commit " + Version)
	},
}
