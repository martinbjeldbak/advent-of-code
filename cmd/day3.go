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

func day3(diagnosticReports []string) (int, error) {
	var gammaRate []int

	convertedReports := make([][]int, len(diagnosticReports))

	for currentReport, stringReport := range diagnosticReports {
		splitReport := strings.Split(stringReport, "")

		row := make([]int, len(splitReport))
		for i, stringBit := range splitReport {
			bit, err := strconv.Atoi(stringBit)

			if err != nil {
				return 0, err
			}

			row[i] = bit
		}

		convertedReports[currentReport] = row
	}

	fmt.Println("Converted report:")
	fmt.Println(convertedReports)

	for bitIndex := 0; bitIndex < len(diagnosticReports[0]); bitIndex++ {
		fmt.Printf("Looking at bit %v\n", bitIndex)

		numOnes := 0

		for _, diagnosticReport := range convertedReports {
			if diagnosticReport[bitIndex] == 1 {
				numOnes += 1
			}
		}

		if numOnes > len(diagnosticReports)/2 {
			gammaRate = append(gammaRate, 1)
		} else {
			gammaRate = append(gammaRate, 0)
		}
	}

	strconv.ParseInt()

	fmt.Println(gammaRate)

	return 0, nil
}

// day3Cmd represents the day3 command
var day3Cmd = &cobra.Command{
	Use: "day3",
	RunE: func(cmd *cobra.Command, args []string) error {
		powerConsumption, err := day3(inputData)

		if err != nil {
			return err
		}

		fmt.Print(powerConsumption)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(day3Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day3Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day3Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
