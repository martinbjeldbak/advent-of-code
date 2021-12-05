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

	"github.com/spf13/cobra"
)

func dayOne(measurements []int) int {
	numIncreases := 0

	for i, measurement := range measurements {
		if i != 0 && measurement > measurements[i-1] {
			numIncreases += 1
		}
	}

	return numIncreases
}

// day1Cmd represents the day1 command
var day1Cmd = &cobra.Command{
	Use: "day1",
	RunE: func(cmd *cobra.Command, args []string) error {
		var measurements []int

		for _, measurement := range inputData {
			m, err := strconv.Atoi(measurement)

			if err != nil {
				return err
			}

			measurements = append(measurements, m)
		}

		fmt.Println(dayOne(measurements))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(day1Cmd)
}
