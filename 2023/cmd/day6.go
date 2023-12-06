package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func day6(inputData []string) int {
	rawTimes := strings.Fields(inputData[0])[1:]
	rawDistances := strings.Fields(inputData[1])[1:]

	times := make([]int, 0, len(rawTimes))
	minDistances := make([]int, 0, len(rawDistances))
	numWaysPerRace := make([]int, 0, len(rawDistances))

	for _, v := range rawTimes {
		t, _ := strconv.Atoi(v)
		times = append(times, t)
	}

	for _, v := range rawDistances {
		d, _ := strconv.Atoi(v)
		minDistances = append(minDistances, d)
	}


	for r := 0; r < len(times); r++ {
		time := times[r]
		minDistance := minDistances[r]

		numWay := 0

		for msHeld := 0; msHeld < time; msHeld++ {
			timeLeft := time - msHeld
			distance := timeLeft * msHeld

			if distance > minDistance {
				numWay++
			}
		}

		numWaysPerRace = append(numWaysPerRace, numWay)
	}

	product := 1
	for _, v := range numWaysPerRace {
		product = product * v
	}

	return product

}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "day6",
		Run: func(_ *cobra.Command, _ []string) {
			res := day6(inputData)

			fmt.Printf("Result: %v\n", res)
		},
	},
	)
}
