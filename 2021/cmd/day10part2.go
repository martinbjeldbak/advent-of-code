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

func day10part2(inputData []string) (middleScore int) {
	var symbolScore = map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}
	var scores = make([]int, 0, len(inputData))

	for _, lineData := range inputData {
		s := make(stack, 0)
		corrupted := false
		symbols := strings.Split(lineData, "")

		for _, symbol := range symbols {
			if !corrupted {
				switch symbol {
				case "(", "[", "{", "<":
					s = s.Push(symbol)
				case ")", "]", "}", ">":
					if s.Peek() != oppositeSymbol[symbol] {
						corrupted = true
					} else {
						s, _ = s.Pop()
					}
				}
			}
		}

		if !corrupted { // Line not corrupted, but incomplete
			lineScore := 0

			r, symbol := s.Pop()

			for symbol != "" {
				lineScore = lineScore*5 + symbolScore[oppositeSymbol[symbol]]

				r, symbol = r.Pop()
			}

			scores = append(scores, lineScore)
		}
	}

	sort.Ints(scores)
	l := len(scores)

	return scores[l/2]
}

// day10part2Cmd represents the day10part2 command
var day10part2Cmd = &cobra.Command{
	Use: "day10part2",
	Run: func(cmd *cobra.Command, args []string) {
		res := day10part2(inputData)

		fmt.Printf("Result %v\n", res)
	},
}

func init() {
	rootCmd.AddCommand(day10part2Cmd)
}
