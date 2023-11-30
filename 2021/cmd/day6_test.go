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

func Test_day6(t *testing.T) {
	type args struct {
		inputData []string
		numdays   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example input from https://adventofcode.com/2021/day/6, 18 days",
			args: args{
				inputData: []string{"3,4,3,1,2"},
				numdays:   18,
			},
			want: 26,
		},
		// {
		// 	name: "Example input from https://adventofcode.com/2021/day/6, 80 days",
		// 	args: args{
		// 		inputData: []string{"3,4,3,1,2"},
		// 		numdays:   80,
		// 	},
		// 	want: 5934,
		// },
		// {
		// 	name: "Problem input, 80 days",
		// 	args: args{inputData: testutils.ParseTestFile("day6_input.txt"),
		// 		numdays: 80,
		// 	},
		// 	want: 5934,
		// },
		// {
		// 	name: "Example input, 256 days",
		// 	args: args{
		// 		inputData: []string{"3,4,3,1,2"},
		// 		numdays:   256,
		// 	},
		// 	want: 26984457539,
		// },
		// {
		// 	name: "Problem input, 256 days",
		// 	args: args{inputData: testutils.ParseInputFile("day6_input.txt"),
		// 		numdays: 256,
		// 	},
		// 	want: 26984457539,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day6(tt.args.inputData, tt.args.numdays); got != tt.want {
				t.Errorf("day6() = %v, want %v", got, tt.want)
			}
		})
	}
}
