package cmd

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func day1(inputData []string) (int, error) {
	listOne := make([]int, 0, len(inputData))
	listTwo := make([]int, 0, len(inputData))

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
		listOne = append(listOne, d1)
		listTwo = append(listTwo, d2)
	}

	sort.Ints(listOne)
	sort.Ints(listTwo)

	totalDistance := 0
	for i, l1 := range listOne {
		d := l1 - listTwo[i]
		if d < 0 {
			d = -d
		}

		totalDistance += d
	}

	return totalDistance, nil
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "day1",
		RunE: func(_ *cobra.Command, _ []string) error {
			res, err := day1(inputData)
			if err != nil {
				return err
			}

			fmt.Printf("Result: %v\n", res)

			return nil
		},
	},
		)
}
