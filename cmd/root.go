/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rand",
	Short: "A cli utility that generates random strings based on given paramenters",
	Long:  `A cli utility that generates random strings based on given paramenters`,
	// Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var lower bool
var upper bool

func init() {
	charCmd.Flags().BoolVarP(&lower, "lower", "l", false, "Only lower case characters")
	charCmd.Flags().BoolVarP(&upper, "upper", "u", false, "Only upper case characters")
	alphaCmd.Flags().BoolVarP(&lower, "lower", "l", false, "Only lower case characters")
	alphaCmd.Flags().BoolVarP(&upper, "upper", "u", false, "Only upper case characters")
	passwordCmd.Flags().BoolVarP(&lower, "lower", "l", false, "Only lower case characters")
	passwordCmd.Flags().BoolVarP(&upper, "upper", "u", false, "Only upper case characters")
}

func getLength(args []string) int {
	if len(args) < 1 {
		fmt.Println("Provide the length of the desired string")
		os.Exit(1)
	}

	i, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Argument should be a positive number")
		os.Exit(1)
	}
	return i
}

func checkLowerUpper() {
	if lower && upper {
		fmt.Println("May provide lower or upper, but not both")
		os.Exit(1)
	}
}
