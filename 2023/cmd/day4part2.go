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

	test := calcScratchards(1, cardWinningNumberCount)
	fmt.Printf("Lol: %v\n", test)

	return -1
}

func calcScratchards(curCard int, cards map[int]int) int {
	matchingNumbers := cards[curCard]

	copies := matchingNumbers

	v := make([]int, 0, matchingNumbers)
	for i := 1; i <= matchingNumbers; i++ {
		copies += calcScratchards(curCard+i, cards)
		v = append(v, calcScratchards(curCard+i, cards))
	}

	fmt.Printf("Card %v wins copies %v\n", curCard, copies)
	fmt.Println(v)

	return copies
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
