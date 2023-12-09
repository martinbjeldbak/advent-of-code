package cmd

import "testing"

func Test_day8(t *testing.T) {
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
					"RL",
					"",
					"AAA = (BBB, CCC)",
					"BBB = (DDD, EEE)",
					"CCC = (ZZZ, GGG)",
					"DDD = (DDD, DDD)",
					"EEE = (EEE, EEE)",
					"GGG = (GGG, GGG)",
					"ZZZ = (ZZZ, ZZZ)",
				},
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				inputData: []string{
					"LLR",
					"",
					"AAA = (BBB, BBB)",
					"BBB = (AAA, ZZZ)",
					"ZZZ = (ZZZ, ZZZ)",
				},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day8(tt.args.inputData); got != tt.want {
				t.Errorf("day8() = %v, want %v", got, tt.want)
			}
		})
	}
}
