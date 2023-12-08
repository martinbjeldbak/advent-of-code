package cmd

import "testing"

func Test_day7(t *testing.T) {
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
					"32T3K 765",
					"T55J5 684",
					"KK677 28",
					"KTJJT 220",
					"QQQJA 483",
				},
			},
			want: 6440,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day7(tt.args.inputData); got != tt.want {
				t.Errorf("day7() = %v, want %v", got, tt.want)
			}
		})
	}
}
