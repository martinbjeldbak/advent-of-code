package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func day04part2(inputData []string) (int, error) {
	wordSearch := make([][]string, len(inputData))
	numTimes := 0

	for i, row := range inputData {
		wordSearch[i] = strings.Split(row, "")
	}

	for rowIdx, row := range wordSearch {
		for colIdx, letter := range row {
			if letter == "A" {
				if rowIdx-1 < 0 || rowIdx+1 == len(wordSearch) {
					continue
				}
				if colIdx-1 < 0 || colIdx+1 == len(row) {
					continue
				}

				candidates := []string{
					wordSearch[rowIdx-1][colIdx-1] + "A" + wordSearch[rowIdx+1][colIdx+1],
					wordSearch[rowIdx-1][colIdx+1] + "A" + wordSearch[rowIdx+1][colIdx-1],
					wordSearch[rowIdx+1][colIdx+1] + "A" + wordSearch[rowIdx-1][colIdx-1],
					wordSearch[rowIdx+1][colIdx-1] + "A" + wordSearch[rowIdx-1][colIdx+1],
				}


				found := 0
				for _, candidate := range candidates {
					if candidate == "MAS" {
						found++
					}
				}

				if found == 2 {
					numTimes++
				}
			}
		}
	}

	return numTimes, nil
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "day04part2",
		RunE: func(_ *cobra.Command, _ []string) error {
			res, err := day04part2(inputData)
			if err != nil {
				return err
			}

			fmt.Printf("Result: %v\n", res)

			return nil
		},
	},
	)
}
