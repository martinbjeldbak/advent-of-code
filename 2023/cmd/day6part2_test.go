package cmd

import "testing"

func Test_day6part2(t *testing.T) {
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
					"Time:      7  15   30",
					"Distance:  9  40  200",
				},
			},
			want: 71503,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day6part2(tt.args.inputData); got != tt.want {
				t.Errorf("day6part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
