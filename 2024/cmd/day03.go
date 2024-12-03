package cmd

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func day03(inputData []string) (int, error) {
	mulExpression := regexp.MustCompile("mul\\(\\d+,\\d+\\)")
	numberRe := regexp.MustCompile("\\d+")
	sum := 0

	for _, program := range inputData {
		expressions := mulExpression.FindAllString(program, -1)

		for _, expression := range expressions {
			parts := strings.Split(expression, ",")

			left, err := strconv.Atoi(numberRe.FindString(parts[0]))
			if err != nil {
				return 0, err
			}

			right, err := strconv.Atoi(numberRe.FindString(parts[1]))
			if err != nil {
				return 0, err
			}

			sum = sum + left * right
		}
	}

	return sum, nil
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "day03",
		RunE: func(_ *cobra.Command, _ []string) error {
			res, err := day03(inputData)
			if err != nil {
				return err
			}

			fmt.Printf("Result: %v\n", res)

			return nil
		},
	},
	)
}
