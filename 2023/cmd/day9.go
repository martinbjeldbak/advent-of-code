package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func day9(inputData []string) int {
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
			if i == len(differences)-1 {
				differences[i] = append(differences[i], 0)
				continue
			}

			leftOfMe := differences[i][len(differences[i])-1]
			below := differences[i+1][len(differences[i+1])-1]

			placeholder := leftOfMe + below

			differences[i] = append(differences[i], placeholder)
		}

		sum += differences[0][len(differences[0])-1]
	}

	return sum
}

func allZero(i []int) bool {
	for _, v := range i {
		if v != 0 {
			return false
		}
	}
	return true
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "day9",
		Run: func(_ *cobra.Command, _ []string) {
			start := time.Now()

			res := day9(inputData)

			fmt.Printf("Result: %v ran in %v\n", res, time.Since(start))
		},
	},
	)
}
