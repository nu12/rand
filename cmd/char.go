package cmd

import (
	"fmt"

	rand "github.com/nu12/rand/pkg"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(charCmd)
}

var charCmd = &cobra.Command{
	Use:   "char <length>",
	Short: "Generate a random string of a given length",
	Long:  `Generate a random string of a given length`,
	Run: func(cmd *cobra.Command, args []string) {
		i := getLength(args)
		checkLowerUpper()
		fmt.Println(rand.Char(i, lower, upper))
	},
}
