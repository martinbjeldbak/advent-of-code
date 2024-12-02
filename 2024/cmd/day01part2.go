package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func day01part2(inputData []string) (int, error) {
	leftList := make([]int, 0, len(inputData))
	rightList := make([]int, 0, len(inputData))

	for _, pairUnparsed := range inputData {
		pair := strings.Split(pairUnparsed, "   ")

		d1, err := strconv.Atoi(pair[0])
		if err != nil {
			return 0, err
		}
		d2, err := strconv.Atoi(pair[1])
		if err != nil {
			return 0, err
		}
		leftList = append(leftList, d1)
		rightList = append(rightList, d2)
	}

	similarityScore := 0

	for _, l := range leftList {
		appearances := 0
		for _, r := range rightList {
			if r == l {
				appearances++
			}
		}
		similarityScore += l * appearances
	}

	return similarityScore, nil
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "day01part2",
		RunE: func(_ *cobra.Command, _ []string) error {
			res, err := day01part2(inputData)
			if err != nil {
				return err
			}

			fmt.Printf("Result: %v\n", res)

			return nil
		},
	},
		)
}
