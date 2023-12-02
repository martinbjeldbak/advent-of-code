package cmd

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/spf13/cobra"
)

func day2part2(inputData []string) (int, error) {
	fullGameRegex := regexp.MustCompile(`(Game (\d+))(.*)`)
	pullRegex := regexp.MustCompile(`(\d+ \w+)(, )?(; )?`)
	redMatcher := regexp.MustCompile(`(\d+) red`)
	greenMatcher := regexp.MustCompile(`(\d+) green`)
	blueMatcher := regexp.MustCompile(`(\d+) blue`)

	sum := 0

	for _, game := range inputData {
		parsedGame := fullGameRegex.FindStringSubmatch(game)

		pulls := pullRegex.FindAllString(parsedGame[3], -1)

		maxRed := 0
		maxGreen := 0
		maxBlue := 0

		for _, grab := range pulls {
			redPull := redMatcher.FindStringSubmatch(grab)
			if redPull != nil {
				redCount, _ := strconv.Atoi(redPull[1])
				if redCount > maxRed {
					maxRed = redCount
				}
			}

			greenPull := greenMatcher.FindStringSubmatch(grab)
			if greenPull != nil {
				greenCount, _ := strconv.Atoi(greenPull[1])
				if greenCount > maxGreen {
					maxGreen = greenCount
				}
			}

			bluePull := blueMatcher.FindStringSubmatch(grab)
			if bluePull != nil {
				blueCount, _ := strconv.Atoi(bluePull[1])
				if blueCount > maxBlue {
					maxBlue = blueCount
				}
			}
		}
		sum += maxRed * maxGreen * maxBlue
	}

	return sum, nil
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "day2part2",
		RunE: func(_ *cobra.Command, _ []string) error {
			res, err := day2part2(inputData)
			if err != nil {
				return err
			}

			fmt.Printf("Result: %v\n", res)

			return nil
		},
	},
	)
}
