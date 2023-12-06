package cmd

import "testing"

func Test_day6(t *testing.T) {
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
			want: 288, // 4 * 8 * 9
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day6(tt.args.inputData); got != tt.want {
				t.Errorf("day6() = %v, want %v", got, tt.want)
			}
		})
	}
}
