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

func Test_dayOne(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		out  int
	}{
		{"Given example", []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 363}, 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			num_increases := dayOne(tt.in)

			if num_increases != tt.out {
				t.Errorf("got %v, want %v", num_increases, tt.out)
			}
		})
	}
}
