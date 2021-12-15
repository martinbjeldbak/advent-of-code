package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// Checks matrix to see if all values are 0
func isSimultaneousFlash(energyLevels [][]int) (simultaneousFlash bool) {
	n := len(energyLevels)
	m := len(energyLevels[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if energyLevels[i][j] != 0 {
				return false
			}
		}
	}

	return true
}

func day11part2(inputData []string) (firstSyncFlashStep int) {
	n := len(inputData)
	m := len(strings.Split(inputData[0], ""))

	energyLevels := make([][]int, n)
	for i := range energyLevels {
		energyLevels[i] = make([]int, 0, m)
	}

	// Convert input data to 10x10 matrix
	for i, row := range inputData {
		rowEnergyData := strings.Split(row, "")
		for _, v := range rowEnergyData {
			l, _ := strconv.Atoi(v)
			energyLevels[i] = append(energyLevels[i], l)
		}
	}

	fmt.Println("Before any steps:")
	prettyPrint2D(energyLevels)

	for !isSimultaneousFlash(energyLevels) {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				energyLevels[i][j]++
			}
		}

		energyLevels, _ := flash(energyLevels)

		for needFlash(energyLevels) {
			energyLevels, _ = flash(energyLevels)
		}

		energyLevels = resetFlashedEnergy(energyLevels)

		prettyPrint2D(energyLevels)
		firstSyncFlashStep++
	}

	return firstSyncFlashStep
}

var day11part2Cmd = &cobra.Command{
	Use: "day11part2",
	Run: func(cmd *cobra.Command, args []string) {
		res := day11part2(inputData)

		fmt.Printf("Res: %v\n", res)
	},
}

func init() {
	rootCmd.AddCommand(day11part2Cmd)
}
