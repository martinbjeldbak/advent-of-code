package cmd

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func isSafe(report []int) bool {
	reportGood := false
	if slices.IsSortedFunc(report, func(a, b int) int {
		if a > b {
			return 1
		} else {
			return -1
		}
	}) || slices.IsSortedFunc(report, func(a, b int) int {
		if a < b {
			return 1
		} else {
			return -1
		}
	}) {
		for j, level := range report {
			if j == 0 {
				continue
			}
			prevLevel := report[j-1]
			difference := prevLevel - level
			if difference < 0 {
				difference = -difference
			}

			if difference > 3 {
				reportGood = false
				break
			}
			reportGood = true
		}
	}

	return reportGood
}

func day2part2(inputData []string) (int, error) {
	reports := make([][]int, len(inputData))

	for i, unparsedReport := range inputData {
		unparsedLevels := strings.Split(unparsedReport, " ")
		reports[i] = make([]int, len(unparsedLevels))

		for j, unparsedLevel := range unparsedLevels {
			level, err := strconv.Atoi(unparsedLevel)
			if err != nil {
				return 0, err
			}
			reports[i][j] = level
		}
	}

	numSafeReports := 0
	for _, report := range reports {
		if isSafe(report) {
			numSafeReports++
		} else {
			for i := 0; i < len(report); i++ {
				updatedReport := make([]int, len(report))
				for l, x := range report {
					updatedReport[l] = x
				}

				updatedReport = slices.Delete(updatedReport, i, i+1)
				if isSafe(updatedReport) {
					numSafeReports++
					break
				}
			}
		}
	}

	return numSafeReports, nil
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "day02part2",
		RunE: func(_ *cobra.Command, _ []string) error {
			res, err := day2part2(inputData)
			if err != nil {
				return err
			}

			fmt.Printf("Result: %v\n", res)

			return nil
		},
	},
	)
}
