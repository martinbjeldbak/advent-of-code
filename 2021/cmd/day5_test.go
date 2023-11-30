/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"testing"

	"github.com/martinbjeldbak/advent-of-code/testutils"
)

func Test_overlappingLines(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Test example from https://adventofcode.com/2021/day/5",
			args: args{lines: testutils.ParseInputFile("../test/day5_example_input.txt")},
			want: 5,
		},
		{name: "Problem from https://adventofcode.com/2021/day/5",
			args: args{lines: testutils.ParseInputFile("../test/day5_input.txt")},
			want: 7473,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := overlappingLines(tt.args.lines); got != tt.want {
				t.Errorf("overlappingLines() = %v, want %v", got, tt.want)
			}
		})
	}
}
