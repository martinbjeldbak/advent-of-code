package cmd

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

// Src: https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// greatest common divisor (GCD) via Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func leastCommonMultiple(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = leastCommonMultiple(result, integers[i])
	}

	return result
}

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

	lengths := make([]int, len(curNodes))

	steps := 0
	atEnd := false
	for !atEnd {
		atEnd = true
		for i, curNode := range curNodes {
			targets := network[curNode]

			curNode = targets[instructionIndex[steps%len(instructionIndex)]]
			curNodes[i] = curNode
		}

		steps++

		for i, curNode := range curNodes {
			if curNode[2] == 'Z' {
				lengths[i] = steps
			}
		}

		for _, l := range lengths {
			if l == 0 {
				atEnd = false
			}
		}
	}

	return leastCommonMultiple(lengths[0], lengths[1], lengths[1:]...)
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
