package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func day5part2(inputData []string) int {
	seedsRaw := strings.Fields(strings.Split(inputData[0], ":")[1])
	seeds := make([]int, 0, len(seedsRaw))

	totalAlamancCount := 7
	almanac := make([][]*mapping, totalAlamancCount)
	for i := 0; i < totalAlamancCount; i++ {
		almanac[i] = make([]*mapping, 0)
	}

	for _, s := range seedsRaw {
		seed, _ := strconv.Atoi(s)
		seeds = append(seeds, seed)
	}

	curAlmanac := 0
	for i := 1; i < len(inputData); i++ {
		row := inputData[i]

		if strings.Contains(row, "map:") {
			// name := strings.Fields(row)[0]

			i++
			for i < len(inputData) && inputData[i] != "" {
				instructionsRaw := strings.Fields(inputData[i])
				instructions := make([]int, 0, len(instructionsRaw))
				for _, v := range instructionsRaw {
					value, _ := strconv.Atoi(v)

					instructions = append(instructions, value)
				}

				almanac[curAlmanac] = append(almanac[curAlmanac], &mapping{
					dst:    instructions[0],
					src:    instructions[1],
					length: instructions[2],
				})
				i++
			}
			curAlmanac++
		}
	}

	memo := make(map[int]int)

	for curSeedIndex := 0; curSeedIndex < len(seeds); curSeedIndex += 2 {
		curSeed := seeds[curSeedIndex]

		for offset := 0; offset < seeds[curSeedIndex+1]; offset++ {
			loc := curSeed + offset

			_, ok := memo[loc]

			if ok {
				continue
			}

			for curAlmanac := 0; curAlmanac < len(almanac); curAlmanac++ {
				loc = destinationFor(loc, almanac[curAlmanac])
			}

			memo[curSeed+offset] = loc
		}
	}

	minLoc := memo[seeds[0]]
	for _, v := range memo {
		if v < minLoc {
			minLoc = v
		}
	}

	return minLoc
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "day5part2",
		Run: func(_ *cobra.Command, _ []string) {
			start := time.Now()
			res := day5part2(inputData)

			fmt.Printf("Result: %v ran in %v\n", res, time.Since(start))
		},
	},
	)
}
