package cmd

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

func day8part2(inputData []string) int {
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

	curNodes := make([]string, 0)

	for _, row := range rest {
		partial := strings.Split(row, "=")
		node := strings.TrimSpace(partial[0])

		if node[2] == 'A' {
			curNodes = append(curNodes, node)
		}

		LR := re.FindStringSubmatch(partial[1])

		network[node] = []string{LR[1], LR[2]}
	}

	fmt.Println("Starting at")
	fmt.Println(curNodes)

	steps := 0
	atEnd := false
	for !atEnd {
		// fmt.Println(curNodes)

		atEnd = true
		for i, curNode := range curNodes {
			targets := network[curNode]

			curNode = targets[instructionIndex[steps%len(instructionIndex)]]
			curNodes[i] = curNode

			if curNode[2] != 'Z' {
				atEnd = false
			}
		}

		steps++

		if atEnd == true {
			break
		}
	}

	return steps
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "day8part2",
		Run: func(_ *cobra.Command, _ []string) {
			res := day8part2(inputData)

			fmt.Printf("Result: %v\n", res)
		},
	},
	)
}
