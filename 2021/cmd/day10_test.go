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

func Test_day10(t *testing.T) {
	type args struct {
		inputData []string
	}
	tests := []struct {
		name      string
		args      args
		wantScore int
	}{
		{
			name:      "Example data from https://adventofcode.com/2021/day/10",
			args:      args{inputData: testutils.ParseTestFile("day10_example_input.txt")},
			wantScore: 26397,
		},
		{
			name:      "Example data from https://adventofcode.com/2021/day/10",
			args:      args{inputData: testutils.ParseTestFile("day10_input.txt")},
			wantScore: 392043,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotScore := day10(tt.args.inputData); gotScore != tt.wantScore {
				t.Errorf("day10() = %v, want %v", gotScore, tt.wantScore)
			}
		})
	}
}
