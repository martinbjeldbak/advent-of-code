package cmd

import "testing"

func Test_day1part2(t *testing.T) {
	type args struct {
		inputData []string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "Day 1 part 2 example",
			args:    args{
				inputData: []string{
					"3   4",
					"4   3",
					"2   5",
					"1   3",
					"3   9",
					"3   3",
				},
			},
			want:    31,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := day1part2(tt.args.inputData)
			if (err != nil) != tt.wantErr {
				t.Errorf("day1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("day1() = %v, want %v", got, tt.want)
			}
		})
	}
}
