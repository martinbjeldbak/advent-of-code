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

func Test_day2(t *testing.T) {
	type args struct {
		commands []string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "https://adventofcode.com/2021/day/2 sample input",
			args: args{commands: []string{
				"forward 5",
				"down 5",
				"forward 8",
				"up 3",
				"down 8",
				"forward 2",
			}},
			want:    150,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := day2(tt.args.commands)
			if (err != nil) != tt.wantErr {
				t.Errorf("day2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("day2() = %v, want %v", got, tt.want)
			}
		})
	}
}
