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

func Test_day7part2(t *testing.T) {
	type args struct {
		inputData []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example input from https://adventofcode.com/2021/day/7#part2",
			args: args{inputData: []string{"16,1,2,0,4,2,7,1,2,14"}},
			want: 168,
		},
		{
			name: "Problem input",
			args: args{inputData: testutils.ParseTestFile("day7_input.txt")},
			want: 97038163,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day7part2(tt.args.inputData); got != tt.want {
				t.Errorf("day7part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
