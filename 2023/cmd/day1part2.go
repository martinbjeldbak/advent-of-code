package cmd

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/spf13/cobra"
)

func day1part2(inputData []string) (int, error) {
	digitLetterToDigit := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	calibrationValues := make([]int, len(inputData))

	for _, row := range inputData {
		runes := make([]rune, 0)
		for _, c := range row {
			runes = append(runes, c)
		}

		firstIdx := len(row)
		firstValue := ""
		lastIdx := -1
		lastValue := ""

		for idx, c := range runes {
			if unicode.IsDigit(c) {
				if idx < firstIdx {
					firstIdx = idx
					firstValue = string(runes[firstIdx])
				}
				if idx > lastIdx {
					lastIdx = idx
					lastValue = string(runes[lastIdx])
				}
			} else {
				chars := ""
				for i := idx; i < len(runes); i++ {
					chars += string(runes[i])

					num, ok := digitLetterToDigit[chars]

					if ok {
						if idx < firstIdx {
							firstIdx = idx
							firstValue = num
						}
						if idx > lastIdx {
							lastIdx = idx
							lastValue = num
						}
					}
				}
			}
		}

		calibrationValue, err := strconv.Atoi(firstValue + lastValue)
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

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "day1part2",
		RunE: func(_ *cobra.Command, _ []string) error {
			res, err := day1part2(inputData)
			if err != nil {
				return err
			}

			fmt.Printf("Result: %v\n", res)

			return nil
		},
	},
	)
}
