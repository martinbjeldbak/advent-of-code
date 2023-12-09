package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func day9part2(inputData []string) int {
	histories := make([][]int, 0, len(inputData))

	for _, row := range inputData {
		readingsRaw := strings.Fields(row)
		readings := make([]int, len(readingsRaw))

		for i, r := range readingsRaw {
			reading, _ := strconv.Atoi(r)

			readings[i] = reading
		}

		histories = append(histories, readings)
	}

	sum := 0

	for _, history := range histories {
		differences := make([][]int, 0)

		differences = append(differences, history)
		curDifference := history

		for !allZero(curDifference) {
			diff := make([]int, len(curDifference)-1)
			for i := 0; i < len(curDifference)-1; i++ {
				diff[i] = curDifference[i+1] - curDifference[i]
			}

			curDifference = diff
			differences = append(differences, curDifference)
		}

		for i := len(differences) - 1; i >= 0; i-- {
			newDiff := make([]int, 1, len(differences)+1)

			if i == len(differences)-1 {
				differences[i] = append(newDiff, differences[i]...)
				continue
			}

			rightOfMe := differences[i][0]
			below := differences[i+1][0]

			placeholder := rightOfMe - below

			newDiff[0] = placeholder

			differences[i] = append(newDiff, differences[i]...)
		}
		fmt.Println(differences)

		sum += differences[0][0]
	}

	return sum
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "day9part2",
		Run: func(_ *cobra.Command, _ []string) {
			start := time.Now()

			res := day9part2(inputData)

			fmt.Printf("Result: %v ran in %v\n", res, time.Since(start))
		},
	},
	)
}
