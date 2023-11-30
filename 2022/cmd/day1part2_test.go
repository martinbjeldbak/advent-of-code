package cmd

import (
	"testing"

	"github.com/martinbjeldbak/advent-of-code/testutils"
)

func Test_twentytwentytwoDay1Part2(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Test example from https://adventofcode.com/2022/day/1#part2",
			args: args{lines: testutils.ParseTestFile("2022d1_example_input.txt")},
			want: 45000,
		},
		{name: "Test example from https://adventofcode.com/2022/day/1#part2",
			args: args{lines: testutils.ParseTestFile("2022d1_input.txt")},
			want: 207148,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twentytwentytwoDay1Part2(tt.args.lines); got != tt.want {
				t.Errorf("d1p2() = %v, want %v", got, tt.want)
			}
		})
	}
}
