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

	"github.com/spf13/cobra"
)

func day4part2(draws []int, boards [][][]int) int {
	boardsWon := make(map[int]bool)
	var winningScores []int

draw:
	for _, draw := range draws {
		for boardIndex, board := range boards {
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
						winningScores = append(winningScores, draw*sumUnmarked(board))
						boardsWon[boardIndex] = true

						if len(boardsWon) == len(boards) {
							break draw
						}
					}
				}
			}
		}
	}

	return winningScores[len(winningScores)-1]
}

// day4part2Cmd represents the day4part2 command
var day4part2Cmd = &cobra.Command{
	Use: "day4part2",
	Run: func(cmd *cobra.Command, args []string) {
		draws, boards := day4proccessRawInput(inputData)

		fmt.Printf("Last to win:")

		fmt.Println(day4part2(draws, boards))
	},
}

func init() {
	rootCmd.AddCommand(day4part2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day4part2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day4part2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
