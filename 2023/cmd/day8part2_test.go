package cmd

import "testing"

func Test_day8part2(t *testing.T) {
	type args struct {
		inputData []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example",
			args: args{
				inputData: []string{
					"LR",
					"",
					"11A = (11B, XXX)",
					"11B = (XXX, 11Z)",
					"11Z = (11B, XXX)",
					"22A = (22B, XXX)",
					"22B = (22C, 22C)",
					"22C = (22Z, 22Z)",
					"22Z = (22B, 22B)",
					"XXX = (XXX, XXX)",
				},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day8part2(tt.args.inputData); got != tt.want {
				t.Errorf("day8part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
