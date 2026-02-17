package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show current version",
	Long:  `Show current version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v1.2.1")
	},
}
