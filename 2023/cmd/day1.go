package cmd

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/spf13/cobra"
)

func day1(inputData []string) (int, error) {
	calibrationValues := make([]int, len(inputData))

	for _, row := range inputData {
		runes := make([]rune, 0)
		for _, c := range row {
			runes = append(runes, c)
		}

		firstIdx := len(row)
		lastIdx := -1

		for idx, c := range row {
			if unicode.IsDigit(c) {
				if idx < firstIdx {
					firstIdx = idx
				}
				if idx > lastIdx {
					lastIdx = idx
				}
			}

		}

		value := string(runes[firstIdx]) + string(runes[lastIdx])

		calibrationValue, err := strconv.Atoi(value)
		if err != nil {
			return -1, err
		}

		calibrationValues = append(calibrationValues, int(calibrationValue))
	}

	sum := 0
	for _, v := range calibrationValues {
		sum += v
	}

	return sum, nil
}

var cmd = &cobra.Command{
	Use: "day1",
	RunE: func(_ *cobra.Command, _ []string) error {
		res, err := day1(inputData)
		if err != nil {
			return err
		}

		fmt.Printf("Result: %v\n", res)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(cmd)
}
