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
	"reflect"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func sumUnmarked(board [][]int) int {
	sumUnmarked := 0

	for _, row := range board {
		for _, v := range row {
			if v > 0 {
				sumUnmarked += v
			}
		}

	}

	return sumUnmarked
}

func day4(draws []int, boards [][][]int) int {
	winningScore := 0

	for _, draw := range draws {
		for _, board := range boards {
			for _, row := range board {
				for colIndex, value := range row {
					if value == draw {
						row[colIndex] = -1
					}

					// see if current board is winning after this update
					var currentColumn []int
					for _, row := range board {
						currentColumn = append(currentColumn, row[colIndex])
					}

					if reflect.DeepEqual(row, []int{-1, -1, -1, -1, -1}) || reflect.DeepEqual(currentColumn, []int{-1, -1, -1, -1, -1}) {
						return draw * sumUnmarked(board)
					}
				}
			}
		}
	}

	fmt.Printf("\n\nAfter first draw, state board 1:")
	fmt.Println(boards[0])

	fmt.Printf("\n\nAfter first draw, state board 2:")
	fmt.Println(boards[1])

	fmt.Printf("\n\nAfter first draw, state board 3:")
	fmt.Println(boards[2])

	return winningScore
}

// day4Cmd represents the day4 command
var day4Cmd = &cobra.Command{
	Use: "day4",
	Run: func(cmd *cobra.Command, args []string) {

		// Convert draw input to array of ints
		drawInput := inputData[0]
		var draws []int

		for _, draw := range strings.Split(drawInput, ",") {
			d, _ := strconv.Atoi(draw)

			draws = append(draws, d)
		}

		var boards [][][]int
		currentBoard := boards[0]
		currentBoardIndex := 0
		rowIndex := 0

		for _, rowData := range inputData[1:] {
			if rowData == "" { // setup a fresh board
				currentBoard = boards[currentBoardIndex+1]
				currentBoardIndex += 1
				rowIndex = 0
				continue
			}

			var row []int
			for _, v := range strings.Split(rowData, " ") {
				cellValue, _ := strconv.Atoi(v)

				row = append(row, cellValue)
			}

			currentBoard[rowIndex] = row
			rowIndex++
		}

		res := day4(draws, boards)

		fmt.Printf("Winning score is: %v", res)
	},
}

func init() {
	rootCmd.AddCommand(day4Cmd)
}
