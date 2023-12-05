package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

type mapping struct {
	src    int
	dst    int
	length int
}

func (m *mapping) dest(source int) (int, error) {
	return -1, nil
}

func (m *mapping) String() string {
	return fmt.Sprintf("(dst: %v, src: %v, len: %v)", m.dst, m.src, m.length)
}

func day5(inputData []string) int {
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

	fmt.Println(seeds)
	fmt.Println(almanac)

	return -1
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "day5",
		Run: func(_ *cobra.Command, _ []string) {
			res := day5(inputData)

			fmt.Printf("Result: %v\n", res)
		},
	},
	)
}
