/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

func twentytwentytwoDay1Part2(inputData []string) int {
	totalPerElf := make([]int, 0)
	currHeldCalories := 0
	for _, calories := range inputData {
		if calories == "" {
			totalPerElf = append(totalPerElf, currHeldCalories)
			currHeldCalories = 0
			continue
		}

		cal, err := strconv.Atoi(calories)
		if err != nil {
			log.Fatal(err)
		}
		currHeldCalories += cal
	}
	// append the last elf
	totalPerElf = append(totalPerElf, currHeldCalories)

	topThree := make([]int, 3)
	smallestIndex := 0
	for _, elfCalories := range totalPerElf {
		if elfCalories > topThree[smallestIndex] {
			topThree[smallestIndex] = elfCalories
		}

		for i, cals := range topThree {
			if cals < topThree[smallestIndex] {
				smallestIndex = i
			}

		}
	}

	totalCalories := 0
	for _, elfCalories := range topThree {
		totalCalories += elfCalories
	}

	return totalCalories
}

// twentytwentytwoDay1Cmd represents the twentytwentytwoDay1 command
var twentytwentytwoDay1Part2Cmd = &cobra.Command{
	Use: "twentytwentytwoDay1",
	Run: func(cmd *cobra.Command, args []string) {
		res := twentytwentytwoDay1(inputData)
		fmt.Printf("Result: %v\n", res)
	},
}

func init() {
	rootCmd.AddCommand(twentytwentytwoDay1Part2Cmd)
}
