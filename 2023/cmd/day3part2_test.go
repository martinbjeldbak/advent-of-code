package cmd

import "testing"

func Test_day3part2(t *testing.T) {
	type args struct {
		inputData []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example input",
			args: args{
				inputData: []string{
					"467..114..",
					"...*......",
					"..35..633.",
					"......#...",
					"617*......",
					".....+.58.",
					"..592.....",
					"......755.",
					"...$.*....",
					".664.598..",
				},
			},
			want: 467835,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day3part2(tt.args.inputData); got != tt.want {
				t.Errorf("day3part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
