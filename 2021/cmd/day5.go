/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func overlappingLines(linesData []string) int {
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
		} else {
			fmt.Println("Got a non-horizontal or vertial line, not currently supported")
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

// day5Cmd represents the day5 command
var day5Cmd = &cobra.Command{
	Use: "day5",
	Run: func(cmd *cobra.Command, args []string) {
		res := overlappingLines(inputData)

		fmt.Printf("Number of overlapping lines is: %v\n", res)
	},
}

func init() {
	rootCmd.AddCommand(day5Cmd)
}
