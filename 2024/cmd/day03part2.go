package cmd

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func day03part2(inputData []string) (int, error) {
	mulExpr := regexp.MustCompile("mul\\(\\d+,\\d+\\)")
	doExpr := regexp.MustCompile("do\\(\\)")
	dontExpr := regexp.MustCompile("don't\\(\\)")
	program := strings.Join(inputData, "\n") // we don't want line breaks
	sum := 0

	numExpressions := 0
	expressions := mulExpr.FindAllStringIndex(program, -1)
	dos := doExpr.FindAllStringIndex(program, -1)
	donts := dontExpr.FindAllStringIndex(program, -1)

	fmt.Printf("Dos: %v\n", len(dos))
	fmt.Printf("Don'ts: %v\n", len(donts))
	fmt.Printf("Expressions: %v\n", len(expressions))

	numExpressions += len(expressions)
	maxDo := 0
	maxDont := 0

	for _, expressionIdx := range expressions {
		expression := program[expressionIdx[0]:expressionIdx[1]]

		for _, doIdx := range dos {
			if (doIdx[1] - 1) < expressionIdx[0] {
				maxDo = doIdx[0] - 1
			}
		}
		for _, dontIdx := range donts {
			if (dontIdx[1] - 1) < expressionIdx[0] {
				maxDont = dontIdx[0] - 1
			}
		}

		if maxDo == maxDont && maxDo == 0 {
			sum += executeMulExpression(expression)
		}
		if maxDo > maxDont {
			sum += executeMulExpression(expression)
		}
	}

	// fmt.Printf("Num expressions: %v\n", numExpressions)

	return sum, nil
}

func executeMulExpression(expression string) int {
	parts := strings.Split(expression, ",")

	numberRe := regexp.MustCompile("\\d+")
	left, err := strconv.Atoi(numberRe.FindString(parts[0]))
	if err != nil {
		return 0
	}

	right, err := strconv.Atoi(numberRe.FindString(parts[1]))
	if err != nil {
		return 0
	}

	return left * right
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use: "day03part2",
		RunE: func(_ *cobra.Command, _ []string) error {
			res, err := day03part2(inputData)
			if err != nil {
				return err
			}

			fmt.Printf("Result: %v\n", res)

			return nil
		},
	},
	)
}
