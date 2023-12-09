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
				},
			},
			want: 0,
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
