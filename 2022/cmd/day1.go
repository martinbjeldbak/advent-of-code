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

func twentytwentytwoDay1(inputData []string) int {
	maxHeldCalories := 0
	currHeldCalories := 0
	for _, calories := range inputData {
		if calories == "" {
			if currHeldCalories > maxHeldCalories {
				maxHeldCalories = currHeldCalories
			}

			currHeldCalories = 0
			continue
		}

		cal, err := strconv.Atoi(calories)
		if err != nil {
			log.Fatal(err)
		}
		currHeldCalories += cal
	}

	return maxHeldCalories
}

// twentytwentytwoDay1Cmd represents the twentytwentytwoDay1 command
var twentytwentytwoDay1Cmd = &cobra.Command{
	Use: "twentytwentytwoDay1",
	Run: func(cmd *cobra.Command, args []string) {
		res := twentytwentytwoDay1(inputData)
		fmt.Printf("Result: %v\n", res)
	},
}

func init() {
	rootCmd.AddCommand(twentytwentytwoDay1Cmd)
}
