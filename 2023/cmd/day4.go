package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func day4(inputData []string) int {
	totalPoints := 0

	for _, row := range inputData {
		points := 0
		gameAndCards := strings.Split(row, ":")
		numbers := strings.Split(gameAndCards[1], "|")

		winningNumbers := strings.Fields(numbers[0])
		myNumbers := strings.Fields(numbers[1])

		winningNumber := make(map[string]bool)
		for _, w := range winningNumbers {
			winningNumber[w] = true
		}

		for _, v := range myNumbers {
			if winningNumber[v] {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}

		totalPoints += points
	}

	return totalPoints
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "day4",
		Run: func(_ *cobra.Command, _ []string) {
			res := day4(inputData)

			fmt.Printf("Result: %v\n", res)
		},
	},
	)
}
