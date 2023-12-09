package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func day9(inputData []string) int {
	return -1
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "day9",
		Run: func(_ *cobra.Command, _ []string) {
			res := day9(inputData)

			fmt.Printf("Result: %v\n", res)
		},
	},
	)
}
