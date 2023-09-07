package cmd

import (
	"fmt"

	rand "github.com/nu12/rand/internal"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(alphaCmd)
}

var alphaCmd = &cobra.Command{
	Use:   "alpha <length>",
	Short: "Generate a random string of characters and numbers of a given length",
	Long:  `Generate a random string of characters and numbers of a given length`,
	Run: func(cmd *cobra.Command, args []string) {
		i := getLength(args)
		checkLowerUpper()
		fmt.Println(rand.GenerateAlpha(i, lower, upper))
	},
}
