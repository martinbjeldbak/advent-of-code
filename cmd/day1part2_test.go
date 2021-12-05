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

func Test_day2part2(t *testing.T) {
	type args struct {
		measurements []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"From website", args{[]string{
			"199  A      ",
			"200  A B    ",
			"208  A B C  ",
			"210    B C D",
			"200  E   C D",
			"207  E F   D",
			"240  E F G  ",
			"269    F G H",
			"260      G H",
			"263        H",
		}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day2part2(tt.args.measurements); got != tt.want {
				t.Errorf("day2part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
