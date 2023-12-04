package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func day3(inputData []string) int {
	schematic := make([][]string, 0, len(inputData))

	for _, row := range inputData {
		entries := strings.Split(row, "")
		schematic = append(schematic, entries)
	}

	partNumbers := make([]int, 0)

	for i, row := range schematic {
		for j, v := range row {

			_, err := strconv.Atoi(v)
			if v != "." && err != nil {
				rowAbove := i - 1
				rowBelow := i + 1

				partNumber, _, err := parseNumber(schematic[i], j-1)
				if err == nil {
					partNumbers = append(partNumbers, partNumber)
				}

				partNumber, _, err = parseNumber(schematic[i], j+1)
				if err == nil {
					partNumbers = append(partNumbers, partNumber)
				}

				if rowAbove >= 0 {
					partNumber, endIndex, err := parseNumber(schematic[rowAbove], j-1)
					if err == nil {
						partNumbers = append(partNumbers, partNumber)
					}

					for offset := j - 1; offset <= j+1; offset++ {
						if endIndex < offset {
							partNumber, endIndex, err = parseNumber(schematic[rowAbove], offset)
							if err == nil {
								partNumbers = append(partNumbers, partNumber)
							}
						}
					}
				}

				if rowBelow <= len(schematic) {
					partNumber, endIndex, err := parseNumber(schematic[rowBelow], j-1)
					if err == nil {
						partNumbers = append(partNumbers, partNumber)
					}
					for offset := j - 1; offset <= j+1; offset++ {
						if endIndex < offset {
							partNumber, endIndex, err = parseNumber(schematic[rowBelow], offset)
							if err == nil {
								partNumbers = append(partNumbers, partNumber)
							}
						}
					}
				}
			}
		}
	}

	sum := 0
	for _, partNumber := range partNumbers {
		sum += partNumber
	}

	return sum
}

func parseNumber(row []string, startCol int) (partNumber int, endIndex int, err error) {
	// 1. Check if actually a number
	_, err = strconv.Atoi(row[startCol])
	if err != nil {
		return 0, -1, err
	}

	// 2. Find first digit
	startIndexOfNumber := startCol
	for startIndexOfNumber >= 0 {
		_, err := strconv.Atoi(row[startIndexOfNumber])
		if err != nil {
			break
		}
		startIndexOfNumber--
	}
	startIndexOfNumber += 1

	// 3. Parse from first digit location
	number := ""
	for j := startIndexOfNumber; j < len(row); j++ {
		digit := row[j]

		_, err := strconv.Atoi(digit)
		if err != nil {
			break
		}

		number += digit
	}
	partNumber, _ = strconv.Atoi(number)

	endIndex = startIndexOfNumber + len(number)

	return partNumber, endIndex, nil
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "day3",
		Run: func(_ *cobra.Command, _ []string) {
			res := day3(inputData)

			fmt.Printf("Result: %v\n", res)
		},
	},
	)
}
