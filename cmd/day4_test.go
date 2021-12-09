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

import "testing"

func Test_day4(t *testing.T) {
	type args struct {
		draws  []int
		boards [][][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Sample from https://adventofcode.com/2021/day/4",
			args: args{
				draws: []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1},
				boards: [][][]int{
					{
						{22, 13, 17, 11, 0},
						{8, 2, 23, 4, 24},
						{21, 9, 14, 16, 7},
						{6, 10, 3, 18, 5},
						{1, 12, 20, 15, 19},
					},
					{
						{3, 15, 0, 2, 22},
						{9, 18, 13, 17, 5},
						{19, 8, 7, 25, 23},
						{20, 11, 10, 24, 4},
						{14, 21, 16, 12, 6},
					},
					{
						{14, 21, 17, 24, 4},
						{10, 16, 15, 9, 19},
						{18, 8, 23, 26, 20},
						{22, 11, 13, 6, 5},
						{2, 0, 12, 3, 7},
					},
				},
			},
			want: 4512,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day4(tt.args.draws, tt.args.boards); got != tt.want {
				t.Errorf("day4() = %v, want %v", got, tt.want)
			}
		})
	}
}
