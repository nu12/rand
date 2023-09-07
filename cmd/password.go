package cmd

import (
	"fmt"

	rand "github.com/nu12/rand/internal"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(passwordCmd)
}

var passwordCmd = &cobra.Command{
	Use:   "password <length>",
	Short: "Generate a random string of characters, numbers and special signs of a given length",
	Long:  `Generate a random string of characters, numbers and special signs of a given length`,
	Run: func(cmd *cobra.Command, args []string) {
		i := getLength(args)
		checkLowerUpper()
		fmt.Println(rand.GeneratePassword(i, lower, upper))
	},
}
