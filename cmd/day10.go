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

type stack []string

func (s stack) Push(v string) []string {
	return append(s, v)
}

func (s stack) Pop() (stack, string) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) Peek() string {
	l := len(s)
	return s[l-1]
}

var oppositeSymbol = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
	">": "<",
}

var scores = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

func day10(inputData []string) (score int) {

	for _, lineData := range inputData {
		s := make(stack, 0)
		symbols := strings.Split(lineData, "")

		for _, symbol := range symbols {
			switch symbol {
			case "(", "[", "{", "<":
				s = s.Push(symbol)
			case ")", "]", "}", ">":
				if s.Peek() != oppositeSymbol[symbol] {
					score += scores[symbol]
					goto endofline
				} else {
					s, _ = s.Pop()
				}
			}
		}
	endofline:
	}

	return score
}

var day10Cmd = &cobra.Command{
	Use: "day10",
	Run: func(cmd *cobra.Command, args []string) {
		res := day10(inputData)

		fmt.Printf("Result %v\n", res)
	},
}

func init() {
	rootCmd.AddCommand(day10Cmd)
}
