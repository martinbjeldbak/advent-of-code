package cmd

import "testing"

func Test_day2part2(t *testing.T) {
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
			name:    "Test input",
			args:    args{

				inputData: []string{
					"7 6 4 2 1",
					"1 2 7 8 9",
					"9 7 6 2 1",
					"1 3 2 4 5",
					"8 6 4 4 1",
					"1 3 6 7 9",
				},
			},
			want:    4,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := day2part2(tt.args.inputData)
			if (err != nil) != tt.wantErr {
				t.Errorf("day2part2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("day2part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
