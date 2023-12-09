package cmd

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

func day8(inputData []string) int {
	instructions := strings.Split(inputData[0], "")
	instructionIndex := make([]int, len(instructions))
	for i, instruction := range instructions {
		if instruction == "L" {
			instructionIndex[i] = 0
		} else {
			instructionIndex[i] = 1
		}
	}


	network := make(map[string][]string)

	re := regexp.MustCompile(`\((\w+), (\w+)\)`)

	rest := inputData[2:]

	for _, row := range rest {
		partial := strings.Split(row, "=")
		node := strings.TrimSpace(partial[0])

		LR := re.FindStringSubmatch(partial[1])

		network[node] = []string{LR[1], LR[2]}
	}

	curNode := "AAA"
	targetNode := "ZZZ"

	steps := 0
	for curNode != targetNode {
		targets := network[curNode]

		curNode = targets[instructionIndex[steps % len(instructionIndex)]]
		steps++
	}


	return steps
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "day8",
		Run: func(_ *cobra.Command, _ []string) {
			res := day8(inputData)

			fmt.Printf("Result: %v\n", res)
		},
	},
	)
}
