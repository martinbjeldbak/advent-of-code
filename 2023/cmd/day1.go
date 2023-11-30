package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func day1(inputData []string) int {
	res, err := strconv.Atoi(inputData[0])

	if err != nil {
		return -1
	}

	return res
}

var cmd = &cobra.Command{
	Use: "day1",
	Run: func(_ *cobra.Command, _ []string) {
		res := day1(inputData)


		fmt.Printf("Result: %v\n", res)
	},
}

func init() {
	rootCmd.AddCommand(cmd)
}
