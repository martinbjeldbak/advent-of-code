package cmd

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type hand struct {
	cards    []string
	bid      int
	strength int // 5 of a kind: 7, 4 of a kind: 6, and so on
}

func (h *hand) String() string {
	return fmt.Sprintf("(%v %v: %v)", h.strength, h.cards, h.bid)
}

func day7(inputData []string) int {
	hcStrength := map[string]int{
		"A": 13,
		"K": 12,
		"Q": 11,
		"J": 10,
		"T": 9,
	}
	for i := 2; i < 10; i++ {
		hcStrength[fmt.Sprint(i)] = i - 1
	}

	// Parse input into struct
	hands := make([]*hand, 0, len(inputData))
	for _, row := range inputData {
		handAndBids := strings.Fields(row)

		hand := hand{}
		for _, card := range strings.Split(handAndBids[0], "") {
			hand.cards = append(hand.cards, card)
		}

		bid, _ := strconv.Atoi(handAndBids[1])
		hand.bid = bid

		counts := make(map[string]int)
		for _, card := range hand.cards {
			counts[card]++
		}

		// Find card counts
		maxDuplicates := -1
		maxCard := ""
		for card, count := range counts {
			if count > maxDuplicates {
				maxDuplicates = count
				maxCard = card
			}
		}

		switch maxDuplicates {
		case 5:
			hand.strength = 8
		case 4:
			hand.strength = 7
		case 3:
			delete(counts, maxCard)
			hand.strength = 4
			for _, count := range counts {
				if count == 2 {
					hand.strength = 5
				}
			}
		case 2:
			delete(counts, maxCard)
			hand.strength = 2
			for _, count := range counts {
				if count == 2 {
					hand.strength = 3
				}
			}
		case 1:
			hand.strength = 1
		}

		hands = append(hands, &hand)

	}

	slices.SortFunc(hands, func(a, b *hand) int {
		if a.strength != b.strength {
			return cmp.Compare(a.strength, b.strength)
		}

		for i := 0; i < 6; i++ {
			if a.cards[i] != b.cards[i] {
				return cmp.Compare(hcStrength[a.cards[i]], hcStrength[b.cards[i]])
			}
		}

		return -1
	})

	sum := 0
	for i, hand := range hands {
		sum += hand.bid * (i + 1)
	}

	return sum
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "day7",
		Run: func(_ *cobra.Command, _ []string) {
			res := day7(inputData)

			fmt.Printf("Result: %v\n", res)
		},
	},
	)
}
