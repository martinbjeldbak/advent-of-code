package cmd

import (
	"fmt"
	"slices"
	"strings"

	"github.com/spf13/cobra"
)

func day04(inputData []string) (int, error) {
	wordSearch := make([][]string, len(inputData))

	for i, row := range inputData {
		wordSearch[i] = strings.Split(row, "")
	}

	candidates := make([]string, 0)

	for rowIdx, row := range wordSearch {
		for colIdx, letter := range row {
			if letter == "X" {
				letters := make([]string, 4)

				// written forwards
				if colIdx+3 < len(row) {
					letters = row[colIdx : colIdx+4]
					candidates = maybeAdd(candidates, letters)
				}

				// written backwards
				if colIdx-3 >= 0 {
					letters = row[colIdx-3 : colIdx+1]
					reversed := make([]string, 4)
					for i, l := range letters {
						reversed[i] = l
					}
					slices.Reverse(reversed)
					candidates = maybeAdd(candidates, reversed)
				}

				// up
				if rowIdx-3 >= 0 {
					letters = []string{
						letter,
						wordSearch[rowIdx-1][colIdx],
						wordSearch[rowIdx-2][colIdx],
						wordSearch[rowIdx-3][colIdx],
					}
					candidates = maybeAdd(candidates, letters)
				}

				// down
				if rowIdx+3 < len(wordSearch) {
					letters = []string{
						letter,
						wordSearch[rowIdx+1][colIdx],
						wordSearch[rowIdx+2][colIdx],
						wordSearch[rowIdx+3][colIdx],
					}
					candidates = maybeAdd(candidates, letters)
				}

				// upperleft
				if colIdx-3 >= 0 && rowIdx-3 >= 0 {
					letters = []string{
						letter,
						wordSearch[rowIdx-1][colIdx-1],
						wordSearch[rowIdx-2][colIdx-2],
						wordSearch[rowIdx-3][colIdx-3],
					}
					candidates = maybeAdd(candidates, letters)
				}

				// upper right
				if colIdx+3 < len(row) && rowIdx-3 >= 0 {
					letters = []string{
						letter,
						wordSearch[rowIdx-1][colIdx+1],
						wordSearch[rowIdx-2][colIdx+2],
						wordSearch[rowIdx-3][colIdx+3],
					}
					candidates = maybeAdd(candidates, letters)
				}

				// bottom left
				if colIdx-3 >= 0 && rowIdx+3 < len(wordSearch) {
					letters = []string{
						letter,
						wordSearch[rowIdx+1][colIdx-1],
						wordSearch[rowIdx+2][colIdx-2],
						wordSearch[rowIdx+3][colIdx-3],
					}
					candidates = maybeAdd(candidates, letters)
				}

				// bottom right
				if colIdx+3 < len(row) && rowIdx+3 < len(wordSearch) {
					letters = []string{
						letter,
						wordSearch[rowIdx+1][colIdx+1],
						wordSearch[rowIdx+2][colIdx+2],
						wordSearch[rowIdx+3][colIdx+3],
					}
					candidates = maybeAdd(candidates, letters)
				}
			}
		}
	}


	return len(candidates), nil
}

func maybeAdd(candidates, letters []string) []string {
	if strings.Join(letters, "") == "XMAS" {
		return append(candidates, strings.Join(letters, ""))
	}

	return candidates
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "day04",
		RunE: func(_ *cobra.Command, _ []string) error {
			res, err := day04(inputData)
			if err != nil {
				return err
			}

			fmt.Printf("Result: %v\n", res)

			return nil
		},
	},
	)
}
