package cmd

import (
	"testing"

	"github.com/martinbjeldbak/advent-of-code/testutils"
)

func Test_day11(t *testing.T) {
	type args struct {
		inputData []string
		numSteps  int
	}
	tests := []struct {
		name           string
		args           args
		wantNumFlashes int
	}{
		{
			name: "Smaller sample data from https://adventofcode.com/2021/day/11",
			args: args{numSteps: 1, inputData: []string{
				"11111",
				"19991",
				"19191",
				"19991",
				"11111",
			}},
			wantNumFlashes: 9,
		},
		{
			name: "Smaller sample data from https://adventofcode.com/2021/day/11",
			args: args{numSteps: 2, inputData: []string{
				"11111",
				"19991",
				"19191",
				"19991",
				"11111",
			}},
			wantNumFlashes: 9,
		},
		{
			name:           "Sample data from https://adventofcode.com/2021/day/11",
			args:           args{numSteps: 2, inputData: testutils.ParseTestFile("day11_example_input.txt")},
			wantNumFlashes: 35,
		},
		{
			name:           "Sample data from https://adventofcode.com/2021/day/11",
			args:           args{numSteps: 100, inputData: testutils.ParseTestFile("day11_example_input.txt")},
			wantNumFlashes: 1656,
		},
		{
			name:           "Problem input",
			args:           args{numSteps: 100, inputData: testutils.ParseTestFile("day11_input.txt")},
			wantNumFlashes: 1700,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNumFlashes := day11(tt.args.inputData, tt.args.numSteps); gotNumFlashes != tt.wantNumFlashes {
				t.Errorf("day11() = %v, want %v", gotNumFlashes, tt.wantNumFlashes)
			}
		})
	}
}
