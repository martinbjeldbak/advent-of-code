package cmd

import (
	"testing"

	"github.com/martinbjeldbak/advent-of-code/testutils"
)

func Test_day11part2(t *testing.T) {
	type args struct {
		inputData []string
	}
	tests := []struct {
		name                      string
		args                      args
		wantFirstStepSimultaneous int
	}{
		{
			name:                      "Sample data from https://adventofcode.com/2021/day/11",
			args:                      args{inputData: testutils.ParseTestFile("day11_example_input.txt")},
			wantFirstStepSimultaneous: 195,
		},
		{
			name:                      "Problem input",
			args:                      args{inputData: testutils.ParseTestFile("day11_input.txt")},
			wantFirstStepSimultaneous: 273,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFirstStepSimultaneous := day11part2(tt.args.inputData); gotFirstStepSimultaneous != tt.wantFirstStepSimultaneous {
				t.Errorf("day11() = %v, want %v", gotFirstStepSimultaneous, tt.wantFirstStepSimultaneous)
			}
		})
	}
}
