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
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func day9(inputData []string) int {
	heightMap := make([][]int, len(inputData))

	for rowNum, heightData := range inputData {
		heights := strings.Split(heightData, "")

		if heightMap[rowNum] == nil {
			heightMap[rowNum] = make([]int, len(heights))
		}

		for colNum, height := range heights {
			h, _ := strconv.Atoi(height)

			heightMap[rowNum][colNum] = h
		}
	}

	var lowHeights []int

	for i, row := range heightMap {
		for j, v := range row {
			if i-1 < 0 { // top row
				if j-1 < 0 { // top left corner
					if v < heightMap[i][j+1] && v < heightMap[i+1][j] {
						lowHeights = append(lowHeights, v)
					}
				} else if j+1 == len(heightMap[i]) { // top right corner
					if v < heightMap[i][j-1] && v < heightMap[i+1][j] {
						lowHeights = append(lowHeights, v)
					}
				} else {
					if v < heightMap[i][j-1] && v < heightMap[i][j+1] && v < heightMap[i+1][j] {
						lowHeights = append(lowHeights, v)
					}
				}
			} else if i+1 == len(heightMap) { // bottom row
				if j-1 < 0 { // bottom left corner
					if v < heightMap[i][j+1] && v < heightMap[i-1][j] {
						lowHeights = append(lowHeights, v)
					}
				} else if j+1 == len(row) { // bottom right corner
					if v < heightMap[i][j-1] && v < heightMap[i-1][j] {
						lowHeights = append(lowHeights, v)
					}
				} else {
					if v < heightMap[i][j-1] && v < heightMap[i][j+1] && v < heightMap[i-1][j] {
						lowHeights = append(lowHeights, v)
					}
				}
			} else { //in middle
				if j-1 < 0 {
					if v < heightMap[i][j+1] && v < heightMap[i-1][j] && v < heightMap[i+1][j] {
						lowHeights = append(lowHeights, v)
					}
				} else if j+1 == len(row) {
					if v < heightMap[i][j-1] && v < heightMap[i-1][j] && v < heightMap[i+1][j] {
						lowHeights = append(lowHeights, v)
					}
				} else {
					if v < heightMap[i-1][j] && v < heightMap[i+1][j] && v < heightMap[i][j-1] && v < heightMap[i][j+1] {
						lowHeights = append(lowHeights, v)
					}
				}

			}
		}

	}

	riskLevel := 0

	for _, height := range lowHeights {
		riskLevel += height + 1
	}

	return riskLevel
}

var day9Cmd = &cobra.Command{
	Use: "day9",
	Run: func(cmd *cobra.Command, args []string) {
		res := day9(inputData)

		fmt.Printf("Result: %v\n", res)
	},
}

func init() {
	rootCmd.AddCommand(day9Cmd)
}
