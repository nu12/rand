package cmd

import (
	"fmt"

	rand "github.com/nu12/rand/pkg"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(uuidCmd)
}

var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Generate an UUID",
	Long:  `Generate an UUID`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(rand.UUID())
	},
}
