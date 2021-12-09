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

func narrow(reports []string, bitIndex int, mostCommon bool) string {
	if len(reports) == 1 {
		return reports[0]
	} else {
		var ones []string
		var zeroes []string

		for _, report := range reports {
			reportBits := strings.Split(report, "") // consider doing this logic in preproccessing

			if reportBits[bitIndex] == "1" {
				ones = append(ones, report)
			} else {
				zeroes = append(zeroes, report)
			}
		}

		if (mostCommon && len(ones) >= len(zeroes)) || (!mostCommon && len(ones) < len(zeroes)) {
			return narrow(ones, bitIndex+1, mostCommon)
		} else {
			return narrow(zeroes, bitIndex+1, mostCommon)
		}
	}
}

func day3part2(diagnosticReports []string) (int, error) {
	oxygenGeneratorReport := narrow(diagnosticReports, 0, true)

	oxygenGeneratorRating, err := strconv.ParseInt(oxygenGeneratorReport, 2, 64)

	if err != nil {
		return 0, err
	}

	co2scrubberReport := narrow(diagnosticReports, 0, false)
	co2scrubberRating, err := strconv.ParseInt(co2scrubberReport, 2, 64)

	if err != nil {
		return 0, err
	}

	return int(oxygenGeneratorRating) * int(co2scrubberRating), nil
}

// day3part2Cmd represents the day3part2 command
var day3part2Cmd = &cobra.Command{
	Use: "day3part2",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := day3part2(inputData)

		if err != nil {
			return err
		}

		fmt.Println("Life support rating:")
		fmt.Println(res)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(day3part2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day3part2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day3part2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
