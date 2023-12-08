package cmd

import (
	"fmt"
	"math"
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

	time, _ := strconv.ParseFloat(rawTime, 32)
	minDistance, _ := strconv.ParseFloat(rawDistance, 32)

	d1 := -time + math.Sqrt(time*time - 4 * minDistance) / 2
	d2 := -time - math.Sqrt(time*time - 4 * minDistance) / 2

	return int(d1-d2)
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
