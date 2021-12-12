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
	"strings"

	"github.com/spf13/cobra"
)

func day8(inputData []string) int {
	numSimpleDigits := 0

	for _, entryData := range inputData {
		digitsAndDecode := strings.Split(entryData, "|")

		digitData := strings.Trim(digitsAndDecode[1], "")

		for _, digit := range strings.Split(digitData, " ") {
			switch len(digit) {
			case 2, 3, 4, 7:
				numSimpleDigits++
			}
		}
	}

	return numSimpleDigits
}

var day8Cmd = &cobra.Command{
	Use: "day8",
	Run: func(cmd *cobra.Command, args []string) {
		res := day8(inputData)

		fmt.Printf("Result: %v\n", res)
	},
}

func init() {
	rootCmd.AddCommand(day8Cmd)
}
