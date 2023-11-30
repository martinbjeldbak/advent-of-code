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

func prettyPrint2D(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%3d", matrix[i][j])
		}
		fmt.Printf("\n")
	}
}

func needFlash(energyLevels [][]int) (needsFlash bool) {
	n := len(energyLevels)
	m := len(energyLevels[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if energyLevels[i][j] > 9 {
				return true
			}
		}
	}

	return false
}

func resetFlashedEnergy(energyLevels [][]int) [][]int {
	n := len(energyLevels)
	m := len(energyLevels[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if energyLevels[i][j] < 0 {
				energyLevels[i][j] = 0
			}
		}
	}
	return energyLevels
}

func flash(energyLevels [][]int) ([][]int, int) {
	flashes := 0
	n := len(energyLevels)
	m := len(energyLevels[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if energyLevels[i][j] > 9 {
				if i-1 >= 0 {
					energyLevels[i-1][j]++ // top

					if j-1 >= 0 {
						energyLevels[i-1][j-1]++ // top left
					}
					if j+1 < m {
						energyLevels[i-1][j+1]++ // top right
					}
				}

				if j-1 >= 0 {
					energyLevels[i][j-1]++ // left
				}

				if j+1 < m {
					energyLevels[i][j+1]++ // right
				}

				if i+1 < n {
					energyLevels[i+1][j]++ // below
					if j-1 >= 0 {
						energyLevels[i+1][j-1]++ // below left
					}
					if j+1 < m {
						energyLevels[i+1][j+1]++ // below right
					}
				}

				energyLevels[i][j] = -10 // mark as flashed
				flashes++
			}
		}
	}

	return energyLevels, flashes
}

func day11(inputData []string, numSteps int) (numFlashes int) {
	n := len(inputData)
	m := len(strings.Split(inputData[0], ""))

	energyLevels := make([][]int, n)
	for i := range energyLevels {
		energyLevels[i] = make([]int, 0, m)
	}

	// Convert input data to 10x10 matrix
	for i, row := range inputData {
		rowEnergyData := strings.Split(row, "")
		for _, v := range rowEnergyData {
			l, _ := strconv.Atoi(v)
			energyLevels[i] = append(energyLevels[i], l)
		}
	}

	fmt.Println("Before any steps:")
	prettyPrint2D(energyLevels)

	for step := 0; step < numSteps; step++ {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				energyLevels[i][j]++
			}
		}

		energyLevels, flashes := flash(energyLevels)
		numFlashes += flashes

		for needFlash(energyLevels) {
			energyLevels, flashes = flash(energyLevels)
			numFlashes += flashes
		}

		energyLevels = resetFlashedEnergy(energyLevels)

		fmt.Printf("After step %v\n\n", step+1)

		prettyPrint2D(energyLevels)
	}

	return numFlashes
}

var day11Cmd = &cobra.Command{
	Use: "day11",
	Run: func(cmd *cobra.Command, args []string) {
		res := day11(inputData, 100)

		fmt.Printf("Res: %v\n", res)
	},
}

func init() {
	rootCmd.AddCommand(day11Cmd)
}
