package cmd

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func day4part2(inputData []string) int {
	cardWinningNumberCount := make(map[int]int)
	re := regexp.MustCompile(`\d+`)

	for _, row := range inputData {
		points := 0
		cardAndNumbers := strings.Split(row, ":")
		numbers := strings.Split(cardAndNumbers[1], "|")

		match := re.FindString(cardAndNumbers[0])

		cardNumber, _ := strconv.Atoi(match)

		winningNumbers := strings.Fields(numbers[0])
		myNumbers := strings.Fields(numbers[1])

		winningNumber := make(map[string]bool)
		for _, w := range winningNumbers {
			winningNumber[w] = true
		}

		for _, v := range myNumbers {
			if winningNumber[v] {
				points += 1
			}
		}

		cardWinningNumberCount[cardNumber] = points
	}

	queue := make([]int, 0, len(inputData))

	// warm queue with original cards
	for j := 1; j <= len(inputData); j++ {
		queue = append(queue, j)
	}

	i := 0
	for i < len(queue) {
		currentCard := queue[i]

		wonCardCount := cardWinningNumberCount[currentCard]

		for j := 1; j <= wonCardCount; j++ {
			queue = append(queue, currentCard + j)
		}

		i++
	}

	return len(queue)
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "day4part2",
		Run: func(_ *cobra.Command, _ []string) {
			res := day4part2(inputData)

			fmt.Printf("Result: %v\n", res)
		},
	},
	)
}
