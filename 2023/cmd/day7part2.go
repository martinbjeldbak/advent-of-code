package cmd

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func calcStrength(cards []string) (strength int) {
	counts := make(map[string]int)
	for _, card := range cards {
		counts[card]++
	}

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
		strength = 8
	case 4:
		strength = 7
	case 3:
		delete(counts, maxCard)
		strength = 4
		for _, count := range counts {
			if count == 2 {
				strength = 5
			}
		}
	case 2:
		delete(counts, maxCard)
		strength = 2
		for _, count := range counts {
			if count == 2 {
				strength = 3
			}
		}
	case 1:
		strength = 1
	}

	return strength
}

func maxStrength(cards []string, curIndex int, strengthToCard map[int]string) int {
	max := calcStrength(cards)

	for i := curIndex; i < len(cards); i++ {
		if cards[i] == "J" {
			for j := 0; j < 14; j++ {
				cardsModified := make([]string, len(cards))
				copy(cardsModified, cards)
				cardsModified[i] = strengthToCard[j]
				strength := maxStrength(cardsModified, curIndex+1, strengthToCard)
				if strength > max {
					max = strength
				}
			}
		}
	}

	return max
}

func day7part2(inputData []string) int {
	hcStrength := map[string]int{
		"A": 13,
		"K": 12,
		"Q": 11,
		"J": 0,
		"T": 9,
	}
	for i := 2; i < 10; i++ {
		hcStrength[fmt.Sprint(i)] = i - 1
	}

	reverseMap := make(map[int]string)
	for card, strength := range hcStrength {
		reverseMap[strength] = card
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
		hand.strength = maxStrength(hand.cards, 0, reverseMap)

		hands = append(hands, &hand)
	}

	slices.SortFunc(hands, func(a, b *hand) int {
		if a.strength != b.strength {
			return cmp.Compare(a.strength, b.strength)
		}

		for i := 0; i < 5; i++ {
			if a.cards[i] != b.cards[i] {
				return cmp.Compare(hcStrength[a.cards[i]], hcStrength[b.cards[i]])
			}
		}

		fmt.Println("Shouldn't be here")
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
		Use: "day7part2",
		Run: func(_ *cobra.Command, _ []string) {
			res := day7part2(inputData)

			fmt.Printf("Result: %v\n", res)
		},
	},
	)
}
