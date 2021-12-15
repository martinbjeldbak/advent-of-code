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

func day7part2(inputData []string) int {
	crabData := strings.Split(inputData[0], ",")
	var crabs []int

	for _, v := range crabData {
		crabPos, _ := strconv.Atoi(v)

		crabs = append(crabs, crabPos)
	}

	maxPos := math.MinInt
	for _, v := range crabs {
		if v > maxPos {
			maxPos = v
		}
	}

	minFuelCost := math.MaxInt

	currentFuelModifier := 0

	for targetPos := 1; targetPos <= maxPos; targetPos++ {
		//fmt.Printf("Evaluating cost of target pos of %v\n", targetPos)
		fuelCost := 0

		for _, v := range crabs {
			numSteps := int(math.Abs(float64(v - targetPos)))
			fuelCost += (numSteps*numSteps + numSteps) / 2 // solution to binomial coefficient
		}

		if fuelCost < minFuelCost {
			minFuelCost = fuelCost
			// bestPos = targetPos
		}

		// fmt.Printf(" cost is %v (%v)\n", fuelCost, currentFuelModifier)

		currentFuelModifier++
	}

	// fmt.Printf("Best position %v\n", bestPos)

	return minFuelCost
}

var day7part2Cmd = &cobra.Command{
	Use: "day7part2",
	Run: func(cmd *cobra.Command, args []string) {
		res := day7part2(inputData)

		fmt.Printf("Result: %v\n", res)
	},
}

func init() {
	rootCmd.AddCommand(day7part2Cmd)
}
