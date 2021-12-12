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
)

func Test_day8part2(t *testing.T) {
	type args struct {
		inputData []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Example input from https://adventofcode.com/2021/day/8#part2",
			args: args{inputData: []string{
				"acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf",
			}},
			want: 5353,
		},
		// {name: "Example day input https://adventofcode.com/2021/day/8#part2",
		// 	args: args{inputData: testutils.ParseTestFile("day8_example_data.txt")},
		// 	want: 61229,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day8part2(tt.args.inputData); got != tt.want {
				t.Errorf("day8part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
