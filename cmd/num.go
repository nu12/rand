package cmd

import (
	"fmt"

	rand "github.com/nu12/rand/internal"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(numCmd)
}

var numCmd = &cobra.Command{
	Use:   "num <length>",
	Short: "Generate a random number of a given length",
	Long:  `Generate a random number of a given length`,
	Run: func(cmd *cobra.Command, args []string) {
		i := getLength(args)
		fmt.Println(rand.GenerateNum(i))
	},
}
