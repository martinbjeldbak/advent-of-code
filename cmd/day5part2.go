package cmd

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func day5part2(linesData []string) int {
	coordinates := make(map[int]map[int]int)
	overlappingPoints := 0

	for _, lineData := range linesData {
		pointsData := strings.Split(lineData, "->")

		begin := strings.Split(strings.Trim(pointsData[0], " "), ",")
		end := strings.Split(strings.Trim(pointsData[1], " "), ",")

		x1, _ := strconv.Atoi(begin[0])
		y1, _ := strconv.Atoi(begin[1])

		x2, _ := strconv.Atoi(end[0])
		y2, _ := strconv.Atoi(end[1])

		fmt.Printf("%v, %v -> ", x1, y1)
		fmt.Printf("%v, %v\n", x2, y2)

		if y1 == y2 { // Draw horizontal line
			for i := int(math.Min(float64(x1), float64(x2))); i <= int(math.Max(float64(x1), float64(x2))); i++ {
				_, ok := coordinates[i]

				if !ok {
					coordinates[i] = make(map[int]int)
				}

				coordinates[i][y1] += 1
			}
		} else if x1 == x2 { // Draw vertical line
			for i := int(math.Min(float64(y1), float64(y2))); i <= int(math.Max(float64(y1), float64(y2))); i++ {
				_, ok := coordinates[x1]

				if !ok {
					coordinates[x1] = make(map[int]int)
				}

				coordinates[x1][i] += 1
			}
		} else { // diagonal line
			xDir := 1
			if x2 < x1 {
				xDir = -1
			}

			yDir := 1
			if y2 < y1 {
				yDir = -1
			}

			for i := 0; i <= int(math.Abs(float64(x2-x1))); i++ {
				_, ok := coordinates[x1+i*xDir]

				if !ok {
					coordinates[x1+i*xDir] = make(map[int]int)
				}

				coordinates[x1+i*xDir][y1+i*yDir] += 1
			}
		}
	}

	for _, column := range coordinates {
		for _, numPoints := range column {
			if numPoints > 1 {
				overlappingPoints++
			}
		}
	}

	return overlappingPoints
}

// day5part2Cmd represents the day5part2 command
var day5part2Cmd = &cobra.Command{
	Use: "day5part2",
	Run: func(cmd *cobra.Command, args []string) {

		res := day5part2(inputData)

		fmt.Printf("Number of overlapping lines is: %v\n", res)
	},
}

func init() {
	rootCmd.AddCommand(day5part2Cmd)
}
