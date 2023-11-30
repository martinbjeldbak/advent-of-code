package cmd

import (
	"testing"

	"github.com/martinbjeldbak/advent-of-code/testutils"
)

func Test_day5part2(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Test example from https://adventofcode.com/2021/day/5#part2",
			args: args{lines: testutils.ParseTestFile("day5part2_example_input.txt")},
			want: 12,
		},
		{name: "Problem input https://adventofcode.com/2021/day/5#part2",
			args: args{lines: testutils.ParseTestFile("day5_input.txt")},
			want: 24164,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day5part2(tt.args.lines); got != tt.want {
				t.Errorf("day5part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
