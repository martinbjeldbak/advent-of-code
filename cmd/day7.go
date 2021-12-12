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

func day7(inputData []string) int {
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
	bestPos := 0

	for targetPos := 1; targetPos <= maxPos; targetPos++ {
		fmt.Printf("Trying target pos: %v\n", targetPos)
		fuelCost := 0

		for _, v := range crabs {
			fuelCost += int(math.Abs(float64(v - targetPos)))
		}
		fmt.Printf("  got fuel cost %v\n", fuelCost)

		if fuelCost < minFuelCost {
			minFuelCost = fuelCost
			bestPos = targetPos
		}
	}

	fmt.Printf("Best position: %v\n", bestPos)

	return minFuelCost
}

// day7Cmd represents the day7 command
var day7Cmd = &cobra.Command{
	Use: "day7",
	Run: func(cmd *cobra.Command, args []string) {
		res := day7(inputData)

		fmt.Printf("Result: %v\n", res)
	},
}

func init() {
	rootCmd.AddCommand(day7Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day7Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day7Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
