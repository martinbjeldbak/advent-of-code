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
			name: "Example",
			args: args{
				inputData: []string{
					"two1nine",
					"eightwothree",
					"abcone2threexyz",
					"xtwone3four",
					"4nineeightseven2",
					"zoneight234",
					"7pqrstsixteen",
				},
			},
			want:    281,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := day1part2(tt.args.inputData)
			if (err != nil) != tt.wantErr {
				t.Errorf("day1part2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("day1part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
