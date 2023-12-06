package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func day6part2(inputData []string) int {
	rawTimes := strings.Fields(inputData[0])[1:]
	rawDistances := strings.Fields(inputData[1])[1:]

	rawTime := ""
	for _, v := range rawTimes {
		rawTime += v
	}

	rawDistance := ""
	for _, v := range rawDistances {
		rawDistance += v
	}

	time, _ := strconv.Atoi(rawTime)
	minDistance, _ := strconv.Atoi(rawDistance)

	numWay := 0

	for msHeld := 0; msHeld < time; msHeld++ {
		timeLeft := time - msHeld
		distance := timeLeft * msHeld

		if distance > minDistance {
			numWay++
		}
	}

	return numWay
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "day6part2",
		Run: func(_ *cobra.Command, _ []string) {
			res := day6part2(inputData)

			fmt.Printf("Result: %v\n", res)
		},
	},
	)
}
