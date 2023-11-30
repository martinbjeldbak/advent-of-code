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
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

type byLength []string

func (a byLength) Len() int           { return len(a) }
func (a byLength) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byLength) Less(i, j int) bool { return len(a[i]) < len(a[j]) }

func day8part2(inputData []string) int {
	numbers := []string{
		"abcefg",  // 0
		"cf",      // 1
		"aedeg",   // 2
		"acdfg",   // 3
		"bcdf",    // 4
		"abdfg",   // 5
		"abdefg",  // 6
		"acf",     // 7
		"abcdefg", // 8
		"abcdfg",  // 9
	}
	fmt.Println(numbers)

	for _, entryData := range inputData {
		digitsAndDecode := strings.Split(entryData, "|")
		// decodeData := strings.TrimSpace(digitsAndDecode[1])

		// Decode signals
		uniqSignalData := strings.TrimSpace(digitsAndDecode[0])
		fmt.Printf("Reviewing signal data %v\n", uniqSignalData)
		uniqSignals := strings.Split(uniqSignalData, " ")

		signalMapping := make(map[string]string)
		sort.Sort(byLength(uniqSignals))

		// set up mapping of initial, known digits (1,7,4,8)
		for _, digit := range uniqSignals {
			segments := strings.Split(digit, "")

			for _, segment := range segments {
				_, ok := signalMapping[segment]

				if !ok {
					switch len(digit) {
					case 2: // number 1
						signalMapping[segment] = "cf"
					case 3: // number 7
						signalMapping[segment] = "a"
					case 4: // number 4
						signalMapping[segment] = "bd"
					case 7: // number 8
						signalMapping[segment] = "eg"
					}
				}
			}
		}

		for _, digit := range uniqSignals {
			fmt.Printf("Reviewing signal %v\n", digit)
			segments := strings.Split(digit, "")

			// Skip known values
			if len(digit) == 2 || len(digit) == 3 || len(digit) == 4 || len(digit) == 7 {
				continue
			}

			mappings := make(map[string]string)
			for _, segment := range segments {
				//fmt.Printf("  looking at segment %v\n", segment)

				mappings[signalMapping[segment]] = segment

			}

			fmt.Printf("  %v", mappings)

		}
	}

	return 0
}

var day8part2Cmd = &cobra.Command{
	Use: "day8part2",
	Run: func(cmd *cobra.Command, args []string) {
		res := day8part2(inputData)

		fmt.Printf("Result: %v\n", res)
	},
}

func init() {
	rootCmd.AddCommand(day8part2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day8part2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day8part2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
