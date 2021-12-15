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

func day6(inputData []string, numDays int) int {
	fmt.Printf("Initial state: %v\n", inputData[0])
	splitData := strings.Split(inputData[0], ",")

	// Initial setup
	var currentState []int
	for _, v := range splitData {
		fishAge, _ := strconv.Atoi(v)

		currentState = append(currentState, fishAge)
	}

	for i := 0; i < numDays; i++ {
		var nextState []int

		for _, age := range currentState {
			age--

			switch age {
			case -1:
				nextState = append(nextState, 6) // reset internal timer
				nextState = append(nextState, 8) // spawn new fish
			default:
				nextState = append(nextState, age)
			}
		}

		currentState = nextState
		// fmt.Printf("After %v days: %v\n", i, nextState)
	}

	return len(currentState)
}

var numDays int

// day6Cmd represents the day6 command
var day6Cmd = &cobra.Command{
	Use: "day6",
	Run: func(cmd *cobra.Command, args []string) {
		res := day6(inputData, numDays)

		fmt.Printf("Result: %v\n", res)
	},
}

func init() {
	rootCmd.AddCommand(day6Cmd)

	day6Cmd.Flags().IntVarP(&numDays, "numDays", "d", 0, "Number of days to simulate")
	err := day6Cmd.MarkFlagRequired("numDays")

	if err != nil {
		fmt.Printf("Got err %v\n", err)
	}
}
