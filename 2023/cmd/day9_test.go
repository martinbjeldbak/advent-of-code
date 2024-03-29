package cmd

import "testing"

func Test_day9(t *testing.T) {
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
					"0 3 6 9 12 15",
					"1 3 6 10 15 21",
					"10 13 16 21 30 45",
				},
			},
			want: 114,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := day9(tt.args.inputData); got != tt.want {
				t.Errorf("day9() = %v, want %v", got, tt.want)
			}
		})
	}
}
