package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func day3part2(inputData []string) int {
	schematic := make([][]string, 0, len(inputData))

	for _, row := range inputData {
		entries := strings.Split(row, "")
		schematic = append(schematic, entries)
	}

	partNumbers := make([]int, 0)

	for i, row := range schematic {
		for j, v := range row {

			if v != "*" {
				continue
			}

			rowAbove := i - 1
			rowBelow := i + 1

			localPartNumbers := make([]int, 0)

			partNumber, _, err := parseNumber(schematic[i], j-1)
			if err == nil {
				localPartNumbers = append(localPartNumbers, partNumber)
			}

			partNumber, _, err = parseNumber(schematic[i], j+1)
			if err == nil {
				localPartNumbers = append(localPartNumbers, partNumber)
			}

			if rowAbove >= 0 {
				partNumber, endIndex, err := parseNumber(schematic[rowAbove], j-1)
				if err == nil {
					localPartNumbers = append(localPartNumbers, partNumber)
				}

				for offset := j - 1; offset <= j+1; offset++ {
					if endIndex < offset {
						partNumber, endIndex, err = parseNumber(schematic[rowAbove], offset)
						if err == nil {
							localPartNumbers = append(localPartNumbers, partNumber)
						}
					}
				}
			}

			if rowBelow <= len(schematic) {
				partNumber, endIndex, err := parseNumber(schematic[rowBelow], j-1)
				if err == nil {
					localPartNumbers = append(localPartNumbers, partNumber)
				}
				for offset := j - 1; offset <= j+1; offset++ {
					if endIndex < offset {
						partNumber, endIndex, err = parseNumber(schematic[rowBelow], offset)
						if err == nil {
							localPartNumbers = append(localPartNumbers, partNumber)
						}
					}
				}

			}

			if len(localPartNumbers) == 2 {
				partNumbers = append(partNumbers, localPartNumbers[0] * localPartNumbers[1])
			}
		}
	}

	sum := 0
	for _, v := range partNumbers {
		sum += v
	}

	return sum
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "day3part2",
		Run: func(_ *cobra.Command, _ []string) {
			res := day3part2(inputData)

			fmt.Printf("Result: %v\n", res)
		},
	},
	)
}
