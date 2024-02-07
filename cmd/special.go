package cmd

import (
	"fmt"

	rand "github.com/nu12/rand/pkg"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(specialCmd)
}

var specialCmd = &cobra.Command{
	Use:   "special <length>",
	Short: "Generate a random string of special characters of a given length",
	Long:  `Generate a random string of special characters of a given length`,
	Run: func(cmd *cobra.Command, args []string) {
		i := getLength(args)
		fmt.Println(rand.Special(i))
	},
}
